package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"net/http"
	"strconv"
	"strings"
)

func Fetch(url string) *goquery.Document {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)")
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		logrus.Error("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	res, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		logrus.Error("Http status codess:", err)
	}
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		logrus.Fatal(err)
	}
	return doc
}

func ip66() {
	indexUrl := "http://www.66ip.cn"
	fetchIndex := Fetch(indexUrl)
	pageNum := fetchIndex.Find("#PageList > a:nth-child(12)").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/%d.html", indexUrl, i)
		fetch := Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").NextAll().Each(func(i int, selection *goquery.Selection) {
				ip := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				port := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				proxyType := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				verifyTime := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				fmt.Println(ip)
				fmt.Println(port)
				fmt.Println(proxyLocation)
				fmt.Println(proxyType)
				fmt.Println(verifyTime)
				fmt.Println("======")
			})
		})
	}
}

func ip89() {
	indexUrl := "https://www.89ip.cn/"
	fetchIndex := Fetch(indexUrl)
	pageNum := fetchIndex.Find("#layui-laypage-1 > a:nth-child(7)").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/index_%d.html", indexUrl, i)
		fetch := Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				ip := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				port := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				serviceProvider := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				recordingTime := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				fmt.Println(ip)
				fmt.Println(port)
				fmt.Println(proxyLocation)
				fmt.Println(serviceProvider)
				fmt.Println(recordingTime)
				fmt.Println("======")
			})
		})
	}
}

func ip3366(proxyType int) {
	indexUrl := "http://www.ip3366.net/free"
	fetchIndex := Fetch(indexUrl)
	pageNum := fetchIndex.Find("#listnav > ul > a:nth-child(8)").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/?stype=%d&page=%d", indexUrl, proxyType, i)
		fetch := Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				ip := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				port := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				anonymity := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				protocolType := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				responseSpeed := strings.TrimSpace(selection.Find("td:nth-child(6)").Text())
				lastVerifyTime := strings.TrimSpace(selection.Find("td:nth-child(7)").Text())
				fmt.Println(ip)
				fmt.Println(port)
				fmt.Println(anonymity)
				fmt.Println(protocolType)
				fmt.Println(proxyLocation)
				fmt.Println(responseSpeed)
				fmt.Println(lastVerifyTime)
				fmt.Println("======")
			})
		})
	}
}

func kuaidaili(proxyType string) {
	indexUrl := "https://www.kuaidaili.com/free"
	fetchIndex := Fetch(indexUrl)
	pageNum := fetchIndex.Find("#listnav > ul > li:nth-child(9) > a").Text()
	fmt.Println(pageNum)

	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/%s/%d", indexUrl, proxyType, i)
		fetch := Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				ip := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				port := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				anonymity := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				protocolType := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				responseSpeed := strings.TrimSpace(selection.Find("td:nth-child(6)").Text())
				lastVerifyTime := strings.TrimSpace(selection.Find("td:nth-child(7)").Text())
				fmt.Println(ip)
				fmt.Println(port)
				fmt.Println(anonymity)
				fmt.Println(protocolType)
				fmt.Println(proxyLocation)
				fmt.Println(responseSpeed)
				fmt.Println(lastVerifyTime)
				fmt.Println("======")
				fmt.Println("#list > table > tbody > tr:nth-child(1) > td:nth-child(1)")
				fmt.Println("#list > table > tbody > tr:nth-child(1) > td:nth-child(2)")
				fmt.Println("#list > table > tbody > tr:nth-child(2) > td:nth-child(1)")
			})
		})
	}
}

func proxylistplus(proxyType string) {
	indexUrl := "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
	fetchIndex := Fetch(indexUrl)
	pageNum := fetchIndex.Find("#listnav > ul > li:nth-child(9) > a").Text()
	fmt.Println(pageNum)

	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/%s/%d", indexUrl, proxyType, i)
		fetch := Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				ip := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				port := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				anonymity := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				protocolType := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				responseSpeed := strings.TrimSpace(selection.Find("td:nth-child(6)").Text())
				lastVerifyTime := strings.TrimSpace(selection.Find("td:nth-child(7)").Text())
				fmt.Println(ip)
				fmt.Println(port)
				fmt.Println(anonymity)
				fmt.Println(protocolType)
				fmt.Println(proxyLocation)
				fmt.Println(responseSpeed)
				fmt.Println(lastVerifyTime)
				fmt.Println("======")
				fmt.Println("#list > table > tbody > tr:nth-child(1) > td:nth-child(1)")
				fmt.Println("#list > table > tbody > tr:nth-child(1) > td:nth-child(2)")
				fmt.Println("#list > table > tbody > tr:nth-child(2) > td:nth-child(1)")
			})
		})
	}
}

func main() {
	//ip66()
	//ip89()

	// // 国内高匿代理
	//ip3366(1)
	// // 国内普通代理
	//ip3366(2)

	// // 国内高匿代理
	//kuaidaili("inha")
	// // 国内普通代理
	//kuaidaili("intr")

}
