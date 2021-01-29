package useragentutil

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUA(t *testing.T) {
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
	agent := GetUserAgent(ua)
	marshal, err := json.Marshal(agent)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(marshal))
}
