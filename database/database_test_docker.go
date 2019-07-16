package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // "empty import" to register driver with database/sql
	"net/url"
	"os/exec"
	"strings"
	"testing"
	"time"
)

type PostgresDB struct {
	// Of form 'host:port'
	Host string

	// The container id assigned by docker
	cid string
}

type PostgresConfig struct {
	Password string
	Username string // defaults to "postgres"
	Database string // defaults to "username"
	Version  string // defaults to "latest"
}

func NewPostgresDB(c PostgresConfig) (*PostgresDB, error) {
	img := "postgres:latest"
	if c.Version == "" {
		img = "postgres:" + c.Version
	}

	// docker's run command has the nasty habbit of pulling images if you don't have them.
	// Warn user they need to pull the image, don't automatically pull for them.
	if exec.Command("docker", "inspect", img).Run() != nil {
		return nil, fmt.Errorf("db requires docker image %s, please pull or specify a different version", img)
	}

	// Running on port 0 instructs the operating system to pick an available port.
	dockerArgs := []string{"run", "-d", "-p", "127.0.0.1:0:5432"}
	envvars := map[string]string{
		"POSTGRES_PASSWORD": c.Password,
		"POSTGRES_USER":     c.Username,
		"POSTGRES_DB":       c.Database,
	}
	for key, val := range envvars {
		if val != "" {
			dockerArgs = append(dockerArgs, "-e", key+"="+val)
		}
	}
	dockerArgs = append(dockerArgs, img)

	// Start the docker container.
	out, err := exec.Command("docker", dockerArgs...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("docker run: %v: %s", err, out)
	}
	cid := strings.TrimSpace(string(out))
	db := &PostgresDB{cid: cid}

	db.Host, err = portMapping(cid, "5432/tcp")
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func portMapping(cid, containerPort string) (hostAddr string, err error) {
	out, err := exec.Command("docker", "inspect", cid).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("docker inspect: %v: %s", err, out)
	}

	// anonymous struct for unmarshalling JSON into
	var inspectResp []struct {
		NetworkSettings struct {
			Ports map[string][]struct {
				HostIp   string
				HostPort string
			}
		}
	}
	if err := json.Unmarshal(out, &inspectResp); err != nil {
		return "", fmt.Errorf("decoding docker inspect result failed: %v: %s", err, out)
	}
	if len(inspectResp) != 1 {
		return "", fmt.Errorf("expected one inspect result, got %d", len(inspectResp))
	}
	ports := inspectResp[0].NetworkSettings.Ports[containerPort]
	if len(ports) != 1 {
		return "", fmt.Errorf("expected one port mapping, got %d", len(ports))
	}
	return ports[0].HostIp + ":" + ports[0].HostPort, nil
}

// Close removes the container running the postgres database.
func (db *PostgresDB) Close() error {
	out, err := exec.Command("docker", "rm", "-f", db.cid).CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker rm: %v: %s", err, out)
	}
	return nil
}

type DBTest func(t *testing.T, conn *sql.DB)

// run test

func RunDBTest(t *testing.T, dbVersion string, test DBTest, config ConfigObj) {
	password := config.DBConfig.Password
	username := config.DBConfig.Username
	dbname := config.DBConfig.DBName
	c := PostgresConfig{password, username, dbname, dbVersion}

	// create a postgres container
	db, err := NewPostgresDB(c)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close() // destroy the postgres container after the test

	// create a connection URL
	// http://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
	connURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.Username, c.Password),
		Host:     db.Host,
		Path:     "/" + c.Database,
		RawQuery: "sslmode=disable",
	}

	// connect to database
	conn, err := sql.Open("postgres", connURL.String())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()

	// ping the database until it comes up
	timeout := time.Now().Add(time.Second * 20)
	for time.Now().Before(timeout) {
		if err = conn.Ping(); err == nil {
			// yay! we've connected to the database, time to run the test
			test(t, conn)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Errorf("failed to connect to database: %v", err)
}
