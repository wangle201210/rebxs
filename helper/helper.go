package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
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



const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "KeeOQqdFYTa3gTcUy2yrEABU"
	APISECRET = "CGrSFRHdLEdTZXF5Qy6zi9wXcZvdTZaO"
	TEMPLATE = "882f23fbc8c3bda063747a0784ecabe6"
)

type GetType map[string]map[string][]map[string]string

type DataList struct {
	User string
	Score int64
}

var client *ocr.OCRClient

func init() {
client = ocr.NewOCRClient(APIKEY, APISECRET)
}



func GeneralRecognizeBasic(f string) (r map[string]DataList,e error) {
	fmt.Println(f)
	rs, err := client.GeneralRecognizeWithLocation(
		vision.MustFromFile(f),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
	var getData GetType
	_ = rs.ToJSON(&getData)
	if len(getData["data"]["ret"]) == 0 {
		e = errors.New("没得数据哈！！")
		return
	}
	lists := formatData(getData)
	r = lists
	//fmt.Println(rs.ToString())
	return
}
func formatData(data GetType) map[string]DataList {
	res := make(map[string]DataList)
	for _,v := range data["data"]["ret"] {
		strArr := strings.Split(v["word_name"],"#")
		var r DataList
		r = res[strArr[1]]
		if strArr[2] == "名字" {
			if len(v["word"]) > 9 && v["word"][len(v["word"])-9:] == "副族长"{
				r.User = v["word"][0:len(v["word"])-9]
			} else if len(v["word"]) > 6{
				r.User = v["word"][0:len(v["word"])-6]
			} else {
				continue
			}
		} else {
			i, _ := strconv.ParseInt(v["word"], 10, 64)
			r.Score = i
		}
		res[strArr[1]] = r
	}
	return res
}