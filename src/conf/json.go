package conf

import (
	"encoding/json"
	"github.com/trist725/myleaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel      string
	LogPath       string
	WSAddr        string
	CertFile      string
	KeyFile       string
	TCPAddr       string
	MaxConnNum    int
	ConsolePort   int
	ProfilePath   string
	MgoUrl        string
	WebUrl        string
	UnionPlatUrl  string
	MgoSessionNum int
	MgoName       string
	XlsxPath      string
	MerchantCode  string
	ProductCode   string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
