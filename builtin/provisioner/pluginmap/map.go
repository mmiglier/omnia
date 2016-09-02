package pluginmap

import (
	"fmt"

	"github.com/mmiglier/omnia/builtin/provisioner/docker"
	"github.com/mmiglier/omnia/provisioner"
)

var Map = map[string]provisioner.Provisioner{
	"omnia-provisioner-docker": &docker.Provisioner{},
}

func Load(name string) (provisioner.Provisioner, error) {

	var p provisioner.Provisioner
	if p = Map["omnia-provisioner-"+name]; p == nil {
		return nil, fmt.Errorf("No provisioner found with id %s", name)
	}

	return p, nil
}
