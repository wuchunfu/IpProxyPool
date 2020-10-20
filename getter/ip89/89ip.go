package ip89
//
//import (
//	logger "github.com/sirupsen/logrus"
//	"io/ioutil"
//	"net/http"
//	"proxypool-go/models/ipModel"
//	"strconv"
//
//	"regexp"
//	"strings"
//)
//
////IP89 get ip from www.89ip.cn
//func IP89() (result []*ipModel.IP) {
//	logger.Info("89IP] start test")
//	var ExprIP = regexp.MustCompile(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\:([0-9]+)`)
//	pollURL := "http://www.89ip.cn/tqdl.html?api=1&num=100&port=&address=%E7%BE%8E%E5%9B%BD&isp="
//
//	resp, err := http.Get(pollURL)
//	if err != nil {
//		logger.Error(err.Error())
//		return
//	}
//
//	if resp.StatusCode != 200 {
//		logger.Error(err.Error())
//		return
//	}
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	bodyIPs := string(body)
//	ips := ExprIP.FindAllString(bodyIPs, 100)
//
//	for index := 0; index < len(ips); index++ {
//		ip := new(ipModel.IP)
//		space := strings.TrimSpace(ips[index])
//		split := strings.Split(space, ":")
//		ip.ProxyHost = split[0]
//		atoi, _ := strconv.Atoi(split[1])
//		ip.ProxyPort = atoi
//		ip.ProxyType = "http"
//		logger.Infof("[89IP] ip = %s, type = %s", ip.ProxyHost, ip.ProxyType)
//		result = append(result, ip)
//	}
//
//	logger.Info("89IP done.")
//	return
//}
