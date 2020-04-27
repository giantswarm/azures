#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

# For future editors: there is a lot of pushd/popd in this script.
# When we run certain commands like `go build`, go pulls versions
# from `go.mod` in the current directory. Additionally, come commands
# can't generate files outside of the current directory. For this reason,
# we have to move between the script directory and the repo root directory
# several times.

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
toolpath="$dir/bin"

export GOPATH=$(go env GOPATH)
export GOROOT=$(go env GOROOT)

# We need specific versions of code generation tools so we install them locally
# in scripts/bin using versions from scripts/go.mod
install_tool() {
  local module=$1
  local package=$2
  local bin=$3
  rm go.sum 2> /dev/null || true
  version=$(go list -m -f '{{.Version}}' "$module")
  echo "Rebuilding $bin@$version"
  mkdir -p "$toolpath"
  go build -o "$toolpath/$bin" "$module""$package" 2> /dev/null
  git checkout go.mod 2> /dev/null
}

pushd "$dir" > /dev/null

install_tool sigs.k8s.io/controller-tools /cmd/controller-gen controller-gen
install_tool k8s.io/code-generator /cmd/deepcopy-gen deepcopy-gen
install_tool k8s.io/code-generator /cmd/client-gen client-gen
install_tool golang.org/x/tools /cmd/goimports goimports
install_tool github.com/markbates/pkger /cmd/pkger pkger
install_tool sigs.k8s.io/kustomize/kustomize/v3 "" kustomize

popd > /dev/null

# Set up variables for deepcopy-gen and client-gen
module="github.com/giantswarm/apiextensions"
input_dirs=$(find ./pkg/apis -maxdepth 2 -mindepth 2 | tr '\r\n' ',')
input_dirs=${input_dirs%?}
groups=${input_dirs//.\/pkg\/apis\//}
header=scripts/boilerplate.go.txt

# deepcopy-gen creates DeepCopy functions for each custom resource
echo "Generating deepcopy funcs"
"$toolpath/deepcopy-gen" \
  --input-dirs "$input_dirs" \
  --output-file-base zz_generated.deepcopy \
  --go-header-file "$header"

# client-gen creates typed go clients for CRUD operations for each custom resource
echo "Generating clientset"
"$toolpath/client-gen" \
  --clientset-name versioned \
  --input "$groups" \
  --input-base "$module/pkg/apis" \
  --output-package "$module/pkg/clientset" \
  --output-base "$dir" \
  --go-header-file "$header"

# client-gen expects to be run in $GOPATH so it generates files in
# ./github.com/giantswarm/apiextensions/pkg/clientset which need to
# be manually moved into pkg.
echo "Moving generated files to expected location"
cp -R "$dir/$module/pkg"/* "$dir/../pkg"
rm -rf "$dir/github.com"

# code-generator doesn't group local imports separately from third-party
# imports, so run goimports after generating new code.
pushd "$dir/.." > /dev/null
echo "Fixing imports in-place with goimports"
"$toolpath/goimports" -local $module -w ./pkg

# Ensure that we have fresh CRDs in case any have been deleted
rm -rf "$dir/../config/crd/bases"

# Using kubebuilder comments, create new CRDs from CR definitions in source files
echo "Generating all CRDs as v1beta1"
"$toolpath/controller-gen" \
  crd \
  paths=./pkg/apis/... \
  output:dir=config/crd/bases \
  crd:crdVersions=v1beta1

# Overwrite CRDs infrastructure.giantswarm.io as v1 until all other
# groups can be migrated to v1
echo "Generating infrastructure.giantswarm.io CRDs as v1"
"$toolpath/controller-gen" \
  crd \
  paths=./pkg/apis/infrastructure/v1alpha2 \
  output:dir=config/crd/bases \
  crd:crdVersions=v1

# Generate Cluster API CRDs from external library (version from scripts/go.mod)
pushd "$dir" > /dev/null
"$toolpath/controller-gen" \
  crd \
  paths=sigs.k8s.io/cluster-api/api/v1alpha2 \
  output:dir=../config/crd/bases \
  crd:crdVersions=v1beta1
popd > /dev/null
# We only want Cluster and MachineDeployment for now, so delete the other two CAPI CRDs.
rm config/crd/bases/cluster.x-k8s.io_machines.yaml
rm config/crd/bases/cluster.x-k8s.io_machinesets.yaml

# Add .metadata.name validation to Release CRD using kustomize since
# kubebuilder comments can't modify metav1.ObjectMeta
echo "Kustomizing CRDs"
"$toolpath/kustomize" build \
  config/crd \
  -o config/crd/bases/release.giantswarm.io_releases.yaml

# Package CRD YAMLs into a virtual filesystem so they can be accessed by New*CRD() functions
hash=$(find config/crd/bases -type f -print0 | xargs -0 sha1sum | sort -df | sha1sum)
prevhash=$(cat "$dir/.hash" 2> /dev/null || echo "")
if [ "$hash" != "$prevhash" ]; then
  echo "Detected changes in CRD YAMLs"
  echo "Using pkger to package CRDs into go source virtual file system"
  "$toolpath/pkger" -include /config/crd/bases -o pkg/crd
  echo "$hash" > "$dir/.hash"
fi

echo "Applying linter patch to generated files"
git apply "$dir/generated.patch"
