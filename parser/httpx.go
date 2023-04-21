package parser

import (
	"encoding/json"
	"strings"
)

type HttpxRecord struct {
	Timestamp     string
	Hash          HttpxHash
	Port          string
	Url           string
	Input         string
	Title         string
	Scheme        string
	Webserver     string
	ContentType   string `json:"content_type"`
	Method        string
	Host          string
	Path          string
	Time          string
	A             []string
	Cname         []string
	Words         uint
	Lines         uint
	StatusCode    uint `json:"status_code"`
	ContentLength uint `json:"content_length"`
}

type HttpxHash struct {
	Body_md5     string
	Body_mmh3    string
	Body_sha256  string
	Body_simhash string

	Header_md5     string
	Header_mmh3    string
	Header_sha256  string
	Header_simhash string
}

func ParseHttpx(input string) []HttpxRecord {
	lines := strings.Split(input, "\n")
	var records []HttpxRecord

	for _, ln := range lines {
		data := &HttpxRecord{}
		byt := []byte(ln)
		if err := json.Unmarshal(byt, &data); err != nil {
			continue
		}
		records = append(records, *data)

		// _, err := json.MarshalIndent(data, "", "\t")
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
	}

	return records
}
