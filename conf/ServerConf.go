package conf 

import (
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"log"
	"fmt"
)

type ServerYaml struct {
	Conf Server `yaml:"server"` 
	Callback_url JxhzSt `yaml:"callback_url"` 
}

type Server struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
	LogPath string `yaml:"logPath"`
}

//jxhz配置
type JxhzSt struct {
	Jxhz string `yaml:"jxhz"`
}

type Conf struct {
	file []byte
	ServerAddr string
	ServerConf *ServerYaml
}

func New() *Conf{
	yamlFile, err := ioutil.ReadFile("./conf/Serverconf.yaml")
	if err != nil {
		log.Fatal("read yaml file error:", err)
	}

	conf := &Conf{
		file: []byte(yamlFile),
	}
    
	return conf
}


func (c *Conf) GetAddr() string{
	if c.ServerAddr != "" {
		return c.ServerAddr
	}

	serverConf := new(ServerYaml)
	err :=  yaml.Unmarshal(c.file, &serverConf)
	if err != nil {
		log.Fatal("yaml unmarshal err:", err)
	}

	c.ServerAddr = fmt.Sprintf("%s:%s", serverConf.Conf.Addr, serverConf.Conf.Port)
	c.ServerConf = serverConf

	return c.ServerAddr
}
