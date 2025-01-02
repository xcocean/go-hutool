package gg

import "testing"

func TestSnowflake(t *testing.T) {
	snowflake2 := Id_NewSnowflake2(1, 1)
	println(snowflake2.NextID())
	println(snowflake2.ParseTimestamp(snowflake2.NextID()).Unix())

	println(Id_getSnowflake().NextID())
	println(Id_getSnowflake().NextIDStr())
	println(Id_getSnowflake().NextIDTo36())

	println(Id_ObjectId().Next())
}

func TestUUID(t *testing.T) {
	println(Id_RandomUUID())
	println(Id_FastUUID())
	println(Id_NanoId())
	println(Id_NanoIdLength(5))
}

func TestId01(t *testing.T) {
	println(NewObjectId().Next())
}
