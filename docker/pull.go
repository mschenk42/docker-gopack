package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// Pull ...
type Pull struct {
	Host string
	Repo string
	Tag  string

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (p Pull) Run(runActions ...action.Enum) gopack.ActionRunStatus {
	p.setDefaults()
	return p.RunActions(&p, p.registerActions(), runActions)
}

func (p Pull) registerActions() action.Methods {
	return action.Methods{
		action.Run: p.run,
	}
}

func (p *Pull) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (p Pull) String() string {
	return fmt.Sprintf("pull %s %s %s", p.Host, p.Repo, p.Tag)
}

func (p Pull) run() (bool, error) {
	t := task.Command{
		Name:   "docker",
		Args:   []string{"pull", fmt.Sprintf("%s/%s:%s", p.Host, p.Repo, p.Tag)},
		Stream: true,
	}
	return t.Run(action.Run)[action.Run], nil
}