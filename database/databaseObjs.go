package database

import (
	influx "github.com/influxdata/influxdb/client/v2"
)

// Future fields may include statistics???
type DBObj struct {
	DBClient influx.Client
}
