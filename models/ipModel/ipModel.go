package ipModel

import (
	logger "github.com/sirupsen/logrus"
	"proxypool-go/middleware/database"
	"proxypool-go/util"
)

// IP struct
type IP struct {
	ID         int64  `gorm:"primary_key; auto_increment; not null" json:"-"`
	Data       string `gorm:"type:varchar(50); not null; unique" json:"ip"`
	Type1      string `gorm:"type:varchar(50); not null" json:"type1"`
	Type2      string `gorm:"type:varchar(50); default null" json:"type2,omitempty"`
	Speed      int64  `gorm:"type:int(11); not null; default 0" json:"speed,omitempty"`
	CreateTime string `gorm:"type:varchar(50); not null" json:"createTime"`
	UpdateTime string `gorm:"type:varchar(50); DEFAULT ''" json:"updateTime"`
}

// NewIP .
func NewIP() *IP {
	// init the speed to 100 Sec
	return &IP{
		Speed:      100,
		CreateTime: util.FormatDateTime(),
		UpdateTime: util.FormatDateTime(),
	}
}

//InsertIps SaveIps save ips info to database
func SaveIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := GetOneIp(ip.Data)
	if ipModel.Data == "" {
		err := db.Model(new(IP)).Create(ip)
		if err.Error != nil {
			logger.Errorf("save ip: %s, error msg: %v", ip.Data, err.Error)
			db.Rollback()
		}
	} else {
		UpdateIp(ipModel)
	}
	db.Commit()
}

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

func DeleteIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := ip
	err := db.Model(new(IP)).Delete(ipModel)
	if err.Error != nil {
		logger.Errorf("delete ip: %s, error msg: %v", ipModel.Data, err.Error)
		db.Rollback()
	}
	db.Commit()
}

func GetOneIp(ipStr string) *IP {
	db := database.GetDB()
	ipModel := new(IP)
	//只获取第一条记录
	err := db.Model(new(IP)).Where("data = ?", ipStr).Find(ipModel)
	if err.Error != nil {
		logger.Errorf("get ip: %s, error msg: %v", ipModel.Data, err.Error)
		return nil
	}
	return ipModel
}

func GetAllIp() []IP {
	db := database.GetDB()
	list := make([]IP, 0)
	err := db.Model(new(IP)).Find(&list)
	ipCount := len(list)
	if err.Error != nil || ipCount == 0 {
		logger.Errorf("ip count: %d, error msg: %v\n", ipCount, err.Error)
		return nil
	}
	return list
}

func GetIpByProxyType(proxyType string) ([]IP, error) {
	db := database.GetDB()
	list := make([]IP, 0)
	err := db.Model(new(IP)).Where("type1=?", proxyType).Find(&list)
	if err != nil {
		logger.Errorf("ip list: %v, error msg: %v, \n", list, err.Error)
		return list, err.Error
	}
	return list, nil
}

func UpdateIp(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := ip
	ipMap := make(map[string]interface{}, 0)
	ipMap["update_time"] = util.FormatDateTime()
	err := db.Model(new(IP)).Where("id = ?", ipModel.ID).Updates(ipMap)
	if err.Error != nil {
		logger.Errorf("update ip: %s, error msg: %v", ipModel.Data, err.Error)
		db.Rollback()
	}
	db.Commit()
}
