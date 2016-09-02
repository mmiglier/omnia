package provisioner

type Provisioner interface {
	Compile(provisionerData ProvisionerData, destDir string) error
	Deploy() error
}
