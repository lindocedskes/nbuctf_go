package config

type K3s struct {
	K3sMaster string `mapstructure:"k3s-master" json:"k3s-master" yaml:"k3s-master"`
	K3sSlave  string `mapstructure:"k3s-slave" json:"k3s-slave" yaml:"k3s-slave"`
}
