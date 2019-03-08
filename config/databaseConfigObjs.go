package config

// DatabaseConfigObj Config for InfluxDB
type DatabaseConfigObj struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBIP     string `json:"dbip"`
}
