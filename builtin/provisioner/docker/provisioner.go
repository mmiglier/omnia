package docker

//go:generate go-bindata -pkg $GOPACKAGE -o data.go data/...

import (
	"log"
	"os"
	"os/exec"

	"github.com/mmiglier/omnia/helper"
	"github.com/mmiglier/omnia/provisioner"
	"github.com/pkg/errors"
)

type Provisioner struct {
}

func (Provisioner) Compile(provisionerData provisioner.ProvisionerData, destDir string) error {

	for toolName, toolAsset := range provisionerData.ToolsAssets {

		toolDestDir := destDir + "/" + toolName

		if err := os.MkdirAll(toolDestDir, 0755); err != nil {
			return err
		}

		if err := helper.ParseTemplate(toolAsset.SetupScript, toolAsset, toolDestDir+"/setup.sh"); err != nil {
			return errors.Wrapf(err, "Failed parsing setup script for tool %s", toolName)
		}

		if err := os.MkdirAll(toolDestDir+"/conf/common", 0755); err != nil {
			return err
		}
		if err := os.MkdirAll(toolDestDir+"/conf/metrics", 0755); err != nil {
			return err
		}

		if len(toolAsset.Conf.Common) > 0 {
			for fileName, file := range toolAsset.Conf.Common {
				if err := helper.ParseTemplate(file, toolAsset, toolDestDir+"/conf/common/"+fileName); err != nil {
					return errors.Wrapf(err, "Failed parsing configuration file %s for tool %s", fileName, toolName)
				}
			}
		}

		if len(toolAsset.Conf.Metrics) > 0 {
			for dirName, dir := range toolAsset.Conf.Metrics {
				if err := os.MkdirAll(toolDestDir+"/conf/metrics/"+dirName, 0755); err != nil {
					return err
				}
				for fileName, file := range dir {
					if err := helper.ParseTemplate(file, toolAsset,
						toolDestDir+"/conf/metrics/"+dirName+"/"+fileName); err != nil {
						return errors.Wrapf(err, "Failed parsing configuration file %s for metric %s for tool %s", fileName, dirName, toolName)
					}
				}
			}
		}

		if err := helper.ParseTemplate(toolAsset.RunScript, toolAsset, toolDestDir+"/run.sh"); err != nil {
			return errors.Wrapf(err, "Failed parsing run script for tool %s", toolName)
		}

		dockerfileTmpl, err := Asset("data/Dockerfile")
		if err != nil {
			return err
		}
		if err := helper.ParseTemplate(dockerfileTmpl, toolAsset, toolDestDir+"/Dockerfile"); err != nil {
			return errors.Wrapf(err, "Failed parsing Dockerfile template for tool %s", toolName)
		}
	}

	dockerComposeTmpl, err := Asset("data/docker-compose.yml")
	if err != nil {
		return err
	}

	if err := helper.ParseTemplate(dockerComposeTmpl, provisionerData, destDir+"/docker-compose.yml"); err != nil {
		return errors.Wrapf(err, "Failed parsing docker-compose.yml file")
	}

	return nil
}

func (Provisioner) Deploy() error {
	out, err := exec.Command("date").Output()
	if err != nil {
		return errors.Wrapf(err, "Failed to execute date")
	}
	log.Printf("The date is %s", out)
	return nil
}
