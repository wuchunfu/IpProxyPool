package kuaidaili

import (
	"github.com/Aiicy/htmlquery"
	logger "github.com/sirupsen/logrus"
	"proxypool-go/models/ipModel"
	"proxypool-go/util"
)

// KDL get ip from kuaidaili.com
func KDL() (result []*ipModel.IP) {
	pollURL := "http://www.kuaidaili.com/free/inha/"
	doc, _ := htmlquery.LoadURL(pollURL)
	trNode, err := htmlquery.Find(doc, "//table[@class='table.table-bordered.table-striped']//tbody//tr")
	if err != nil {
		logger.Error(err.Error())
	}
	for i := 0; i < len(trNode); i++ {
		tdNode, _ := htmlquery.Find(trNode[i], "//td")
		ip := htmlquery.InnerText(tdNode[0])
		port := htmlquery.InnerText(tdNode[1])
		Type := htmlquery.InnerText(tdNode[3])
		speed := htmlquery.InnerText(tdNode[5])

		IP := ipModel.NewIP()
		IP.Data = ip + ":" + port
		if Type == "HTTPS" {
			IP.Type1 = "https"
			IP.Type2 = "https"
		} else if Type == "HTTP" {
			IP.Type1 = "http"
		}
		IP.Speed = util.ExtractSpeed(speed)
		result = append(result, IP)
	}

	logger.Info("[kuaidaili] done")
	return
}
