package jsonutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetJson(path string, jsonInterface interface{}) {
	content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err);
    }
	jsonErr := json.Unmarshal(content, &jsonInterface);
	if jsonErr != nil {
        log.Fatal(err);
    }
}