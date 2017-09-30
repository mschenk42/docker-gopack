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
func (p Pull) Run(runActions ...action.Name) gopack.ActionRunStatus {
	p.setDefaults()
	return p.RunActions(&p, p.registerActions(), runActions)
}

func (p Pull) registerActions() action.Funcs {
	return action.Funcs{
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
	args := []string{"pull"}
	switch p.Host {
	case "":
		args = append(args, fmt.Sprintf("%s:%s", p.Repo, p.Tag))
	default:
		args = append(args, fmt.Sprintf("%s/%s:%s", p.Host, p.Repo, p.Tag))
	}
	t := task.Command{
		Name:   "docker",
		Args:   args,
		Stream: true,
	}
	return t.Run(action.Run)[action.Run], nil
}
