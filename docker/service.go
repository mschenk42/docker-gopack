package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/systemd-gopack/systemd"
)

// Service installs and configures the docker service for a host
type Service struct {
	Version  string
	EdgeRepo bool

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (s Service) Run(runActions ...action.Name) gopack.ActionRunStatus {
	s.setDefaults()
	return s.RunActions(&s, s.registerActions(), runActions)
}

func (s Service) registerActions() action.Funcs {
	return action.Funcs{
		action.Install: s.install,
		action.Enable:  s.enable,
		action.Start:   s.start,
	}
}

func (s *Service) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (s Service) String() string {
	return fmt.Sprintf("docker %s edge %t", s.Version, s.EdgeRepo)
}

func (s Service) start() (bool, error) {
	status := systemd.SystemCtl{Service: "docker"}.Run(action.Start)
	return status[action.Start], nil
}

func (s Service) enable() (bool, error) {
	status := systemd.SystemCtl{Service: "docker"}.Run(action.Enable)
	return status[action.Enable], nil
}
