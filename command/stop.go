package command

import (
	"os"

	provisioners "github.com/mmiglier/omnia/builtin/provisioner/pluginmap"
	"github.com/mmiglier/omnia/omniafile"
	"github.com/pkg/errors"
)

func Stop(omniafileName string, omniaDir string) error {
	projectDir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "Failed to get working directory")
	}
	if err := os.Chdir(omniaDir); err != nil {
		return errors.Wrapf(err, "Failed to change directory to %s", omniaDir)
	}
	ofile := omniafile.Omniafile{}
	if err := ofile.Load(omniafileName); err != nil {
		return errors.Wrapf(err, "Failed to load file %s", omniafileName)
	}

	prov, err := provisioners.Load(ofile.Provisioner.Name)
	if err != nil {
		return errors.Wrapf(err, "Failed to load provisioner %s", ofile.Provisioner.Name)
	}

	if err = prov.Stop(); err != nil {
		return errors.Wrap(err, "Failed to stop")
	}

	if err := os.Chdir(projectDir); err != nil {
		return errors.Wrapf(err, "Failed to change directory to %s", projectDir)
	}

	return nil
}
