package list

import "github.com/spf13/cobra"

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the supported devices present in the system",
		Long:  "Scan the system for a list of storage devices and show relevant identifying information",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
