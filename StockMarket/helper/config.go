package helper

import	(
	"io"
	"fmt"
	"gopkg.in/yaml.v2"
)

const filePath := "../config/config.yaml"

type Config struct{
}

func (config *Config) load(){
	data,err := io.ReadAll(filePath)

	if err != nil {

		fmt.Println("Error occured during config load)
		
		return

	}

	err := yaml.Unmarshal(data,config)

	if err != nil{

