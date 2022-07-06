package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	Header         = "www.baidu.com"
	HeaderLength   = 13
	SaveDataLength = 4
)

// 封包
func Packet(message []byte) []byte {
	// 头部信息 + body长度 + 消息
	return append(append([]byte(Header), IntToBytes(len(message))...), message...)
}

// 解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		if length < i+HeaderLength+SaveDataLength { // 长度为最小头部信息长（不包含）
			break
		}
		if string(buffer[i:i+HeaderLength]) == Header { // 是否是我们约定的头部信息
			messageLength := BytesToInt(buffer[i+HeaderLength : i+HeaderLength+SaveDataLength]) // 信息body长度
			if length < i+HeaderLength+SaveDataLength+messageLength {
				break
			}
			data := buffer[i+HeaderLength+SaveDataLength : i+HeaderLength+SaveDataLength+messageLength] // 信息
			readerChannel <- data

			i += HeaderLength + SaveDataLength + messageLength - 1 // end index
		}
	}

	if i == length { // 没有找到我们约定的头部信息, return empty []byte
		return make([]byte, 0)
	}
	return buffer[i:] // return message
}

// 整形转换成字节 int32 4个字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
