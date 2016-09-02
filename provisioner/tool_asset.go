package provisioner

type ProvisionerData struct {
	Args        map[string]string
	ToolsAssets map[string]ToolAsset
}

type ToolAsset struct {
	Name        string
	IsAgent     bool
	Ports       []int
	Links       []string
	Linked      []string
	SetupScript []byte
	Conf        ToolConf
	RunScript   []byte
}

type ToolConf struct {
	Common  map[string][]byte
	Metrics map[string]map[string][]byte
}
