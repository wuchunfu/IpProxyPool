package storage

import (
	"proxypool-go/models/ipModel"
	"sync"
	"time"

	sj "github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
	logger "github.com/sirupsen/logrus"
)

// CheckProxy .
func CheckProxy(ip *ipModel.IP) {
	if CheckIP(ip) {
		ipModel.SaveIp(ip)
	}
}

// CheckIP is to check the ip work or not
func CheckIP(ip *ipModel.IP) bool {
	var pollURL string
	var testIP string
	if ip.Type2 == "https" {
		testIP = "https://" + ip.Data
		pollURL = "https://httpbin.org/get"
	} else {
		testIP = "http://" + ip.Data
		pollURL = "http://httpbin.org/get"
	}
	logger.Info(testIP)
	begin := time.Now()
	resp, _, errs := gorequest.New().Proxy(testIP).Get(pollURL).End()
	if errs != nil {
		logger.Warnf("[CheckIP] testIP = %s, pollURL = %s: Error = %v", testIP, pollURL, errs)
		return false
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		//harrybi 20180815 判断返回的数据格式合法性
		_, err := sj.NewFromReader(resp.Body)
		if err != nil {
			logger.Warnf("[CheckIP] testIP = %s, pollURL = %s: Error = %v", testIP, pollURL, err)
			return false
		}

		//harrybi 计算该代理的速度，单位毫秒
		ip.Speed = time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
		ipModel.UpdateIp(ip)
		return true
	}
	return false
}

// CheckProxyDB to check the ip in DB
func CheckProxyDB() {
	record := ipModel.CountIp()
	logger.Infof("Before check, DB has: %d records.", record)
	ips := ipModel.GetAllIp()

	var wg sync.WaitGroup
	for _, v := range ips {
		wg.Add(1)
		go func(ip ipModel.IP) {
			if !CheckIP(&ip) {
				ipModel.DeleteIp(&ip)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	record = ipModel.CountIp()
	logger.Infof("After check, DB has: %d records.", record)
}

// RandomProxy .
func RandomProxy() (ip ipModel.IP) {
	ips := ipModel.GetAllIp()

	ipCount := len(ips)
	logger.Warnf("ProxyHttpsRandom ip count: %d\n", ipCount)

	randomNum := RandInt(0, ipCount)
	logger.Infof("RandomProxy random num: %d\n", randomNum)
	if randomNum == 0 {
		return *ipModel.NewIP()
	}
	return ips[randomNum]
}

// RandomByProxyType .
func RandomByProxyType(proxyType string) (ip ipModel.IP) {
	ips, err := ipModel.FindByProxyType(proxyType)
	if err != nil {
		logger.Warn(err.Error())
		return *ipModel.NewIP()
	}
	ipCount := len(ips)
	logger.Warnf("RandomByProxyType ip count: %d\n", ipCount)

	randomNum := RandInt(0, ipCount)
	logger.Infof("RandomByProxyType random num: %d\n", randomNum)
	if randomNum == 0 {
		return *ipModel.NewIP()
	}
	return ips[randomNum]
}
