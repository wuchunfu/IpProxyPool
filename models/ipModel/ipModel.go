package ipModel

import (
	logger "github.com/sirupsen/logrus"
	"github.com/wuchunfu/IpProxyPool/middleware/database"
	"github.com/wuchunfu/IpProxyPool/util"
)

// IP struct
type IP struct {
	ProxyId       int64  `gorm:"primary_key; auto_increment; not null" json:"-"`
	ProxyHost     string `gorm:"type:varchar(255); not null; unique" json:"proxyHost"`
	ProxyPort     int    `gorm:"type:int(11); not null; unique" json:"proxyPort"`
	ProxyType     string `gorm:"type:varchar(64); not null" json:"proxyType"`
	ProxyLocation string `gorm:"type:varchar(255); default null" json:"proxyLocation"`
	ProxySpeed    int    `gorm:"type:int(20); not null; default 0" json:"proxySpeed"`
	ProxySource   string `gorm:"type:varchar(64); not null;" json:"proxySource"`
	CreateTime    string `gorm:"type:varchar(50); not null" json:"-"`
	UpdateTime    string `gorm:"type:varchar(50); default ''" json:"updateTime"`
}

// SaveIp 保存数据到数据库
func SaveIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := GetIpByProxyHost(ip.ProxyHost)
	if ipModel.ProxyHost == "" {
		err := db.Model(new(IP)).Create(ip)
		if err.Error != nil {
			logger.Errorf("save ip: %s, error msg: %v", ip.ProxyHost, err.Error)
			db.Rollback()
		}
	} else {
		UpdateIp(ipModel)
	}
	db.Commit()
}

// GetIpByProxyHost 根据 proxyHost 获取一条数据
func GetIpByProxyHost(host string) *IP {
	db := database.GetDB()
	ipModel := new(IP)
	err := db.Model(new(IP)).Where("proxy_host = ?", host).Find(ipModel)
	if err.Error != nil {
		logger.Errorf("get ip: %s, error msg: %v", ipModel.ProxyHost, err.Error)
		return nil
	}
	return ipModel
}

// CountIp 查询共有多少条数据
func CountIp() int64 {
	db := database.GetDB()
	var count int64
	err := db.Model(new(IP)).Count(&count)
	if err.Error != nil {
		logger.Errorf("ip count: %d, error msg: %v", count, err.Error)
		return -1
	}
	return count
}

// GetAllIp 获取所有数据
func GetAllIp() []IP {
	db := database.GetDB()
	list := make([]IP, 0)
	err := db.Model(new(IP)).Find(&list)
	ipCount := len(list)
	if err.Error != nil {
		logger.Warnf("ip count: %d, error msg: %v\n", ipCount, err.Error)
		return nil
	}
	return list
}

// GetIpByProxyType 根据 proxyType 获取一条数据
func GetIpByProxyType(proxyType string) ([]IP, error) {
	db := database.GetDB()
	list := make([]IP, 0)
	err := db.Model(new(IP)).Where("proxy_type = ?", proxyType).Find(&list)
	if err.Error != nil {
		logger.Errorf("error msg: %v\n", err.Error)
		return list, err.Error
	}
	return list, nil
}

// UpdateIp 更新数据
func UpdateIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := ip
	ipMap := make(map[string]interface{}, 0)
	ipMap["proxy_speed"] = ip.ProxySpeed
	ipMap["update_time"] = util.FormatDateTime()
	if ipModel.ProxyId != 0 {
		err := db.Model(new(IP)).Where("proxy_id = ?", ipModel.ProxyId).Updates(ipMap)
		if err.Error != nil {
			logger.Errorf("update ip: %s, error msg: %v", ipModel.ProxyHost, err.Error)
			db.Rollback()
		}
	}
	db.Commit()
}

// DeleteIp 删除数据
func DeleteIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := ip
	err := db.Model(new(IP)).Where("proxy_id = ?", ipModel.ProxyId).Delete(ipModel)
	if err.Error != nil {
		logger.Errorf("delete ip: %s, error msg: %v", ipModel.ProxyHost, err.Error)
		db.Rollback()
	}
	db.Commit()
}
