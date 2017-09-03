package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// Docker ...
type Docker struct {
	Version  string
	EdgeRepo bool

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (d Docker) Run(runActions ...action.Enum) gopack.ActionRunStatus {
	d.setDefaults()
	return d.RunActions(&d, d.registerActions(), runActions)
}

func (d Docker) registerActions() action.Methods {
	return action.Methods{
		action.Install: d.install,
	}
}

func (d *Docker) setDefaults() {
	if d.Version == "" {
		d.Version = "default value"
	}
}

// String returns a string which identifies the task with it's property values
func (d Docker) String() string {
	return fmt.Sprintf("docker %s edge: %t", d.Version, d.EdgeRepo)
}

func (d Docker) install() (bool, error) {
	const dockerBreadcrumb = "/var/run/docker-remove"

	setBreadcrumb := task.Command{
		Name: "touch",
		Args: []string{dockerBreadcrumb},
	}

	removeDocker := task.Command{
		Name:   "yum",
		Args:   []string{"remove", "-y", "docker", "docker-common", "docker-selinux", "docker-engine"},
		Stream: true,
	}
	removeDocker.SetNotIf(func() (bool, error) { _, x, err := task.Fexists(dockerBreadcrumb); return x, err })
	removeDocker.SetNotify(setBreadcrumb, action.Run, action.Run, true)
	removeDocker.Run(action.Run)

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
	enableEdgeRepo.SetOnlyIf(func() (bool, error) { return d.EdgeRepo, nil })
	enableEdgeRepo.Run(action.Run)

	task.Command{
		Name:   "yum",
		Args:   []string{"makecache", "fast"},
		Stream: true,
	}.Run(action.Run)

	task.Command{
		Name:   "yum",
		Args:   []string{"install", "-y", d.Version},
		Stream: true,
	}.Run(action.Run)

	return true, nil
}
