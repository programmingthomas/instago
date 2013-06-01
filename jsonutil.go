//This file contains some utility functions that make dealing with the JSON a little easier
//Go's json package (encoding/json) will convert everything to a map[string]interface{}
//however that is not entirely relevant when working with the Instagram API and it can be
//a pain to access even the most simple of functions because Go assumes that values are
//not strings/numbers/arrays/etc by default, however when working with the API we know
//that they are, so we can use this instead
package instago

//JSON Objects are basicallly just maps because the JSON spec insists on keys being strings
type JSON map[string]interface{}

//Check to see if an interface is a string and if it is not, return an empty string
func JSONString(data interface{}) string {
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}

//Check to see if an interface is an int and if it is not it will return 0
func JSONInt(data interface{}) int {
	//N.B. The encoding/json library assumes all numbers as float64 but most Instagram
	//values are actually integers
	return int(JSONFloat(data))
}

//Check to see if an interface is a float and if it is not it will return 0
func JSONFloat(data interface{}) float64 { 
	if number, ok := data.(float64); ok {
		return number
	}
	return 0
}

//Checks to see if an interface is an array, and if not return an empty array
func JSONArray(data interface{}) []interface{} {
	if arr, ok := data.([]interface{}); ok {
		return arr
	} 
	return make([]interface{}, 0)
}

//Checks to see if an object is a JSON object and if not, return an empty object
func JSONObject(data interface{}) JSON {
	if obj, ok := data.(map[string]interface{}); ok {
		return obj
	}
	return make(map[string]interface{}, 0)
}

//Checks to see if an interface is an array of strings, and if not return an empty string
//array
func JSONStringArray(data interface{}) []string {
	arr := JSONArray(data)
	strings := make([]string, 0)
	for _, v := range arr {
		strings = append(strings, JSONString(v))
	}
	return strings
}

//Check to see if it is an array of objects and if not return an empty JSON object array
func JSONObjectArray(data interface{}) []JSON {
	arr := JSONArray(data)
	objs := make([]JSON, 0)
	for _, v := range arr {
		objs = append(objs, JSONObject(v))
	}
	return objs
}

//The following functions provide utility functions for working directly with a JSON //object (a map)

func (json JSON) String(key string) string {
	return JSONString(json[key])
}

func (json JSON) Int(key string) int {
	return JSONInt(json[key])
}

func (json JSON) Float(key string) float64 {
	return JSONFloat(json[key])
}

func (json JSON) Array(key string) []interface{} {
	return JSONArray(json[key])
}

func (json JSON) Object(key string) JSON {
	return JSONObject(json[key])
}

func (json JSON) StringArray(key string) []string {
	return JSONStringArray(json[key])
}

func (json JSON) ObjectArray(key string) []JSON {
	return JSONObjectArray(json[key])
}