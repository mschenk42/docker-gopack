package docker

import (
	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
)

var dockerTask Docker

// Run initializes the properties and runs the pack
func Run(props *gopack.Properties, actions []string) {
	pack := gopack.Pack{
		Name: "docker",
		Props: &gopack.Properties{
			"docker.version":          "docker-ce",
			"docker.enable-edge-repo": false,
		},
		Actions: actions,
		ActionMap: map[string]func(p *gopack.Pack){
			"default": run,
			"install": install,
			"start":   start,
		},
	}
	dockerTask = Docker{
		Version:  pack.Props.Str("docker.version"),
		EdgeRepo: pack.Props.Bool("docker.enable-edge-repo"),
	}
	pack.Run(props)
}

func run(pack *gopack.Pack) {
}

func install(pack *gopack.Pack) {
	dockerTask.Run(action.Install)
}

func start(pack *gopack.Pack) {
	dockerTask.Run(action.Enable, action.Start)
}
