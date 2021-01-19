package disk

import (
	"strings"

	"github.com/StackExchange/wmi"
)

type Win32_DiskDrive struct {
	Caption           *string
	DefaultBlockSize  *uint64
	Description       *string
	DeviceID          *string
	Index             *uint32
	InterfaceType     *string
	Manufacturer      *string
	MediaType         *string
	Model             *string
	Name              *string
	Partitions        *int32
	SerialNumber      *string
	Size              *uint64
	TotalCylinders    *int64
	TotalHeads        *int32
	TotalSectors      *int64
	TotalTracks       *int64
	TracksPerCylinder *int32
}

func getDiskDrives() ([]Win32_DiskDrive, error) {
	const query = "SELECT Caption, DefaultBlockSize, Description, DeviceID, Index, InterfaceType, Manufacturer, MediaType, Model, Name, Partitions, SerialNumber, Size, TotalCylinders, TotalHeads, TotalSectors, TotalTracks, TracksPerCylinder FROM Win32_DiskDrive"

	var physicalDisks []Win32_DiskDrive

	if err := wmi.Query(query, &physicalDisks); err != nil {
		return nil, err
	}

	return physicalDisks, nil
}

type MSFT_PhysicalDisk struct {
	FriendlyName     *string
	HealthStatus     *uint16
	BusType          *uint16
	MediaType        *uint16
	DeviceID         *string
	SerialNumber     *string
	SoftwareVersion  *string
	FirmwareVersion  *string
	Size             *uint64
	AllocatedSize    *uint64
	Model            *string
	PhysicalLocation *string
}

func getPhysicalDisks() ([]MSFT_PhysicalDisk, error) {
	const query = "SELECT FriendlyName, HealthStatus, BusType, MediaType, DeviceID, SerialNumber, SoftwareVersion, FirmwareVersion, Size, AllocatedSize, Model, PhysicalLocation FROM MSFT_PhysicalDisk"
	const namespace = "root/microsoft/windows/storage"

	var physicalDisks []MSFT_PhysicalDisk

	if err := wmi.QueryNamespace(query, &physicalDisks, namespace); err != nil {
		return nil, err
	}

	return physicalDisks, nil
}

func (response *ListDrivesResponse) load() error {

	diskDrives, err := getDiskDrives()

	if err != nil {
		return err
	}

	disks := make([]*Disk, 0)

	for _, diskDrive := range diskDrives {

		disk := &Disk{
			Name:         strings.TrimSpace(*diskDrive.Name),
			Model:        strings.TrimSpace(*diskDrive.Model),
			SerialNumber: strings.TrimSpace(*diskDrive.SerialNumber),
		}

		disks = append(disks, disk)
	}

	response.Disks = disks

	return nil
}
