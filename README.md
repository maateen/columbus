
# Columbus

  

Columbus is a modern service discovery tool for [Docker](https://www.docker.com/) and [Traefik](https://traefik.io/), inspired from [Registrator](https://github.com/gliderlabs/registrator). It uses [Consul](https://www.consul.io/) as service registry.

  

Columbus automatically registers and deregisters services in Consul for any Docker container by inspecting containers as soon as they become healthy. Traefik listens to Consul and routes requests to available containers based on [labels](https://docs.traefik.io/providers/docker/).

  

## To-Do

  

This project is incomplete. I wrote this tool on my weekend to practice GoLang. Please contribute and complete the project if you can.

  

- [ ] Run the tool as a daemon

- [ ] Listen to Docker events

- [ ] Add support for more KV storage

- [ ] Minify labels

- [ ] Update Documentation

- [ ] Add architecture image