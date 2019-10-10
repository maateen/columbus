package main

import (
	"github.com/maateen/columbus/internal/consul"
	"github.com/maateen/columbus/internal/docker"
)

func main() {
	services := docker.DiscoverServices()
	consul.RegisterServices(services)
	// fmt.Println(services)
}
