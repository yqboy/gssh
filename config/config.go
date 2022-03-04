package config

import (
	"gssh/utils"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	User string `yaml:"User"`
	Pwd  string `yaml:"Pwd"`
}

func ReadHost(host string) (Config, error) {
	var conf Config
	name := utils.GetHome() + host
	content, err := os.ReadFile(name)
	if err != nil {
		return conf, err
	}
	if err := yaml.Unmarshal(content, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func WriteHost(host, port, user, pwd string) error {
	conf := Config{
		Host: host,
		Port: port,
		User: user,
		Pwd:  pwd,
	}

	bytes, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	name := utils.GetHome() + conf.Host
	return os.WriteFile(name, bytes, 0644)
}

func DelHost(host string) {
	name := utils.GetHome() + host
	os.Remove(name)

}
