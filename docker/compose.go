package docker

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// Compose ...
type Compose struct {
	Path    string
	Version string

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (c Compose) Run(runActions ...action.Name) gopack.ActionRunStatus {
	c.setDefaults()
	return c.RunActions(&c, c.registerActions(), runActions)
}

func (c Compose) registerActions() action.Funcs {
	return action.Funcs{
		action.Install: c.install,
		action.Up:      c.up,
		action.Down:    c.down,
	}
}

func (c *Compose) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (c Compose) String() string {
	return fmt.Sprintf("compose %s %s", c.Path, c.Version)
}

func (c Compose) install() (bool, error) {
	// TODO: implement
	f := task.File{
		Path: c.Path,
		URL:  "",
	}
	f.SetNotIf(
		func() (bool, error) {
			_, exists, err := task.Fexists(f.Path)
			return exists, err
		},
	)
	return f.Run(action.Install)[action.Install], nil
}

func (c Compose) up() (bool, error) {
	// TODO: implement
	return true, nil
}

func (c Compose) down() (bool, error) {
	// TODO: implement
	return true, nil
}
