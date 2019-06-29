package config

// DatabaseConfigObj Config for InfluxDB
type DatabaseConfigObj struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	ProdDBIP string `json:"proddbip"`
	DevDBIP  string `json:"devdbip"`
	IsDev    bool   `json:"isdev`
}

// ServerConfigObj for general config
type ServerConfigObj struct {
	Port string `json:"port"`
}
