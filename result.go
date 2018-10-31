package ydfy

//Result is the struct of tranlate result
type Result struct {
	TranslateResult [][]struct {
		Src string `json:"src"`
		Tgt string `json:"tgt"`
	} `json:"translateResult"`
	SmartResult struct {
		Entries []string `json:"entries"`
	} `json:"smartResult"`
}
