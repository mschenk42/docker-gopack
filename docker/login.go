package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// Login into Docker repo
type Login struct {
	User     string
	Password string
	Host     string

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (l Login) Run(runActions ...action.Enum) gopack.ActionRunStatus {
	l.setDefaults()
	return l.RunActions(&l, l.registerActions(), runActions)
}

func (l Login) registerActions() action.Methods {
	return action.Methods{
		action.Run: l.run,
	}
}

func (l *Login) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (l Login) String() string {
	return fmt.Sprintf("login %s %s", l.User, l.Host)
}

func (l Login) run() (bool, error) {
	t := task.Command{
		Name:   "docker",
		Args:   []string{"login", "-u", l.User, "-p", l.Password, l.Host},
		Stream: true,
	}
	return t.Run(action.Run)[action.Run], nil
}
