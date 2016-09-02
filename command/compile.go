package command

import (
	"os"

	"github.com/ghodss/yaml"
	provisioners "github.com/mmiglier/omnia/builtin/provisioner/pluginmap"
	"github.com/mmiglier/omnia/builtin/tool"
	"github.com/mmiglier/omnia/omniafile"
	"github.com/mmiglier/omnia/provisioner"
)

const toolsBindataDir string = "builtin/tool/data"

func Compile(omniafileName string, omniaDir string) error {

	ofile := omniafile.Omniafile{}
	if err := ofile.Load(omniafileName); err != nil {
		return err
	}

	if err := ofile.SaveTo(omniaDir + "/" + omniafileName); err != nil {
		return err
	}

	prov, err := provisioners.Load(ofile.Provisioner.Name)
	if err != nil {
		return err
	}

	provisionerData, err := buildProvisionerData(ofile)
	if err != nil {
		return err
	}

	if err := prov.Compile(provisionerData, omniaDir); err != nil {
		return err
	}

	return nil
}

func buildProvisionerData(ofile omniafile.Omniafile) (provisioner.ProvisionerData, error) {

	var provisionerData provisioner.ProvisionerData
	toolsAssets := make(map[string]provisioner.ToolAsset, len(ofile.Tools))
	linkedToolsMap := make(map[string]map[string]bool)

	for toolName, toolDesc := range ofile.Tools {

		currToolBindataDir := toolsBindataDir + "/" + toolName

		var portsFile []byte
		var ports []int

		portsFile, err := tool.Asset(currToolBindataDir + "/ports.yml")
		if !os.IsNotExist(err) {
			if err != nil {
				return provisionerData, err
			}
			if err = yaml.Unmarshal(portsFile, &ports); err != nil {
				return provisionerData, err
			}
		}

		setupScript, err := tool.Asset(currToolBindataDir + "/setup.sh")
		if err != nil {
			return provisionerData, err
		}

		commonConf := make(map[string][]byte)
		commonConfFilesNames, err := tool.AssetDir(currToolBindataDir + "/conf/common")
		if !os.IsNotExist(err) {
			for _, commonConfFileName := range commonConfFilesNames {
				curr, err := tool.Asset(currToolBindataDir + "/conf/common/" + commonConfFileName)
				if err != nil {
					return provisionerData, err
				}
				commonConf[commonConfFileName] = curr
			}
		}

		metricsConf := make(map[string]map[string][]byte, len(ofile.DefaultMetrics))
		for _, metricName := range ofile.DefaultMetrics {
			currMetricConfFilesNames, err := tool.AssetDir(currToolBindataDir + "/conf/metrics/" + metricName)
			if !os.IsNotExist(err) {
				metricsConf[metricName] = make(map[string][]byte, len(currMetricConfFilesNames))
				for _, currMetricConfFileName := range currMetricConfFilesNames {
					curr, err := tool.Asset(currToolBindataDir + "/conf/metrics/" + metricName + "/" + currMetricConfFileName)
					if err != nil {
						return provisionerData, err
					}
					metricsConf[metricName][currMetricConfFileName] = curr
				}
			}
		}

		conf := provisioner.ToolConf{
			Common:  commonConf,
			Metrics: metricsConf,
		}

		var runScript []byte
		runScript, err = tool.Asset(currToolBindataDir + "/run.sh")
		if err != nil {
			return provisionerData, err
		}

		toolsAssets[toolName] = provisioner.ToolAsset{
			Name:        toolName,
			IsAgent:     toolDesc.IsAgent,
			Ports:       ports,
			Links:       toolDesc.Links,
			SetupScript: setupScript,
			Conf:        conf,
			RunScript:   runScript,
		}

		for _, link := range toolDesc.Links {
			if linkedToolsMap[link] == nil {
				linkedToolsMap[link] = make(map[string]bool)
			}
			linkedToolsMap[link][toolName] = true
		}
	}

	for toolName, toolAsset := range toolsAssets {
		for linked := range linkedToolsMap[toolName] {
			toolAsset.Linked = append(toolAsset.Linked, linked)
		}
	}

	provisionerData = provisioner.ProvisionerData{
		Args:        ofile.Provisioner.Args,
		ToolsAssets: toolsAssets,
	}

	return provisionerData, nil
}
