package proxylistplus

import (
	logger "github.com/sirupsen/logrus"
	"proxypool-go/models/ipModel"

	"github.com/Aiicy/htmlquery"
)

//PLP get ip from proxylistplus.com
func PLP() (result []*ipModel.IP) {
	pollURL := "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
	doc, _ := htmlquery.LoadURL(pollURL)
	trNode, err := htmlquery.Find(doc, "//div[@class='hfeed site']//table[@class='bg']//tbody//tr")
	if err != nil {
		logger.Error(err.Error())
	}
	for i := 3; i < len(trNode); i++ {
		tdNode, _ := htmlquery.Find(trNode[i], "//td")
		ip := htmlquery.InnerText(tdNode[1])
		port := htmlquery.InnerText(tdNode[2])
		Type := htmlquery.InnerText(tdNode[6])

		IP := ipModel.NewIP()
		IP.Data = ip + ":" + port

		if Type == "yes" {
			IP.Type1 = "https"
			IP.Type2 = ""

		} else if Type == "no" {
			IP.Type1 = "http"
		}

		logger.Infof("[PLP] ip.Data = %s,ip.Type = %s,%s", IP.Data, IP.Type1, IP.Type2)

		result = append(result, IP)
	}

	logger.Info("PLP done.")
	return
}
