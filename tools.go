// +build tools

package main

import (
	_ "github.com/k14s/kapp/cmd/kapp" // clone
	_ "github.com/k14s/kbld/cmd/kbld"
	_ "github.com/k14s/vendir/cmd/vendir"
	_ "github.com/k14s/ytt/cmd/ytt" // clone
	_ "github.com/cloudfoundry/bosh-cli" // clone
)
