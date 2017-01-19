dargo
=====
Deploy apps directly to docker.

Why?
====
Inspired from Kelsey Hightower's
[Kargo](https://github.com/kelseyhightower/kargo) library, I wanted also be able
to deploy applications directly to Docker Swarm

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
