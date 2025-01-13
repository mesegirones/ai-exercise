package config

type ConfigURL struct {
	service string
}

func (g ConfigURL) GetService() string {
	return g.service
}

type ConfigRest struct {
	prod    bool
	local   bool
	service string
	port    string
}

func (g ConfigRest) IsProd() bool {
	return g.prod
}

func (g ConfigRest) IsLocal() bool {
	return g.local
}

func (g ConfigRest) GetService() string {
	return g.service
}

func (g ConfigRest) GetPort() string {
	return g.port
}

type ConfigHealth struct {
	service string
}

func (g ConfigHealth) GetService() string {
	return g.service
}

type ConfigOpenAI struct {
	apiKey string
}

func (g ConfigOpenAI) GetAPIKEy() string {
	return g.apiKey
}

type Config struct {
	restConfig   ConfigRest
	urlConfig    ConfigURL
	healthConfig ConfigHealth
	openaiConfig ConfigOpenAI
}

func (g Config) GetRestConfig() ConfigRest {
	return g.restConfig
}

func (g Config) GetUrlConfig() ConfigURL {
	return g.urlConfig
}

func (g Config) GetHealthConfig() ConfigHealth {
	return g.healthConfig
}

func (g Config) GetOpenAIConfig() ConfigOpenAI {
	return g.openaiConfig
}
