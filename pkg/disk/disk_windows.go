package disk

import (
	"errors"
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

type MsftMediaType uint16

const (
	MSFT_MEDIA_TYPE_UNSPECIFIED = 0
	MSFT_MEDIA_TYPE_HDD         = 1
	MSFT_MEDIA_TYPE_SSD         = 2
	MSFT_MEDIA_TYPE_SCM         = 3
)

type BusType uint16

const (
	BUS_TYPE_UNKNOWN            = 0
	BUS_TYPE_SCSI               = 1
	BUS_TYPE_ATAPI              = 2
	BUS_TYPE_ATA                = 3
	BUS_TYPE_1394               = 4
	BUS_TYPE_SSA                = 5
	BUS_TYPE_FIBRE_CHANNEL      = 6
	BUS_TYPE_USB                = 7
	BUS_TYPE_RAID               = 8
	BUS_TYPE_ISCSI              = 9
	BUS_TYPE_SAS                = 10
	BUS_TYPE_SATA               = 11
	BUS_TYPE_SD                 = 12
	BUS_TYPE_MMC                = 13
	BUS_TYPE_MAX                = 14
	BUS_TYPE_VIRTUAL            = 15
	BUS_TYPE_STORAGE_SPACES     = 16
	BUS_TYPE_NVME               = 17
	BUS_TYPE_MICROSOFT_RESERVED = 18
)

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

	physicalDrives, err := getPhysicalDisks()

	if err != nil {
		return err
	}

	disks := make([]*Disk, 0)

	for _, diskDrive := range diskDrives {

		// Find the physical drive using the serial number, is there are safer way to do this?
		physicalDisk := findPhysicalDrive(physicalDrives, *diskDrive.SerialNumber)

		if physicalDisk == nil {
			return errors.New("Failed to find physical drive with serial number " + *diskDrive.SerialNumber)
		}

		disk := &Disk{
			Name:            strings.TrimSpace(*diskDrive.Name),
			Model:           strings.TrimSpace(*diskDrive.Model),
			SerialNumber:    strings.TrimSpace(*diskDrive.SerialNumber),
			FirmwareVersion: strings.TrimSpace(*physicalDisk.FirmwareVersion),
			MediaType:       toMediaType(physicalDisk),
			InterfaceType:   toInterfaceType(physicalDisk),
		}

		disks = append(disks, disk)
	}

	response.Disks = disks

	return nil
}

// NOTE: Returning a pointer to a slice element seems wrong, is it?
func findPhysicalDrive(disks []MSFT_PhysicalDisk, serial string) *MSFT_PhysicalDisk {
	for _, disk := range disks {
		if *disk.SerialNumber == serial {
			return &disk
		}
	}

	// NOTE: Should we be returning a nullable pointer, or using an error tuple?
	return nil
}

func toInterfaceType(physicalDisk *MSFT_PhysicalDisk) InterfaceType {
	// NOTE: This should probably be an assert?
	if physicalDisk == nil {
		return INTERFACE_TYPE_UNKNOWN
	}

	switch *physicalDisk.BusType {
	case BUS_TYPE_NVME:
		return INTERFACE_TYPE_NVME
	case BUS_TYPE_ATA:
		return INTERFACE_TYPE_ATA
	case BUS_TYPE_SATA:
		return INTERFACE_TYPE_ATA
	default:
		return INTERFACE_TYPE_SCSI // HACK: Treat everything else as SCSI
	}
}

func toMediaType(physicalDisk *MSFT_PhysicalDisk) MediaType {
	// NOTE: This should probably be an assert
	if physicalDisk == nil {
		return MEDIA_TYPE_UNKNOWN
	}

	// The physical disk structure defines PCIe drives as "SSD"s. Explicitly
	// define them as NVMe drives instead.
	if *physicalDisk.BusType == BUS_TYPE_NVME {
		return MEDIA_TYPE_NVME
	}

	switch *physicalDisk.MediaType {
	case MSFT_MEDIA_TYPE_UNSPECIFIED:
		return MEDIA_TYPE_UNKNOWN
	case MSFT_MEDIA_TYPE_HDD:
		return MEDIA_TYPE_HDD
	case MSFT_MEDIA_TYPE_SSD:
		return MEDIA_TYPE_SSD
	case MSFT_MEDIA_TYPE_SCM: // Storage class memory (SCM) is it worth handling?
		return MEDIA_TYPE_UNKNOWN
	}

	return MEDIA_TYPE_UNKNOWN
}
