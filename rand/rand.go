package rand

import (
	"crypto/rand"
	"io"
	"math/big"
	"time"

	mrnd "math/rand"

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
	bytes  []byte
	length int
	max    int
	bits   int
	mask   int
}

func StringScope(str string) *charScope {
	len := len(str)

	scope := &charScope{
		bytes:  bytesconv.StringToBytes(str),
		length: len,
		bits:   1,
	}

	for scope.mask < len {
		scope.mask = 1<<scope.bits - 1
		scope.bits++
	}

	return scope
}

var (
	Uppercase = StringScope(uppercase) // 大写字母
	Lowercase = StringScope(lowercase) // 小写字母
	Digit     = StringScope(digit)     // 数字
	Nomix     = StringScope(nomix)     // 不混淆字符
	Letter    = StringScope(letter)    // 字母
	Hex       = StringScope(hex)       // 十六进制字符
	AllChars  = StringScope(allChars)  // 所有字符
)

var randSource mrnd.Source = mrnd.NewSource(time.Now().UnixNano())

// 生成指定长度的随机字符串
func (scope *charScope) RandString(n int) string {

	ret := make([]byte, n)
	for i, cache, remain := n-1, randSource.Int63(), scope.max; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), scope.max
		}

		if idx := int(cache & int64(scope.mask)); idx < scope.length {
			ret[i] = scope.bytes[idx]
			i--
		}

		cache >>= int64(scope.bits)
		remain--
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
	sub := max
	if min > 0 {
		sub = max - min
	}

	n, _ := randNumber(int32(sub))
	return n + sub
}

func RandBytes(len int) ([]byte, error) {
	r := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, r)
	return r, err
}

func randNumber(max int32) (int32, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	return int32(r.Int64()), err
}
