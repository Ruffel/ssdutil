package root

import (
	"github.com/Ruffel/ssdutil/pkg/cmd/list"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssdutil",
		Short: "CLI tool for querying and managing storage devices",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
