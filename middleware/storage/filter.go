package storage

import (
	"net/http"
	"net/url"
	"proxypool-go/models/ipModel"
	"proxypool-go/util/randomutil"
	"sync"
	"time"

	logger "github.com/sirupsen/logrus"
)

// CheckProxy .
func CheckProxy(ip *ipModel.IP) {
	if CheckIp(ip) {
		ipModel.SaveIp(ip)
	}
}

// CheckIP is to check the ip work or not
func CheckIp(ip *ipModel.IP) bool {
	// 检测代理iP访问地址
	var testIp string
	var testUrl string
	if ip.Type1 == "https" {
		testIp = "https://" + ip.Data
		testUrl = "https://httpbin.org/get"
	} else {
		testIp = "http://" + ip.Data
		testUrl = "http://httpbin.org/get"
	}
	// 解析代理地址
	proxy, parseErr := url.Parse(testIp)
	if parseErr != nil {
		logger.Errorf("parse error: %v\n", parseErr.Error())
		return false
	}
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(10),
	}
	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	begin := time.Now() //判断代理访问时间
	// 使用代理IP访问测试地址
	res, err := httpClient.Get(testUrl)
	if err != nil {
		logger.Warnf("testIp: %s, testUrl: %s: error msg: %v", testIp, testUrl, err.Error())
		return false
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		// 判断是否成功访问，如果成功访问StatusCode应该为200
		speed := time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
		ip.Speed = speed
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
			if !CheckIp(&ip) {
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
	if ipCount == 0 {
		logger.Warnf("RandomProxy random count: %d\n", ipCount)
		return ipModel.IP{}
	}
	randomNum := randomutil.RandInt(0, ipCount)
	return ips[randomNum]
}

// RandomByProxyType .
func RandomByProxyType(proxyType string) (ip ipModel.IP) {
	ips, err := ipModel.GetIpByProxyType(proxyType)
	if err != nil {
		logger.Warn(err.Error())
		return ipModel.IP{}
	}
	ipCount := len(ips)
	if ipCount == 0 {
		logger.Warnf("RandomByProxyType random count: %d\n", ipCount)
		return ipModel.IP{}
	}
	randomNum := randomutil.RandInt(0, ipCount)
	return ips[randomNum]
}
