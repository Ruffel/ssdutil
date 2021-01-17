// +build !windows

package disk

import (
	"errors"
	"runtime"
)

func (response *ListDrivesResponse) load() error {
	return errors.New("Unimplemented function ListDrives.load on platform " + runtime.GOOS)
}
