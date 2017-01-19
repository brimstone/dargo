// Package dargo enables applications to deploy themselves to a Docker engine.
package dargo

import "os"

// Enable determines if we should deploy when Deploy() is called
var Enable bool

func init() {
	if len(os.Args) > 1 {
		if os.Args[1] == "deploy" {
			Enable = true
		}
	}
}
