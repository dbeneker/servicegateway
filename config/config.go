package config

type Configuration struct {
	Applications map[string]Application `json:"applications"`
	RateLimiting RateLimiting `json:"rate_limiting"`
	Authentication GlobalAuth `json:"authentication"`
}

type Application struct {
	Routing Routing `json:"routing"`
	Backend Backend `json:"backend"`
	Auth ApplicationAuth `json:"auth"`
	Caching Caching `json:"caching"`
	RateLimiting bool `json:"rate_limiting"`
}

type Routing struct {
	Type string `json:"type"`
	Path string `json:"path"`
	Patterns map[string]string `json:"patterns"`
	Hostname string `json:"hostname"`
}

type Backend struct {
	Url string `json:"url"`
}

type ApplicationAuth struct {
	Disable bool `json:"disable"`
}

type GlobalAuth struct {
	Mode string `json:"mode"`
	StorageConfig StorageAuthConfig `json:"storage"`
	GraphicalConfig GraphicalAuthConfig `json:"graphical"`
	VerificationKey string `json:"verification_key"`
	VerificationKeyUrl string `json:"verification_key_url"`
	KeyCacheTtl string `json:"key_cache_ttl"`
}

type StorageAuthConfig struct {
	Mode string `json:"mode"`
	Name string `json:"name"`
}

type GraphicalAuthConfig struct {
	LoginRoute string `json:"login_route"`
}

type Caching struct {
	Enabled bool `json:"enabled"`
	Ttl int `json:"ttl"`
	AutoFlush bool `json:"auto_flush"`
}

type RateLimiting struct {
	Burst int `json:"burst"`
	Window string `json:"window"`
	RequestsPerSecond int `json:"requests_per_second"`
}
