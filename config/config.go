package config

var (
	Config *ProjectConfig
	NodeId string
)

func init() {
	Config = &ProjectConfig{}
}

type ProjectConfig struct {
	HttpServer  HttpServerConfig  `json:"http_server" yaml:"http_server" mapstructure:"http_server"`
	Application ApplicationConfig `json:"application" yaml:"application" mapstructure:"application"`
}

type ApplicationConfig struct {
	Mode               string             `json:"mod" yaml:"mod" mapstructure:"mode"`
	NodeId             string             `json:"node_id" yaml:"node_id" mapstructure:"node_id"`
	ClockSequence      int                `json:"clock_sequence" yaml:"clock_sequence" mapstructure:"clock_sequence"`
	ObjectClientConfig ObjectClientConfig `json:"object_client" yaml:"object_client" mapstructure:"object_client"`
}

type HttpServerConfig struct {
	Port string `json:"port" yaml:"port" mapstructure:"port"`
}

type ObjectClientConfig struct {
	Enable bool   `json:"enable" yaml:"enable"`
	Host   string `json:"host" yaml:"host" mapstructure:"host"`
}
