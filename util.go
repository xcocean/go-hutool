package main

var singleton map[int]interface{}

const (
	Constant_Snowflake = 1
	Constant_ObjectId  = 2
)

func init() {
	singleton = make(map[int]interface{})
}

func GetSingleton(key int) interface{} {
	val := singleton[key]
	if val == nil {
		if key == Constant_Snowflake {
			snowflake, _ := NewSnowflake()
			singleton[key] = snowflake
			return snowflake
		} else if key == Constant_ObjectId {
			obj := NewObjectId()
			singleton[key] = obj
			return obj
		}
	}
	return val
}
