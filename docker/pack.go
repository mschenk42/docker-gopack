package docker

import (
	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
)

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
		},
	}
	pack.Run(props)
}

func run(pack *gopack.Pack) {
}

func install(pack *gopack.Pack) {
	Docker{
		Version:  pack.Props.Str("docker.version"),
		EdgeRepo: pack.Props.Bool("docker.enable-edge-repo"),
	}.Run(action.Install)
}
