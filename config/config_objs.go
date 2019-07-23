package config

// DatabaseConfigObj Config for InfluxDB
type DBConfigObj struct {
	DBName      string `json:"dbname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ProdDBIP    string `json:"proddbip"`
	DevDBIP     string `json:"devdbip"`
	ProdDBPort  string `json:"proddbport"`
	DevDBPPort  string `json:"devdbport"`
	DBConfigStr string ""
}

// ServerConfigObj for general config
type ServerConfigObj struct {
	Port  string `json:"port"`
	IsDev bool   `json:"isdev"`
}
