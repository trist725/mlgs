package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/trist725/myleaf/log"
)

var Client struct {
	TCPAddr string
}

func init() {
	data, err := ioutil.ReadFile("conf/client.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Client)
	if err != nil {
		log.Fatal("%v", err)
	}
}
