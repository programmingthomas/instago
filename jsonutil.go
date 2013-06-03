package instago

//The JSON type can be used when you do not directly want to parse JSON data into a Go
//struct, or when you are dealing with object types that are unknown or constantly 
//changing. The API uses this because a) The structure of some Instagram API requests adds
//a lot of additional unnecessary data that ought not be in the final Go response.
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

//Check to see if an interface is a float(64) and if it is not it will return 0
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

//Utility wrapper around JSONString
func (json JSON) String(key string) string {
	return JSONString(json[key])
}

//Utility wrapper around JSONInt
func (json JSON) Int(key string) int {
	return JSONInt(json[key])
}

//Utility wrapper around JSONFloat
func (json JSON) Float(key string) float64 {
	return JSONFloat(json[key])
}

//Utility wrapper around JSONArray
func (json JSON) Array(key string) []interface{} {
	return JSONArray(json[key])
}

//Utility wrapper around JSONObject
func (json JSON) Object(key string) JSON {
	return JSONObject(json[key])
}

//Utility wrapper around JSONStringArray
func (json JSON) StringArray(key string) []string {
	return JSONStringArray(json[key])
}

//Utility wrapper around JSONObjectArray
func (json JSON) ObjectArray(key string) []JSON {
	return JSONObjectArray(json[key])
}