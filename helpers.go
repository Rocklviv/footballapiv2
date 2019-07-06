package footballapiv2

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func structToMap(i interface{}) (values url.Values) {
	if i != nil {
		values = url.Values{}
		iVal := reflect.ValueOf(i).Elem()
		typ := iVal.Type()
		for i := 0; i < iVal.NumField(); i++ {
			f := iVal.Field(i)
			// You ca use tags here...
			// tag := typ.Field(i).Tag.Get("tagname")
			// Convert each type into a string for the url.Values string map
			var v string
			switch f.Interface().(type) {
			case int, int8, int16, int32, int64:
				v = strconv.FormatInt(f.Int(), 10)
			case uint, uint8, uint16, uint32, uint64:
				v = strconv.FormatUint(f.Uint(), 10)
			case float32:
				v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
			case float64:
				v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
			case []byte:
				v = string(f.Bytes())
			case string:
				v = f.String()
			case []string:
				fmt.Println(reflect.TypeOf(iVal).Kind())
			}
			values.Set(makeFirstLowerCase(typ.Field(i).Name), v)
		}
		return
	}
	return
}

// func setNilIf(v *interface{}) {
// 	*v = nil
// }

func makeFirstLowerCase(str string) string {
	if len(str) < 2 {
		return strings.ToLower(str)
	}

	bts := []byte(str)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
