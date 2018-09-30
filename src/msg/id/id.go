package main

import (
	"encoding/json"
	"mlgs/src/msg"
	"os"
	"reflect"
)

func Gen() {
	idMap := map[string]uint16{}
	var f *os.File
	var err error
	msg.Processor.Range(func(id uint16, t reflect.Type) {
		f, err = os.OpenFile("./conf/id.json", os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		idMap[t.String()] = id
	})
	defer f.Close()

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
