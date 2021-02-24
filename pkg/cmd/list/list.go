package list

import (
	"os"

	"github.com/Ruffel/ssdutil/pkg/disk"
	"github.com/olekukonko/tablewriter"
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Model", "Serial", "Firmware", "Interface"})

	for _, drive := range response.Disks {
		table.Append([]string{drive.Name, drive.Model, drive.SerialNumber, drive.FirmwareVersion, drive.InterfaceType.String()})
	}

	table.Render()

	return nil
}
