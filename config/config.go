package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type mysql struct {
	Debug       bool   `yaml:"debug"`
	Port        int16  `yaml:"port"`
	MaxIdleConn int16  `yaml:"maxIdleConns"`
	MaxOpenConn int16  `yaml:"maxOpenConns"`
	UserName    string `yaml:"username"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	DataBase    string `yaml:"database"`
}

type qiniuCfg struct {
	AccessKey string `yaml:"accessKey"`
	Bucket    string `yaml:"bucket"`
	SecretKey string `yaml:"secretKey"`
}

type serverCfg struct {
	Port       int16  `yaml:"port"`
	ServerName string `yaml:"servername"`
	DataBase   string `yaml:"database"`
}

// 配置
type Config struct {
	Mysql mysql `yaml:"mysql"`
	// Register  bool  `yaml:"register"`
	ServerCfg   map[string]serverCfg
	serNameLock *sync.Mutex
	Qiniu       qiniuCfg
}

var cfg *Config

func init() {
	cfg = &Config{
		ServerCfg:   make(map[string]serverCfg),
		serNameLock: &sync.Mutex{},
	}
	loadDefault("default")
}

func getFileName() string {
	file := "config_prev.yaml"
	env := os.Getenv("PRODUCTION")
	if env == "true" {
		file = "config_production.yaml"
	}

	// file = "D:/work_space/v5-learning/merge/src/robot-server/config_test.yaml"
	// file = os.Getenv("GOPATH") + "/" + file
	file = "/home/m/go/src/robot-server/config_test.yaml"
	return file
}

func loadDefault(cfgName string) {
	file := getFileName()
	temp := readAll(file)
	if cfg.Mysql.Host == "" {
		cfg.Mysql = temp["default"].Mysql
	}

	if cfg.Qiniu.AccessKey == "" {
		cfg.Qiniu = temp["default"].Qiniu
	}

}
func Get() *Config {
	return cfg
}

func LoadService(serverName string) {
	file := getFileName()
	temp := readServerCfg(file, serverName)
	_, ok := cfg.ServerCfg[serverName]

	if !ok {
		cfg.ServerCfg[serverName] = temp[serverName]
	}

	printConfig()
}

func readServerCfg(file, serverName string) map[string]serverCfg {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var temp map[string]serverCfg
	err = yaml.Unmarshal(content, &temp)
	if err != nil {
		fmt.Println("yaml unmarshal:", err)
		os.Exit(-1)
	}

	return temp

}

func readAll(file string) map[string]Config {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var temp map[string]Config
	err = yaml.Unmarshal(content, &temp)
	if err != nil {
		fmt.Println("yaml unmarshal:", err)
		os.Exit(-1)
	}
	return temp
}

func printConfig() {
	fmt.Println(cfg)
}
