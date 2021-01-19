package disk

type Disk struct {
	Name            string
	Model           string
	SerialNumber    string
	FirmwareVersion string
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
