package disk

type InterfaceType int

const (
	INTERFACE_TYPE_UNKNOWN InterfaceType = iota
	INTERFACE_TYPE_ATA                   = iota
	INTERFACE_TYPE_SCSI                  = iota
	INTERFACE_TYPE_NVME                  = iota
)

func (it InterfaceType) String() string {
	switch it {
	case INTERFACE_TYPE_UNKNOWN:
		return "Unknown"
	case INTERFACE_TYPE_ATA:
		return "SATA"
	case INTERFACE_TYPE_SCSI:
		return "SCSI"
	case INTERFACE_TYPE_NVME:
		return "NVMe"
	}

	return "Unknown"
}

type MediaType int

const (
	MEDIA_TYPE_UNKNOWN MediaType = iota
	MEDIA_TYPE_HDD               = iota
	MEDIA_TYPE_SSD               = iota
	MEDIA_TYPE_NVME              = iota
)

func (mt MediaType) String() string {
	switch mt {
	case MEDIA_TYPE_UNKNOWN:
		return "Unknown"
	case MEDIA_TYPE_HDD:
		return "HDD"
	case MEDIA_TYPE_SSD:
		return "SSD"
	case MEDIA_TYPE_NVME:
		return "NVMe"
	}

	return "Unknown"
}

type Disk struct {
	Name            string
	Model           string
	SerialNumber    string
	FirmwareVersion string
	InterfaceType   InterfaceType
	MediaType       MediaType
}

type ListDrivesResponse struct {
	Disks []*Disk
}

func ListDrives() (*ListDrivesResponse, error) {
	response := &ListDrivesResponse{}

	if err := response.load(); err != nil {
		return nil, err
	}

	return response, nil
}
