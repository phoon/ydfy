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
	data := []byte("fanyideskweb" + e + t + "Tbh5E8=q6U3EXe+&L[4c@")
	return fmt.Sprintf("%x", md5.Sum(data))
}
