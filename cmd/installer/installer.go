package main

import (
	"context"
	"flag"

	"github.com/deislabs/smi-sdk-go/pkg/installer"
	"github.com/kubernetes-client/go/kubernetes/client"
	"github.com/kubernetes-client/go/kubernetes/config"
)

var (
	dir = flag.String("dir", "./", "The directory to load YAML from")
)

func main() {
	flag.Parse()

	c, err := config.LoadKubeConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset := client.NewAPIClient(c)
	err = installer.InstallTypesFromYAML(context.Background(), *dir, clientset)
	if err != nil {
		panic(err.Error())
	}
}