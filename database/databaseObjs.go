package database

import (
	influx "github.com/influxdata/influxdb/client/v2"
)

type DBObj struct {
	DBConfig DBConfigObj
	DBClient influx.Client
}
