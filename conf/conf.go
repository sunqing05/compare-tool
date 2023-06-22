package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	DataDir        string `yaml:"data_dir"`
	BaseFolderName string `yaml:"base_folder_name"`
	DiffFolderName string `yaml:"diff_folder_name"`
}

func Config() Conf {
	data, err := os.ReadFile("conf.yaml")
	if err != nil {
		fmt.Printf("读取配置文件错误：%v \n", err)
	}
	var config Conf
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("读取配置文件错误：%v \n", err)
	}
	return config
}
