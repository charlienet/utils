package rand

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/charlienet/utils/bytesconv"
)

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	digit     = "0123456789"
	nomix     = "BCDFGHJKMPQRTVWXY2346789"
	letter    = uppercase + lowercase
	allChars  = uppercase + lowercase + digit
	hex       = digit + "ABCDEF"
)

type charScope struct {
	bytes []byte
}

var (
	Uppercase = &charScope{bytes: bytesconv.StringToBytes(uppercase)} // 大写字母
	Lowercase = &charScope{bytes: bytesconv.StringToBytes(lowercase)} // 小写字母
	Digit     = &charScope{bytes: bytesconv.StringToBytes(digit)}     // 数字
	Nomix     = &charScope{bytes: bytesconv.StringToBytes(nomix)}     // 不混淆字符
	Letter    = &charScope{bytes: bytesconv.StringToBytes(letter)}    // 字母
	Hex       = &charScope{bytes: bytesconv.StringToBytes(hex)}       // 十六进制字符
	AllChars  = &charScope{bytes: bytesconv.StringToBytes(allChars)}  // 所有字符
)

// 生成指定长度的随机字符串
func (scope *charScope) RandString(length int) string {
	charLength := len(scope.bytes) - 1

	ret := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		ret = append(ret, scope.bytes[RandInt(0, charLength)])
	}

	return bytesconv.BytesToString(ret)
}

// 获取指定范围内的随机数
func RandInt(min, max int) int {
	for {
		n, _ := randNumber(int32(max))
		intn := int(n)
		if intn >= min {
			return intn
		}
	}
}

// 生成指定范围的随机数
func RandInt32(min, max int32) int32 {
	for {
		n, _ := randNumber(int32(max))
		if n >= min {
			return n
		}
	}
}

func RandByts(len int) ([]byte, error) {
	r := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, r)
	return r, err
}

func randNumber(max int32) (int32, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	return int32(r.Int64()), err
}
