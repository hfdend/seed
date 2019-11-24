package conf

import (
	"flag"
	"io/ioutil"
	"log"
	"seed/utils"
	"time"

	"gopkg.in/yaml.v3"
)

const TestMode = "test"

type MysqlModel struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DB              string        `yaml:"db"`
	Net             string        `yaml:"net"`
	Timeout         time.Duration `yaml:"timeout"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
}

type Model struct {
	Log struct {
		Dir string `yaml:"dir"`
	} `yaml:"log"`
	DataDir string `yaml:"data_dir"`
}

var (
	Config Model
)

func Init() {
	var (
		configFile string
		configBts  []byte
		err        error
	)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	flag.StringVar(&configFile, "f", "", "config file")
	flag.Parse()

	if configFile == "" {
		if configBts, err = utils.LRead("config.yml", 2); err != nil {
			log.Fatalln(err)
		}
	} else {
		if configBts, err = ioutil.ReadFile(configFile); err != nil {
			log.Fatalln(err)
		}
	}
	if err = yaml.Unmarshal(configBts, &Config); err != nil {
		log.Fatalln(err)
	}

	if configFile == "" {
		if configBts, err = utils.LRead("config.yml", 2); err != nil {
			log.Fatalln(err)
		}
	} else {
		if configBts, err = ioutil.ReadFile(configFile); err != nil {
			log.Fatalln(err)
		}
	}
}
