package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/trist725/myleaf/log"
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
	MgoSessionNum int
	DBName        string
	XlsxPath      string
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
