package ip3366

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	logger "github.com/sirupsen/logrus"
	"github.com/wuchunfu/IpProxyPool/fetcher"
	"github.com/wuchunfu/IpProxyPool/models/ipModel"
	"github.com/wuchunfu/IpProxyPool/util"
	"strconv"
	"strings"
)

// 国内高匿代理
func Ip33661() []*ipModel.IP {
	return Ip3366(1)
}

// 国内普通代理
func Ip33662() []*ipModel.IP {
	return Ip3366(2)
}

func Ip3366(proxyType int) []*ipModel.IP {
	logger.Info("[ip3366] fetch start")
	defer func() {
		recover()
		logger.Warnln("[ip3366] fetch error")
	}()
	list := make([]*ipModel.IP, 0)

	indexUrl := "http://www.ip3366.net/free"
	fetchIndex := fetcher.Fetch(indexUrl)
	pageNum := fetchIndex.Find("#listnav > ul > a:nth-child(8)").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/?stype=%d&page=%d", indexUrl, proxyType, i)
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
				ip.ProxySource = "http://www.ip3366.net"
				ip.CreateTime = util.FormatDateTime()
				ip.UpdateTime = util.FormatDateTime()
				list = append(list, ip)
			})
		})
	}
	logger.Info("[ip3366] fetch done")
	return list
}
