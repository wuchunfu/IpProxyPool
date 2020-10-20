package proxylistplus

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

func ProxyListPlus() []*ipModel.IP {
	logger.Info("[proxylistplus] fetch start")

	list := make([]*ipModel.IP, 0)

	indexUrl := "https://list.proxylistplus.com"
	for i := 1; i <= 6; i++ {
		url := fmt.Sprintf("%s/Fresh-HTTP-Proxy-List-%d", indexUrl, i)
		fetch := fetcher.Fetch(url)
		fetch.Find("table.bg > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				proxyIp := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				proxyPort := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(5)").Text())

				ip := new(ipModel.IP)
				ip.ProxyHost = proxyIp
				ip.ProxyPort, _ = strconv.Atoi(proxyPort)
				ip.ProxyType = "http"
				ip.ProxyLocation = proxyLocation
				ip.ProxySpeed = 100
				ip.CreateTime = util.FormatDateTime()
				ip.UpdateTime = util.FormatDateTime()
				list = append(list, ip)
			})
		})
	}
	logger.Info("[proxylistplus] fetch done")
	return list
}
