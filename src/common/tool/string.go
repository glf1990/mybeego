package tool

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)
func Map_to_json(map_str map[string]interface{}) (string,error){
	b, err := json.Marshal(map_str)
	if err != nil {
		//fmt.Println("json.Marshal failed:", err)
		return "",err
	}
	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(map_str)
	return  string(b),nil
}
func MapGroup_to_jsonGroup(map_str []map[string]interface{}) (string,error){
	b, err := json.Marshal(map_str)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return "",err
	}
	return  string(b),nil
}
func Json_to_map(json_string string) (map[string]interface{},error){
	//json str 转map
	var map_data map[string]interface{}
	if err := json.Unmarshal([]byte(json_string), map_data); err != nil {
		return nil ,err
	}
	return map_data,nil
}
func Json_to_struct(json_string string,config *interface{}) (*interface{},error){
	if err := json.Unmarshal([]byte(json_string), config); err != nil {
		return nil ,err

	}
	return  config,nil
}
func Struct_to_json(config *interface{}) (string,error){
	json_string,err := json.Marshal( config)
	if  err != nil {
		return "" ,err
	}
	return  string(json_string),nil
}

func String_to_int(value string) (int,error) {
	i,err:=strconv.Atoi(value)
	if  err != nil {
		return 0,err
	}
	return i,nil
}
//replace_num为替换次数 为-1时全部替换
func String_repalce(value string,old string,new string ) string{
	return strings.Replace(value, old, new, -1)
}
//replace_num为替换次数
func String_repalce_num(value string,old string,new string,replace_num  int ) string{
	return strings.Replace(value, old, new, replace_num)

}
