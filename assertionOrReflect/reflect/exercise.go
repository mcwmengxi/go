package reflect

import "reflect"

// 1.反射实现类型断言工具
func TypeAssert(v interface{}) bool {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Struct {
		return true
	}
	return false
}

// 2.结构体转 Map 的技巧
func StructToMap(v interface{}) map[string]interface{}{
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	mapData := make(map[string]interface{})
	if TypeAssert(v){
		for i:=0;i<t.NumField();i++{
			mapData[t.Field(i).Name] = val.Interface()
		}
	}
	return mapData
}
