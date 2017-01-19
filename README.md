dargo
=====
Deploy apps directly to docker.

[![GoDoc](https://godoc.org/github.com/brimstone/dargo?status.svg)](https://godoc.org/github.com/brimstone/dargo)

Why?
====
Inspired from Kelsey Hightower's [presentation at dotGo 2016](https://www.youtube.com/watch?v=nhmAyZNlECw)
of [Kargo](https://github.com/kelseyhightower/kargo) library, I wanted also be
able to deploy applications directly to Docker Swarm


Usage
=====
Check out the example app in `example`. It's as simple as:
```
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
```
