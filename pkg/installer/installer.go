package installer

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path"
	"github.com/kubernetes-client/go/kubernetes/client"
	"gopkg.in/yaml.v2"
)

// InstallTypesFromYAML installs a collection of CRDs loaded from a directory
func InstallTypesFromYAML(ctx context.Context, dir string, c *client.APIClient) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	crd := []client.V1beta1CustomResourceDefinition{}
	for ix := range files {
		file, err := os.Open(path.Join(dir, files[ix].Name()))
		if err != nil {
			return err
		}
	    d := yaml.NewDecoder(file)
	
		for {
			var def client.V1beta1CustomResourceDefinition
			if d.Decode(&def) != nil {
				break
			}
			crd = append(crd, def)
		}
	}
	return InstallTypes(ctx, crd, c)
}

// InstallTypes installs a collection of CRDs into a cluster
func InstallTypes(ctx context.Context, crd []client.V1beta1CustomResourceDefinition, c *client.APIClient) error {
	for ix := range crd {
		if crd[ix].Metadata == nil {
			continue
		}
		_, _, err := c.ApiextensionsV1beta1Api.ReadCustomResourceDefinition(ctx, crd[ix].Metadata.Name, nil)
		if err == nil {
			// TODO: validate that the object is the same here.
			log.Printf("%v already installed", crd[ix].Metadata.Name)
			continue
		}
		if _, _, err := c.ApiextensionsV1beta1Api.CreateCustomResourceDefinition(ctx, crd[ix], nil); err != nil {
			return err
		}
		log.Printf("%v Installed.", crd[ix].Metadata.Name)
	}
	return nil
}