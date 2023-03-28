package utils

import (
	"reflect"
	"strconv"
	"time"
)

// map转结构体  m(map)  s(结构体)
func MapToStruct(m map[string]string, s interface{}) error {
	rv := reflect.ValueOf(s).Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.CanSet() {
			continue
		}

		field := rt.Field(i)
		tagName := field.Tag.Get("redis")
		if tagName == "" {
			continue
		}

		val, ok := m[tagName]
		if !ok {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			f.SetString(val)
		case reflect.Uint64:
			if u64, err := strconv.ParseUint(val, 10, 64); err == nil {
				f.SetUint(u64)
			}
		case reflect.Uint8:
			if u8, err := strconv.ParseUint(val, 10, 8); err == nil {
				f.SetUint(u8)
			}
		case reflect.Struct:
			if field.Type == reflect.TypeOf(time.Time{}) {
				if ts, err := strconv.ParseInt(val, 10, 64); err == nil {
					t := time.Unix(ts, 0)
					f.Set(reflect.ValueOf(t))
				}
			}
		}
	}

	return nil
}
