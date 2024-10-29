package yamlutils

import (
	"io/fs"
	"log"

	"gopkg.in/yaml.v2"
)

func GetYamlFS(dir fs.FS, fileName string, configType any) (err error) {

	yamlFile, err := fs.ReadFile(dir, fileName)
	if err != nil {
		log.Printf("yamlFile err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, configType)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return
}
