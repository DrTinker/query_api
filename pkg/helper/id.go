package helper

import (
	"hash/crc32"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GenUid(name string, phone int) int {
	u := uuid.NewV4()
	str := u.String() + name + strconv.Itoa(phone)
	id := crc32.ChecksumIEEE([]byte(str))
	res := int(id)
	if res < 0 {
		res = -res
	}
	return res
}

func GenQid() int {
	u := uuid.NewV4()
	str := u.String() + strconv.Itoa(int(time.Now().Unix()))
	id := crc32.ChecksumIEEE([]byte(str))
	res := int(id)
	if res < 0 {
		res = -res
	}
	return res
}
