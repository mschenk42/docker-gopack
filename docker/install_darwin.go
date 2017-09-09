package docker

import (
	"fmt"
	"runtime"
)

func (s Service) install() (bool, error) {
	return false, fmt.Errorf("install not implemented for %s", runtime.GOOS)
}
