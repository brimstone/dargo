package dargo

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
)

// DeployOptions control attributes about the resulting Docker Image
type DeployOptions struct {
	Tags       []string
	Foreground bool
}

type buildOutput struct {
	Stream string `json:"stream"`
}

func buildImage(o DeployOptions, cli *client.Client) error {
	buildOptions := types.ImageBuildOptions{
		Tags: o.Tags,
	}

	self, err := makeTar()
	if err != nil {
		return fmt.Errorf("Error building tar for image build: %v", err)
	}

	buildResponse, err := cli.ImageBuild(context.Background(), self, buildOptions)
	if err != nil {
		return fmt.Errorf("Error starting image build: %v", err)
	}
	defer buildResponse.Body.Close()

	scanner := bufio.NewScanner(buildResponse.Body)

	for scanner.Scan() {
		var output buildOutput
		err = json.Unmarshal([]byte(scanner.Text()), &output)
		fmt.Print(output.Stream)
	}

	return nil
}

func runImage(runOptions DeployOptions, cli *client.Client) error {
	containerConfig := container.Config{
		Image: runOptions.Tags[0],
	}

	if runOptions.Foreground {
		containerConfig.AttachStdin = true
		containerConfig.AttachStdout = true
		containerConfig.AttachStderr = true
		containerConfig.Tty = true
		containerConfig.OpenStdin = true
		containerConfig.StdinOnce = true
	}
	networkConfig := network.NetworkingConfig{}
	hostConfig := container.HostConfig{}

	createResponse, err := cli.ContainerCreate(context.Background(), &containerConfig, &hostConfig, &networkConfig, "")
	if err != nil {
		return fmt.Errorf("Error creating container: %v", err)
	}

	attachResp, err := cli.ContainerAttach(context.Background(), createResponse.ID, types.ContainerAttachOptions{
		Stdin:  true,
		Stdout: true,
		Stderr: true,
		Stream: true,
	})
	go func() { io.Copy(attachResp.Conn, os.Stdin) }()
	go func() { io.Copy(os.Stdout, attachResp.Reader) }()

	err = cli.ContainerStart(context.Background(), createResponse.ID, types.ContainerStartOptions{})

	if err != nil {
		return fmt.Errorf("Error starting container: %v", err)
	}

	_, err = cli.ContainerWait(context.Background(), createResponse.ID)
	if err != nil {
		return fmt.Errorf("Error waiting for container to exit: %v", err)
	}
	return nil
}

// Deploy deploys to the Docker engine reachable via environment variables and
// defaults.
func Deploy(o DeployOptions) error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("Error creating docker client: %v", err)
	}

	err = buildImage(o, cli)
	if err != nil {
		return fmt.Errorf("Error building image: %v", err)
	}

	err = runImage(o, cli)
	if err != nil {
		return fmt.Errorf("Error starting image: %v", err)
	}

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
