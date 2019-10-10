package docker

import (
	dockerclient "github.com/fsouza/go-dockerclient"
)

// Container is another struct
type Container struct {
	IPs    []string
	Ports  []dockerclient.APIPort
	Labels map[string]string
}

// ServiceList is a type alias
type ServiceList []Container

var (
	services map[string]ServiceList
)

func getRawServiceList() []dockerclient.APIContainers {
	client, err := dockerclient.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	listContainersOpts := dockerclient.ListContainersOptions{
		All: false,
		Filters: map[string][]string{
			"status": {"running"},
			"label":  {"traefik.enable=true"},
		},
	}

	serviceList, err := client.ListContainers(listContainersOpts)
	if err != nil {
		panic(err)
	}

	return serviceList
}

// DiscoverServices returns Docker services.
func DiscoverServices() map[string]ServiceList {
	services = make(map[string]ServiceList)
	rawServiceList := getRawServiceList()
	for _, rawService := range rawServiceList {
		ips := make([]string, 0)
		for _, x := range rawService.Networks.Networks {
			ips = append(ips, x.IPAddress)
		}
		key := rawService.Labels["com.docker.compose.config-hash"]
		service := Container{ips, rawService.Ports, rawService.Labels}
		services[key] = append(services[key], service)
	}

	return services
}
