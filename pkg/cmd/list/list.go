package list

import (
	"github.com/Ruffel/ssdutil/pkg/disk"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the supported devices present in the system",
		Long:  "Scan the system for a list of storage devices and show relevant identifying information",
		RunE:  showDrives,
	}

	return cmd
}

func showDrives(cmd *cobra.Command, args []string) error {

	response, err := disk.ListDrives()

	if err != nil {
		return errors.Wrap(err, "Failed to fetch drive information from host system")
	}

	for _, drive := range response.Disks {
		println(drive.Name)
	}

	return nil
}
