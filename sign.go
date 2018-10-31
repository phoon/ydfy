package ydfy

import (
	"crypto/md5"
	"fmt"
	"time"
)

func timestamp() string {
	//ms level timestamp
	return fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
}

func sign(e, t string) string {
	data := []byte("fanyideskweb" + e + t + "6x(ZHw]mwzX#u0V7@yfwK")
	return fmt.Sprintf("%x", md5.Sum(data))
}
