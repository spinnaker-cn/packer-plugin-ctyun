package main

import (
	"fmt"
	"github.com/ctyun/packer-plugin-ctyun/builder/ctyun/basic"
	"os"

	"github.com/hashicorp/packer/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("basic", new(basic.Builder))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
