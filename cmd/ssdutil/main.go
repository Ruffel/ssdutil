package main

import (
	"os"

	"github.com/Ruffel/ssdutil/internal/build"
	"github.com/Ruffel/ssdutil/pkg/cmd/root"
)

func main() {
	cmd := root.NewCmdRoot(build.Version)

	if _, err := cmd.ExecuteC(); err != nil {
		os.Exit(1)
	}
}
