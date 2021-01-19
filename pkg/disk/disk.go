package disk

type InterfaceType int

const (
	INTERFACE_TYPE_UNKNOWN InterfaceType = iota
	INTERFACE_TYPE_ATA                   = iota
	INTERFACE_TYPE_SCSI                  = iota
	INTERFACE_TYPE_NVME                  = iota
)

type MediaType int

const (
	MEDIA_TYPE_UNKNOWN MediaType = iota
	MEDIA_TYPE_HDD               = iota
	MEDIA_TYPE_SSD               = iota
	MEDIA_TYPE_NVME              = iota
)

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
