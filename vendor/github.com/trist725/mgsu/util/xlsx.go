package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/tealeg/xlsx"
)

func DeserializeStructFromXlsxRow(e interface{}, row *xlsx.Row) (err error) {
	v := reflect.ValueOf(e)
	ve := v.Elem()
	if ve.Kind() != reflect.Struct {
		return fmt.Errorf("[%#v] is not a struct", e)
	}

	t := reflect.TypeOf(e)
	te := t.Elem()
	for i := 0; i < ve.NumField(); i++ {
		column := i + 1
		tf := te.Field(i)
		tagColumn := tf.Tag.Get("excel_column")
		if tagColumn != "" {
			if column, err = strconv.Atoi(tagColumn); err != nil {
				return fmt.Errorf("get %s column fail, %s", tf.Name, err)
			}
		}

		if column >= len(row.Cells) {
			continue
		}

		cell := row.Cells[column]
		if cell.Value == "" {
			continue
		}

		vef := ve.Field(i)
		kind := vef.Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i64, err := cell.Int64()
			if err != nil {
				return fmt.Errorf("could not set %s %s to %s, %s", kind.String(), tf.Name, cell.Value, err)
			}
			vef.SetInt(i64)

		case reflect.String:
			vef.SetString(cell.String())

		case reflect.Float32, reflect.Float64:
			f64, err := cell.Float()
			if err != nil {
				return fmt.Errorf("could not set %s %s to %s, %s", kind.String(), tf.Name, cell.Value, err)
			}
			vef.SetFloat(f64)

		case reflect.Slice, reflect.Struct, reflect.Array:
			if err := json.Unmarshal([]byte(cell.String()), vef.Addr().Interface()); err != nil {
				return fmt.Errorf("could not set %s %s to %s, %s", kind.String(), tf.Name, cell.Value, err)
			}

		default:
			if err := json.Unmarshal([]byte(cell.String()), vef.Interface()); err != nil {
				return fmt.Errorf("could not set %s %s to %s, %s", kind.String(), tf.Name, cell.Value, err)
			}
		}
	}

	return nil
}
