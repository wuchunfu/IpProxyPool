package run

import (
	logger "github.com/sirupsen/logrus"
	"proxypool-go/getter/ip66"
	"proxypool-go/getter/ip89"
	"proxypool-go/middleware/storage"
	"proxypool-go/models/ipModel"
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
		nums := ipModel.CountIps()
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
		//ip66_01.IP66,
		ip66.IP66, //need to remove it
		//ip3366.IP3306,
		//kuaidaili.KDL,
		//proxylistplus.PLP, //need to remove it
		//proxylistplus.PLPSSL,
		ip89.IP89,
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
