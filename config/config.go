package config

type Config struct{}

type ClientsConf struct {
	HarnessBaseUrl         string `envconfig:"HARNESS_BASE_URL" default:"https://app.harness.io/gateway"`
	HarnessPipelineSvcPath string `envconfig:"HARNESS_PIPELINE_SVC_PATH" default:"/pipeline"`
}
