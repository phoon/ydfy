package ydfy

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	referer     = "http://fanyi.youdao.com/"
	cookie      = "OUTFOX_SEARCH_USER_ID=1678668575@10.168.1.8"
	contentType = "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent   = "Mozilla/5.0 (X11; Linux x86_64; rv:64.0) Gecko/20100101 Firefox/64.0"
)

func doPost(postURL string, params string) string {
	c := &http.Client{}
	body := strings.NewReader(params)

	req, err := http.NewRequest("POST", postURL, body)
	/*set headers*/
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Referer", referer)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentType)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(respBody)
}
