package ydfy

const postURL = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"

//En2zhCHS translate English to Chinese
func En2zhCHS(i string) string {
	return doPost(postURL, params(i, e2c))
}

//ZhCHS2en translate Chinese to English
func ZhCHS2en(i string) string {
	return doPost(postURL, params(i, c2e))
}
