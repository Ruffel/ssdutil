package list

import (
	"github.com/Ruffel/ssdutil/pkg/disk"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the supported devices present in the system",
		Long:  "Scan the system for a list of storage devices and show relevant identifying information",
		Run: func(cmd *cobra.Command, args []string) {

			response, _ := disk.ListDrives()

			for _, drive := range response.Disks {
				println(drive.Name)
			}
		},
	}

	return cmd
}
