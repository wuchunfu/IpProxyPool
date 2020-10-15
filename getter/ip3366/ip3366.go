package ip3366

import (
	"github.com/Aiicy/htmlquery"
	logger "github.com/sirupsen/logrus"
	"proxypool-go/models/ipModel"
	"proxypool-go/util"
)

//IP3306 get ip from http://www.ip3366.net/
func IP3306() (result []*ipModel.IP) {
	logger.Info("[IP3366]] start Get IpProxy")
	pollURL := "http://www.ip3366.net/free/?stype=1&page=1"
	doc, _ := htmlquery.LoadURL(pollURL)
	trNode, err := htmlquery.Find(doc, "//div[@id='list']//table//tbody//tr")
	logger.Info("[IP3366] start up")
	if err != nil {
		logger.Info("[IP3366]] parse pollUrl error")
		logger.Error(err.Error())
	}
	//debug begin
	logger.Infof("[IP3366] len(trNode) = %d ", len(trNode))
	for i := 1; i < len(trNode); i++ {
		tdNode, _ := htmlquery.Find(trNode[i], "//td")
		ip := htmlquery.InnerText(tdNode[0])
		port := htmlquery.InnerText(tdNode[1])
		Type := htmlquery.InnerText(tdNode[3])
		speed := htmlquery.InnerText(tdNode[5])

		IP := ipModel.NewIP()
		IP.Data = ip + ":" + port

		if Type == "HTTPS" {
			IP.Type1 = "https"
			IP.Type2 = ""

		} else if Type == "HTTP" {
			IP.Type1 = "http"
		}
		IP.Speed = util.ExtractSpeed(speed)

		logger.Infof("[IP3366] ip.Data = %s,ip.Type = %s,%s ip.Speed = %d", IP.Data, IP.Type1, IP.Type2, IP.Speed)

		result = append(result, IP)
	}

	logger.Info("IP3366 done.")
	return
}
