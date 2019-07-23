package helper

import (
	"fmt"
	"reflect"
)

type data []map[string]interface{}

type dataMap map[interface{}]map[string]interface{}
//func GroupBy(d data,s string) (r map[string]interface{})  {
//	rs := GetValueByKey(d, s)
//	r2 := SliceRemoveDuplicate(rs)
//	for _,row := range d{
//		for k,r := range row{
//			fmt.Println("---",k,r)
//		}
//	}
//	return
//}
//获取某key值下的所有值
func GetValueByKey(d data,s string) (r []interface{})  {
	for _,row := range d{
		if row[s] != nil{
			r = append(r,row[s])
		}
	}
	return
}

//slice去重
func SliceRemoveDuplicate(d interface{}) (r []interface{}) {
	if reflect.TypeOf(d).Kind() != reflect.Slice {
		fmt.Printf("<SliceRemoveDuplicate> <a> is not slice but %T\n", d)
		return r
	}
	va := reflect.ValueOf(d)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i - 1).Interface(), va.Index(i).Interface()) {
			continue
		}
		r = append(r, va.Index(i).Interface())
	}
	return
}
//把map转为slice
func MapToSlice(input map[string]map[string]interface{}) (r []interface{}) {
	for _,row := range input {
 		r = append(r,row)
	}
	return
}
