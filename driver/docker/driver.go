// Runtime driver implementation for Docker
package docker

import (
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/geovanisouza92/lambdac/driver"
	"github.com/geovanisouza92/lambdac/types"
)

func init() {
	driver.Register("docker", new(dockerDriver))
}

type dockerDriver struct {
	c *docker.Client
}

func (d *dockerDriver) Init(options []string) (err error) {
	optionsMap := map[string]string{}
	for _, o := range options {
		pair := strings.Split(o, "=")
		optionsMap[pair[0]] = pair[1]
	}

	host := optionsMap["DOCKER_HOST"]
	if host == "" {
		host = "unix:///var/run/docker.sock"
	}
	d.c, err = docker.NewClient(host)
	return
}

func (d *dockerDriver) Create(function types.Function, runtime types.Runtime) (id string, err error) {
	opts := docker.CreateContainerOptions{
		// Name: "",
		Config: &docker.Config{
			Memory:     int64(1024 * 1024 * function.Memory),
			Env:        []string{},
			Entrypoint: []string{}, // Must override?
			Cmd:        []string{runtime.Command},
			Image:      runtime.Image,
			// WorkingDir
			// TODO Labels
		},
		HostConfig: &docker.HostConfig{
		// RestartPolicy (runtime.Agent)
		},
	}

	if runtime.Agent {
		opts.HostConfig.RestartPolicy = docker.RestartPolicy{Name: "unless-stopped"}
	}

	container, err := d.c.CreateContainer(opts)
	if err != nil {
		return
	}
	id = container.ID

	return
}

func (d *dockerDriver) Start(id string) (err error) {
	// TODO
	return
}

func (d *dockerDriver) Stop(id string) (err error) {
	// TODO
	return
}

func (d *dockerDriver) Destroy(id string) (err error) {
	// TODO
	return
}
