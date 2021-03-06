package crd

import (
	"github.com/google/go-github/v35/github"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Patch func(crd *v1.CustomResourceDefinition)

type Renderer struct {
	GithubClient *github.Client

	LocalCRDDirectory string
	OutputDirectory   string

	Patches map[string]Patch

	UpstreamAssets []ReleaseAssetFileDefinition
}

type ReleaseAssetFileDefinition struct {
	Files    []string
	Owner    string
	Provider string
	Repo     string
	Version  string
}
