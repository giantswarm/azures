package clientset

import (
	"os"
	"path/filepath"

	clientgen "k8s.io/code-generator/cmd/client-gen/generators"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
)

func Generate(genericArgs args.GeneratorArgs) error {
	b, err := genericArgs.NewBuilder()
	if err != nil {
		return err
	}

	c, err := generator.NewContext(b, clientgen.NameSystems(), clientgen.DefaultNameSystem())
	if err != nil {
		return err
	}

	packages := clientgen.Packages(c, &genericArgs)
	if err := c.ExecutePackages(genericArgs.OutputBase, packages); err != nil {
		return err
	}

	if err := os.RemoveAll("pkg/clientset"); err != nil {
		return err
	}
	if err := os.Rename(
		filepath.Join(genericArgs.OutputBase, genericArgs.OutputPackagePath),
		"pkg/clientset",
	); err != nil {
		return err
	}
	if err := os.RemoveAll(genericArgs.OutputBase); err != nil {
		return err
	}
	return nil
}
