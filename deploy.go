package dargo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// DeployOptions control attributes about the resulting Docker Image
type DeployOptions struct {
	Tags []string
}

// Deploy deploys to the Docker engine reachable via environment variables and
// defaults.
func Deploy(o DeployOptions) error {

	buildOptions := types.ImageBuildOptions{
		Tags: o.Tags,
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	self, err := makeTar()
	if err != nil {
		return err
	}

	buildResponse, err := cli.ImageBuild(context.Background(), self, buildOptions)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	all, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}

// DeployAndExit checks to see if we're Enabled, then calls Deploy, logging
// Fatal errors, then exiting
func DeployAndExit(o DeployOptions) {
	if !Enable {
		return
	}
	if err := Deploy(o); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
