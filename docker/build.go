package docker

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// Build ...
type Build struct {
	Image        string
	Tag          string
	DockerFile   string
	DockerIgnore string
	Timeout      time.Duration

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (b Build) Run(runActions ...action.Name) gopack.ActionRunStatus {
	b.setDefaults()
	return b.RunActions(&b, b.registerActions(), runActions)
}

func (b Build) registerActions() action.Funcs {
	return action.Funcs{
		action.Run: b.run,
	}
}

func (b *Build) setDefaults() {
	if b.Timeout == 0 {
		b.Timeout = time.Minute * 30
	}
}

// String returns a string which identifies the task with it's property values
func (b Build) String() string {
	return fmt.Sprintf("build %s %s", b.Image, b.Tag)
}

func (b Build) run() (bool, error) {
	args := []string{"build", "-t", b.Image}
	if b.Tag != "" {
		args[len(args)-1] = fmt.Sprintf("%s:%s", b.Image, b.Tag)
	}
	args = append(args, ".")

	if err := ioutil.WriteFile("Dockerfile", []byte(b.DockerFile), 0755); err != nil {
		return false, err
	}
	if err := ioutil.WriteFile(".dockerignore", []byte(b.DockerIgnore), 0755); err != nil {
		return false, err
	}

	build := task.Command{
		Name:    "docker",
		Args:    args,
		Stream:  true,
		Timeout: b.Timeout,
	}.Run(action.Run)

	return build[action.Run], nil
}
