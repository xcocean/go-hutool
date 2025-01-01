package main

import "testing"

func TestSnowflake(t *testing.T) {
	snowflake2, _ := NewSnowflake2(1, 1)
	println(snowflake2.NextID())
	println(snowflake2.ParseTimestamp(snowflake2.NextID()).Unix())

	println(GetSnowflake().NextID())
	println(GetSnowflake().NextIDStr())
	println(GetSnowflake().NextIDTo36())

	println(GetObjectId().Next())
}

func TestUUID(t *testing.T) {
	println(RandomUUID())
	println(FastUUID())
	println(NanoId())
}
