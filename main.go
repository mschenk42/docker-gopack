package main

import (
	"github.com/mschenk42/docker-runpack/docker"
	"github.com/mschenk42/gopack"
)

func main() {
	docker.Run(gopack.LoadProperties())
}
