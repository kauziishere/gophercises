package yamlUtil

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	MaxDataBytes = 10000
)

func FetchMapFromYAMLFile(filename string) (map[interface{}]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if nil != err {
		return nil, err
	}

	dataMap := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &dataMap)
	if nil != err {
		return nil, err
	}

	return dataMap, err
}
