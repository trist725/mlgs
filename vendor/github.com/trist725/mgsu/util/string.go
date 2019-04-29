package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func DeserializeFromStringArray(i interface{}, stringArray []string) (err error) {
	v := reflect.ValueOf(i).Elem()
	t := reflect.TypeOf(i)
	for i := 0; i < v.NumField(); i++ {
		column := i + 1

		tf := t.Field(i)
		tagColumnValue := tf.Tag.Get("column")
		if tagColumnValue != "" {
			if column, err = strconv.Atoi(tagColumnValue); err != nil {
				return fmt.Errorf("parse [%v][%v] fail, %v\n",
					i+1, t.Field(i).Name, err)
			}
		}

		f := v.Field(i)
		switch f.Type().Name() {
		case "int32", "int", "int64":
			temp, err := strconv.ParseInt(stringArray[column], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetInt(temp)

		case "string":
			f.SetString(stringArray[i+1])

		case "bool":
			temp, err := strconv.ParseInt(stringArray[i+1], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetBool(temp != 0)

		case "Duration":
			temp, err := strconv.ParseInt(stringArray[i+1], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetInt(temp * int64(time.Millisecond))

		default:
			return fmt.Errorf("parse [%v][%v] fail, unsupported field type",
				column, t.Field(i).Name)
		}
	}
	return nil
}

// 格式化字段名, 比如: xlsx表字段, mysql表字段
func FormatFieldName(raw string) (name string) {
	words := strings.Split(raw, "_")
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		word = strings.Title(word)
		name += word
	}

	name = strings.Replace(name, "Id", "ID", 1)

	return
}
