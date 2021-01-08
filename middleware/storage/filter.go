package storage

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net"
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
	if ip.ProxyType == "https" {
		testIp = fmt.Sprintf("https://%s:%d", ip.ProxyHost, ip.ProxyPort)
		testUrl = "https://httpbin.org/get"
	} else {
		testIp = fmt.Sprintf("http://%s:%d", ip.ProxyHost, ip.ProxyPort)
		testUrl = "http://httpbin.org/get"
	}
	// 解析代理地址
	proxy, parseErr := url.Parse(testIp)
	if parseErr != nil {
		logger.Errorf("parse error: %v\n", parseErr.Error())
		return false
	}
	dialer := &net.Dialer{
		// 限制创建一个TCP连接使用的时间（如果需要一个新的链接）
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	//设置网络传输
	netTransport := &http.Transport{
		DialContext: dialer.DialContext,
		Proxy:       http.ProxyURL(proxy),
		// true为代表开启长连接
		DisableKeepAlives: true,
		MaxConnsPerHost:   20,
		// 是长连接在关闭之前，连接池对所有host的最大链接数量
		MaxIdleConns: 20,
		// 连接池对每个host的最大链接数量(MaxIdleConnsPerHost <= MaxIdleConns,如果客户端只需要访问一个host，那么最好将MaxIdleConnsPerHost与MaxIdleConns设置为相同，这样逻辑更加清晰)
		MaxIdleConnsPerHost: 20,
		// 连接最大空闲时间，超过这个时间就会被关闭
		IdleConnTimeout:       20 * time.Second,
		ResponseHeaderTimeout: time.Second * time.Duration(10),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	_ = http2.ConfigureTransport(netTransport)
	// 创建连接客户端
	httpClient := &http.Client{
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
		ip.ProxySpeed = int(speed)
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

// AllProxy .
func AllProxy() []ipModel.IP {
	ips := ipModel.GetAllIp()
	ipCount := len(ips)
	if ipCount == 0 {
		logger.Warnf("RandomProxy random count: %d\n", ipCount)
		return []ipModel.IP{}
	}
	return ips
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
