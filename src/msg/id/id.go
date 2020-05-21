package main

import (
	"encoding/json"
	"github.com/trist725/mlgs/src/msg"
	"os"
	"reflect"
)

func Gen() {
	idMap := map[string]uint16{}
	var f *os.File
	var err error
	f, err = os.OpenFile("./conf/id.json", os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	msg.Processor.Range(func(id uint16, t reflect.Type) {
		//5是过滤*msg.
		idMap[t.String()[5:]] = id
	})

	data, err := json.Marshal(idMap)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	Gen()
}
