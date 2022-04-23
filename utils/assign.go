package utils

import "reflect"

// StructAssign copy struct
//current 要修改的struct, former 有数据的struct
func StructAssign(current interface{}, former interface{}) {
	c := reflect.ValueOf(current).Elem()
	f := reflect.ValueOf(former).Elem()

	fType := f.Type()
	for i := 0; i < f.NumField(); i++ {
		name := fType.Field(i).Name
		if ok := c.FieldByName(name).IsValid(); ok {
			c.FieldByName(name).Set(reflect.ValueOf(f.Field(i).Interface()))
		}
	}
}
