package ydfy

const (
	/*English to Chinese*/
	e2c = 0
	/*Chinese to English*/
	c2e = 1
)

func params(i string, w int) string {
	salt := timestamp()
	if w == 0 {
		return "i=" + i +
			"&from=en&to=zh-CHS&smartresult=dict" +
			"&client=fanyideskweb&salt=" + salt +
			"&sign=" + sign(i, salt) +
			"&doctype=json&version=2.1&keyfrom=fanyi.web" +
			"&action=FY_BY_REALTIME&typoResult=false"
	}
	return "i=" + i +
		"&from=zh-CHS&to=en&smartresult=dict" +
		"&client=fanyideskweb&salt=" + salt +
		"&sign=" + sign(i, salt) +
		"&doctype=json&version=2.1&keyfrom=fanyi.web" +
		"&action=FY_BY_REALTIME&typoResult=false"
}
