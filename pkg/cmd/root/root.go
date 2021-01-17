package root

import (
	"github.com/Ruffel/ssdutil/pkg/cmd/list"
	"github.com/spf13/cobra"
)

func NewCmdRoot(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssdutil",
		Short: "CLI tool for querying and managing storage devices",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	// Add --version option to command line
	cmd.Version = version
	cmd.Flags().Bool("version", false, "Show version information")

	// Subcommands
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
