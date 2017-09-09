package docker

import (
	"fmt"
	"runtime"

	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
	"github.com/shirou/gopsutil/host"
)

func (s Service) install() (bool, error) {
	platform, _, _, err := host.PlatformInformation()
	if err != nil {
		return false, err
	}
	switch platform {
	case "centos":
		return s.install_centos()
	default:
		return false, fmt.Errorf("install not implemented for %s %s", runtime.GOOS, platform)
	}
	return false, nil
}

func (s Service) install_centos() (bool, error) {
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
