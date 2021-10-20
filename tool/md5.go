package tool

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: https://github.com/18211167516
//@function: MD5
//@description: md5加密
//@param: str []byte
//@return: string

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
