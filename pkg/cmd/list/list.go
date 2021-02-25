package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ruffel/ssdutil/pkg/disk"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Interface string
}

func NewCmdList() *cobra.Command {
	opts := &ListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the supported devices present in the system",
		Long:  "Scan the system for a list of storage devices and show relevant identifying information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return showDrives(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Interface, "interface", "i", "all", "Filter by interface type: {all|nvme|sata|unknown|scsi}")

	return cmd
}

func showDrives(opts *ListOptions) error {

	response, err := disk.ListDrives()

	if err != nil {
		return errors.Wrap(err, "Failed to fetch drive information from host system")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Model", "Serial", "Firmware", "Interface"})

	var filterInterface disk.InterfaceType

	switch strings.ToLower(opts.Interface) {
	case "nvme":
		filterInterface = disk.INTERFACE_TYPE_NVME
	case "sata":
		filterInterface = disk.INTERFACE_TYPE_ATA
	case "scsi":
		filterInterface = disk.INTERFACE_TYPE_SCSI
	case "unknown":
		filterInterface = disk.INTERFACE_TYPE_UNKNOWN
	case "all":
		break
	default:
		return fmt.Errorf("Invalid interface type: %s", opts.Interface)
	}

	for _, drive := range response.Disks {
		if opts.Interface != "all" {
			if filterInterface != drive.InterfaceType {
				continue
			}
		}

		table.Append([]string{drive.Name, drive.Model, drive.SerialNumber, drive.FirmwareVersion, drive.InterfaceType.String()})
	}

	if table.NumLines() != 0 {
		table.Render()
	}

	return nil
}
