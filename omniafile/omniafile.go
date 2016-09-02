package omniafile

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

type Omniafile struct {
	Provisioner       Provisioner     `json:"provisioner"`
	Tools             map[string]Tool `json:"tools"`
	MonitoredServices []string        `json:"monitored_services"`
	DefaultMetrics    []string        `json:"default_metrics"`
}

type Tool struct {
	IsAgent bool     `json:"is_agent"`
	Links   []string `json:"links"`
}

type Provisioner struct {
	Name string            `json:name`
	Args map[string]string `json:"args"`
}

func (ofile *Omniafile) Load(omniafileName string) error {
	file, err := ioutil.ReadFile(omniafileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &ofile)
	if err != nil {
		return err
	}
	return nil
}

func (ofile *Omniafile) SaveTo(destPath string) error {

	file, err := yaml.Marshal(ofile)
	if err != nil {
		return errors.Wrapf(err, "Failed to marshal omniafile to %s", destPath)
	}

	if err = ioutil.WriteFile(destPath, file, 0644); err != nil {
		return errors.Wrapf(err, "Error while writing omniafile to %s", destPath)
	}

	return nil
}
