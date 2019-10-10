package consul

import (
	"fmt"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/maateen/columbus/config"
	"github.com/maateen/columbus/internal/docker"
)

func getPreparedServiceList(cfg *config.Config, key string, serviceList docker.ServiceList) map[string]string {
	preparedServiceList := make(map[string]string)

	for _, service := range serviceList {
		preparedServiceList["traefik/backends/"+key+"/servers/server"+strconv.Itoa(cfg.Node.Weight)+"/url"] = cfg.Node.Hostname + cfg.Node.Port
		preparedServiceList["traefik/backends/"+key+"/servers/server"+strconv.Itoa(cfg.Node.Weight)+"/weight"] = strconv.Itoa(cfg.Node.Weight)
		preparedServiceList["traefik/frontends/"+key+"/backend"] = service.Labels["traefik.backend"]
		preparedServiceList["traefik/frontends/"+key+"/entrypoints"] = cfg.Node.Scheme
		break
	}
	return preparedServiceList
}

func registerThisServiceList(kv *consulapi.KV, preparedServiceList map[string]string) {
	for k, v := range preparedServiceList {
		d := &consulapi.KVPair{Key: k, Value: []byte(v)}
		_, err := kv.Put(d, nil)

		if err != nil {
			panic(err)
		}
	}
}

// RegisterServices will store key-value pairs in consul
func RegisterServices(services map[string]docker.ServiceList) {
	cfg := config.GetConfig()
	if cfg.Consul.Enabled == true {
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = cfg.Consul.Hostname + ":" + cfg.Consul.Port
		consul, err := consulapi.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}

		kv := consul.KV()
		for key, serviceList := range services {
			preparedServiceList := getPreparedServiceList(cfg, key, serviceList)
			registerThisServiceList(kv, preparedServiceList)
		}
	} else {
		fmt.Println("Consul is not enabled.")
	}
}
