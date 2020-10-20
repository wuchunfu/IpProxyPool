package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"proxypool-go/util/headerutil"
	"time"
)

func Fetch(url string) *goquery.Document {
	logrus.Infof("Fetch url: %s", url)
	// &cookiejar.Options{PublicSuffixList: publicsuffix.List}，这是为了可以根据域名安全地设置cookies
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar:     cookieJar,
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Proxy-Switch-Ip", "yes")
	req.Header.Set("Referer", "https://www.baidu.com")
	//req.Header.Set("Referrer Policy", "strict-origin-when-cross-origin")
	req.Header.Set("User-Agent", headerutil.RandomUserAgent())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Content-Length", "25")
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	//req.Header.Set("X-Anit-Forge-Code", "0")
	//req.Header.Set("X-Anit-Forge-Token", "None")
	//req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("起死回生")
		}
	}()
	if err != nil {
		//logrus.Errorf("http get error: %v", err)
		//return nil
		panic(err)
	}
	//if resp.StatusCode != http.StatusOK {
	//	logrus.Errorf("error http status code: %d", resp.StatusCode)
	//}

	var newResp io.Reader
	var charsetErr error

	var doc *goquery.Document
	var docErr error

	if resp.StatusCode == http.StatusOK {
		newResp, charsetErr = charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if charsetErr != nil {
			logrus.Errorf("charset convert failed: %v", charsetErr)
		}
		doc, docErr = goquery.NewDocumentFromReader(newResp)
		if docErr != nil {
			logrus.Errorf("goquery http response body reader error: %v", docErr)
		}
	}
	return doc
}
