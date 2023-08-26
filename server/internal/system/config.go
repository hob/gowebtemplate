package system

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Profile string

const (
	Prod Profile = "prod"
	Dev  Profile = "dev"
)

var profile = Prod

type Config struct {
	CsrfTokenKey  string `yaml:"csrfTokenKey",env:"CSRF_TOKEN_KEY"`
	HmacKey       string `yaml:"hmacKey",env:"HMAC_KEY"`
	MySqlUsername string `yaml:"mysqlUserName",env:"MYSQL_USER_NAME"`
	MySqlPassword string `yaml:"mysqlPassword",env:"MYSQL_PASSWORD"`
	BaseURL       string `yaml:"baseURL",env:"BASE_URL"`
	MySqlHost     string `yaml:"mysqlHost",env:"MYSQL_HOST"`
	MySqlPath     string `yaml:"mysqlPath",env:"MYSQL_PATH"`
	RedisHost     string `yaml:"redisHost",env:"REDIS_HOST"`
	Profile       Profile
}

func GetConfig() Config {
	var config Config
	//Read from env 1st, then from yaml so that dev values override prod values
	config.CsrfTokenKey = os.Getenv("CSRF_TOKEN_KEY")
	config.HmacKey = os.Getenv("HMAC_KEY")
	config.MySqlUsername = os.Getenv("MYSQL_USER_NAME")
	config.MySqlPassword = os.Getenv("MYSQL_PASSWORD")
	config.BaseURL = os.Getenv("BASE_URL")
	config.MySqlHost = os.Getenv("MYSQL_HOST")
	config.MySqlPath = os.Getenv("MYSQL_PATH")
	config.RedisHost = os.Getenv("REDIS_HOST")

	var yamlFile *os.File
	var err error
	if profile == Prod {
		logrus.Info("Loading prod profile")
		yamlFile, err = os.Open("config.yaml")
	} else {
		logrus.Info("Loading dev profile")
		yamlFile, err = os.Open("config.dev.yaml")
	}
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(yamlFile).Decode(&config); err != nil {
		panic(err)
	}
	config.Profile = profile
	if profile == Dev {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	return config
}
