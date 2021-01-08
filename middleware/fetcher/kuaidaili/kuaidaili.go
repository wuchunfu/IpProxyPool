package kuaidaili

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	logger "github.com/sirupsen/logrus"
	"proxypool-go/middleware/fetcher"
	"proxypool-go/models/ipModel"
	"proxypool-go/util"
	"strconv"
	"strings"
)

// 国内高匿代理
func KuaiDaiLiInha() []*ipModel.IP {
	return KuaiDaiLi("inha")
}

// 国内普通代理
func KuaiDaiLiIntr() []*ipModel.IP {
	return KuaiDaiLi("intr")
}

func KuaiDaiLi(proxyType string) []*ipModel.IP {
	logger.Info("[kuaidaili] fetch start")

	list := make([]*ipModel.IP, 0)

	indexUrl := "https://www.kuaidaili.com/free"
	fetchIndex := fetcher.Fetch(indexUrl)
	pageNum := fetchIndex.Find("#listnav > ul > li:nth-child(9) > a").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/%s/%d", indexUrl, proxyType, i)
		fetch := fetcher.Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				proxyIp := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				proxyPort := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				proxyType := strings.TrimSpace(selection.Find("td:nth-child(4)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())
				proxySpeed := strings.TrimSpace(selection.Find("td:nth-child(6)").Text())

				ip := new(ipModel.IP)
				ip.ProxyHost = proxyIp
				ip.ProxyPort, _ = strconv.Atoi(proxyPort)
				ip.ProxyType = proxyType
				ip.ProxyLocation = proxyLocation
				ip.ProxySpeed, _ = strconv.Atoi(proxySpeed)
				ip.ProxySource = "https://www.kuaidaili.com"
				ip.CreateTime = util.FormatDateTime()
				ip.UpdateTime = util.FormatDateTime()
				list = append(list, ip)
			})
		})
	}
	logger.Info("[kuaidaili] fetch done")
	return list
}
