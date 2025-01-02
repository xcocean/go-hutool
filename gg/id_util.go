package gg

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	epoch             int64 = 1609459200000                                  // 起始时间戳：2021-01-01 00:00:00
	workerIDBits      uint8 = 5                                              // workerId 占用的位数
	datacenterIDBits  uint8 = 5                                              // datacenterId 占用的位数
	sequenceBits      uint8 = 12                                             // 序列号占用的位数
	workerIDMax       int64 = -1 ^ (-1 << workerIDBits)                      // 最大 workerId
	datacenterIDMax   int64 = -1 ^ (-1 << datacenterIDBits)                  // 最大 datacenterId
	sequenceMask      int64 = -1 ^ (-1 << sequenceBits)                      // 序列号掩码
	timestampShift    uint8 = workerIDBits + datacenterIDBits + sequenceBits // 时间戳左移位数
	datacenterIDShift uint8 = workerIDBits + sequenceBits                    // datacenterId 左移位数
	workerIDShift     uint8 = sequenceBits                                   // workerId 左移位数
)

// 雪花算法 start------------------------------------------------------------------------------------------

// Snowflake 结构体
type Snowflake struct {
	mu           sync.Mutex
	lastStamp    int64 // 上次生成ID的时间戳
	workerID     int64 // workerId
	datacenterID int64 // datacenterId
	sequence     int64 // 序列号
}

// 创建一个 Snowflake 实例
func Id_NewSnowflake() *Snowflake {
	return Id_NewSnowflake2(1, 1)
}

// 创建一个 Snowflake 实例
//
//	snowflake2 := Id_NewSnowflake2(1, 1)
//	println(snowflake2.NextID())
//	println(snowflake2.NextID())
//	println(snowflake2.ParseTimestamp(snowflake2.NextID()).Unix())
func Id_NewSnowflake2(workerID, datacenterID int64) *Snowflake {
	if workerID < 0 || workerID > workerIDMax {
		panic(fmt.Sprintf("workerId 超出范围"))
	}
	if datacenterID < 0 || datacenterID > datacenterIDMax {
		panic(fmt.Sprintf("datacenterId 超出范围"))
	}
	return &Snowflake{
		lastStamp:    0,
		workerID:     workerID,
		datacenterID: datacenterID,
		sequence:     0,
	}
}

// NextIDTo36 生成下一个ID 转为36进制字符串 --> 40w28qcxz402  （长度等于12）
func (s *Snowflake) NextIDTo36() string {
	return strconv.FormatInt(s.NextID(), 36)
}

// NextIDStr 生成下一个ID --> 529742306929152000（长度等于18）
func (s *Snowflake) NextIDStr() string {
	return fmt.Sprintf("%010d", s.NextID())
}

// NextID 生成下一个ID --> 529742306929152000（长度等于18）
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 获取当前时间戳
	now := time.Now().UnixNano() / 1e6

	// 如果当前时间小于上次生成ID的时间，说明时钟回拨
	if now < s.lastStamp {
		panic("时钟回拨")
	}

	// 如果是同一毫秒内生成的ID
	if now == s.lastStamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		// 如果序列号溢出，等待下一毫秒
		if s.sequence == 0 {
			for now <= s.lastStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = now

	// 生成ID
	id := ((now - epoch) << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}

// ParseTimestamp 解析ID中的时间戳
func (s *Snowflake) ParseTimestamp(id int64) time.Time {
	// 右移并加上起始时间戳
	timestamp := (id >> timestampShift) + epoch
	return time.Unix(timestamp/1000, (timestamp%1000)*1e6)
}

// 获取 Snowflake 单例对象
func Id_getSnowflake() *Snowflake {
	getSingleton := GetSingleton(Constant_Snowflake)
	return getSingleton.(*Snowflake)
}

// 雪花算法 end------------------------------------------------------------------------------------------

// MongoDB 的 ObjectId 算法 start------------------------------------------------------------------------------------------

// ObjectId 是一个 12 字节的标识符
type ObjectId [12]byte

// 生成一个新的 ObjectId
func NewObjectId() ObjectId {
	var id ObjectId

	// 4 字节时间戳（秒）
	timestamp := uint32(time.Now().Unix())
	binary.BigEndian.PutUint32(id[0:4], timestamp)

	// 3 字节机器标识
	machine := getMachineHash()
	copy(id[4:7], machine[:3])

	// 2 字节进程 ID
	pid := uint16(os.Getpid())
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)

	// 3 字节随机数
	random := Random_uint32()
	id[9] = byte(random >> 16)
	id[10] = byte(random >> 8)
	id[11] = byte(random)

	return id
}

// getMachineHash 获取机器标识的哈希值
func getMachineHash() []byte {
	hostname, _ := os.Hostname()
	hash := md5.Sum([]byte(hostname))
	return hash[:]
}

// 获取一个 objectId 24长度字符串 --> 6776672d1370e04ba81f6922
func (id ObjectId) Next() string {
	return hex.EncodeToString(id[:])
}

// 获取 ObjectId 单例
// Id_ObjectId().Next() --> 24长度字符串
func Id_ObjectId() ObjectId {
	return GetSingleton(Constant_ObjectId).(ObjectId)
}

// MongoDB 的 ObjectId 算法 end------------------------------------------------------------------------------------------

// 生成一个uuid --> b2b26b89-4fd1-405d-945f-5f70f9063422
func Id_RandomUUID() string {
	uuid := Random_randomBytes(16)

	// 设置版本号（第 7 字节的高 4 位为 4）
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// 设置变体（第 9 字节的高 2 位为 10）
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:36], uuid[10:16])
	return string(buf)
}

func Id_FastUUID() string {
	uuid := Random_randomBytes(16)
	// 设置版本为4
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// 设置变体为RFC 4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	// 将UUID转换为字符串，不带横线
	return fmt.Sprintf("%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:])
}

// 随机生成一个 21 位的字符串
func Id_NanoId() string {
	return Id_NanoIdLength(21)
}

// 随机生成一个指定长度的字符串
func Id_NanoIdLength(length int) string {
	if length <= 0 {
		panic("length 必须大于 0")
	}
	return Random_randomString(length)
}
