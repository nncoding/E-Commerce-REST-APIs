package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv fail: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout Failed: %v", err)
				}
				// เปลี่ยนจาก ms. เป็น s.
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRTIE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load write timeout Failed: %v", err)
				}
				// เปลี่ยนจาก ms. เป็น s.
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load body limit failed: %v", err)
				}
				return b
			}(),
			fileLimit: func() int {
				f, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("load file limit failed: %v", err)
				}
				return f
			}(),
			gcpbuket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load db port failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxXonnection: func() int {
				m, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load max connection failed: %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secertKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				f, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load access expires failed: %v", err)
				}
				return f
			}(),
			refreshExpiresAt: func() int {
				f, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refresh expires failed: %v", err)
				}
				return f
			}(),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDBConfig
	Jwt() IJWTConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GPCBuket() string
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int // bytes
	fileLimit    int // bytes
	gcpbuket     string
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GPCBuket() string            { return a.gcpbuket }

type IDBConfig interface {
	Url() string
	MaxOpenConns() int
}

type db struct {
	host          string
	port          int
	protocol      string
	username      string
	password      string
	database      string
	sslMode       string
	maxXonnection int
}

func (c *config) Db() IDBConfig {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host,
		d.port,
		d.username,
		d.password,
		d.database,
		d.sslMode,
	)
}
func (d *db) MaxOpenConns() int { return d.maxXonnection }

type IJWTConfig interface {
	SecretKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

type jwt struct {
	adminKey         string
	secertKey        string
	apiKey           string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (c *config) Jwt() IJWTConfig {
	return c.jwt
}

func (j *jwt) SecretKey() []byte          { return []byte(j.secertKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
