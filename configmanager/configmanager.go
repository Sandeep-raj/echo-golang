package configmanager

import (
	"encoding/json"
	"io/ioutil"
)

type Configmanager struct {
	DatasetList  []string `json:"dataset_list"`
	MongoConnStr string   `json:"mongo_conn_str"`
	Database     string   `json:"database"`
}

var ConfStore *Configmanager

func InitConfig() error {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &ConfStore)
	if err != nil {
		return err
	}
	return nil
}
