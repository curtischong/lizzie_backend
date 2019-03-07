package database

// DatabaseConfigObj Config for InfluxDB
type DatabaseConfigObj struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBIP     string `http://10.8.0.2:8086`
}
