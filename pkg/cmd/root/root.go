package root

import (
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssdutil",
		Short: "CLI tool for querying and managing storage devices",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
