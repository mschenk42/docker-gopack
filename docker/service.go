package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
	"github.com/mschenk42/systemd-runpack/systemd"
)

// Service installs and configures the Service service for a host
type Service struct {
	Version  string
	EdgeRepo bool

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (s Service) Run(runActions ...action.Enum) gopack.ActionRunStatus {
	s.setDefaults()
	return s.RunActions(&s, s.registerActions(), runActions)
}

func (s Service) registerActions() action.Methods {
	return action.Methods{
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

func (s Service) install() (bool, error) {
	const dockerBreadcrumb = "/var/run/docker-installed"

	_, breadcrumb, err := task.Fexists(dockerBreadcrumb)
	if err != nil {
		return false, err
	}

	if !breadcrumb {
		task.Command{
			Name:   "yum",
			Args:   []string{"remove", "-y", "docker", "docker-common", "docker-selinux", "docker-engine"},
			Stream: true,
		}.Run(action.Run)

		task.Command{
			Name:   "yum",
			Args:   []string{"install", "-y", "yum-utils", "device-mapper-persistent-data", "lvm2"},
			Stream: true,
		}.Run(action.Run)

		task.Command{
			Name:   "yum-config-manager",
			Args:   []string{"--add-repo", "https://download.docker.com/linux/centos/docker-ce.repo"},
			Stream: true,
		}.Run(action.Run)

		enableEdgeRepo := task.Command{
			Name:   "yum-config-manager",
			Args:   []string{"--enable", "docker-ce-edge"},
			Stream: true,
		}
		enableEdgeRepo.SetOnlyIf(func() (bool, error) { return s.EdgeRepo, nil })
		enableEdgeRepo.Run(action.Run)

		task.Command{
			Name:   "yum",
			Args:   []string{"makecache", "fast"},
			Stream: true,
		}.Run(action.Run)

		task.Command{
			Name:   "yum",
			Args:   []string{"install", "-y", s.Version},
			Stream: true,
		}.Run(action.Run)

		task.Command{
			Name: "touch",
			Args: []string{dockerBreadcrumb},
		}.Run(action.Run)

		return true, nil
	}
	return false, nil
}

func (s Service) start() (bool, error) {
	status := systemd.SystemCtl{Service: "docker"}.Run(action.Start)
	return status[action.Start], nil
}

func (s Service) enable() (bool, error) {
	status := systemd.SystemCtl{Service: "docker"}.Run(action.Enable)
	return status[action.Enable], nil
}
