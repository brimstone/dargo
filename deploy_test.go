package dargo_test

import (
	"fmt"

	"github.com/brimstone/dargo"
)

func ExampleDeployAndExit() {
	dargo.DeployAndExit(dargo.DeployOptions{
		Tags: []string{"hello-dargo"},
	})

	fmt.Println("Hello world!")
}
