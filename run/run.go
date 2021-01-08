package run

import (
	logger "github.com/sirupsen/logrus"
	"github.com/wuchunfu/IpProxyPool/fetcher/ip3366"
	"github.com/wuchunfu/IpProxyPool/fetcher/ip66"
	"github.com/wuchunfu/IpProxyPool/fetcher/ip89"
	"github.com/wuchunfu/IpProxyPool/middleware/storage"
	"github.com/wuchunfu/IpProxyPool/models/ipModel"
	"sync"
	"time"
)

func Task() {
	ipChan := make(chan *ipModel.IP, 2000)

	// Check the IPs in DB
	go func() {
		storage.CheckProxyDB()
	}()

	// Check the IPs in channel
	for i := 0; i < 50; i++ {
		go func() {
			for {
				storage.CheckProxy(<-ipChan)
			}
		}()
	}

	// Start getters to scraper IP and put it in channel
	for {
		nums := ipModel.CountIp()
		logger.Printf("Chan: %v, IP: %d\n", len(ipChan), nums)
		if len(ipChan) < 100 {
			go run(ipChan)
		}
		time.Sleep(300 * time.Second)
	}
}

func run(ipChan chan<- *ipModel.IP) {
	var wg sync.WaitGroup
	siteFuncList := []func() []*ipModel.IP{
		ip66.Ip66,
		ip89.Ip89,
		ip3366.Ip33661,
		ip3366.Ip33662,
		//kuaidaili.KuaiDaiLiInha,
		//kuaidaili.KuaiDaiLiIntr,
		//proxylistplus.ProxyListPlus,
	}
	for _, siteFunc := range siteFuncList {
		wg.Add(1)
		go func(siteFunc func() []*ipModel.IP) {
			temp := siteFunc()
			for _, v := range temp {
				ipChan <- v
			}
			wg.Done()
		}(siteFunc)
	}
	wg.Wait()
	logger.Println("All getters finished.")
}
