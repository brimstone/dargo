package main

import (
	"fmt"

	"github.com/brimstone/dargo"
)

func main() {
	dargo.DeployAndExit(dargo.DeployOptions{
		Tags: []string{"hello-dargo"},
	})
	fmt.Println("Hello world!")
}
