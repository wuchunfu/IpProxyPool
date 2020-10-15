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
func InsertIps(ip *IP) {
	db := database.GetDB().Begin()
	err := db.Model(new(IP)).Create(ip)
	if err.Error != nil {
		logger.Error(err.Error)
		db.Rollback()
	}
	db.Commit()
}

func CountIps() int64 {
	db := database.GetDB()
	var count int64
	// set id >= 0, fix bug: when this is nothing in the database
	err := db.Model(new(IP)).Where("id >= ?", 0).Count(&count)
	//err := db.Model(new(IP)).Count(&count)
	if err.Error != nil {
		logger.Error(err.Error.Error())
		return -1
	}
	return count
}

func DeleteIP(ip *IP) {
	db := database.GetDB().Begin()
	err := db.Model(new(IP)).Delete(ip)
	if err.Error != nil {
		db.Rollback()
		logger.Error(err.Error)
	}
	db.Commit()
}

func GetOne(ipStr string) *IP {
	db := database.GetDB()
	//只获取第一条记录
	err := db.Model(new(IP)).Where("data = ?", ipStr).Find(new(IP))
	if err.Error != nil {
		logger.Error(err.Error)
		return nil
	}
	return new(IP)
}

func GetAll() []IP {
	db := database.GetDB()
	list := make([]IP, 0)
	err := db.Model(new(IP)).Find(&list)
	ipCount := len(list)
	if err.Error != nil || ipCount == 0 {
		logger.Warnf("error msg: %v, ip count: %d\n", err.Error, ipCount)
		return nil
	}
	return list
}

//Test if have https proxy in database
//just test on MySQL/Mariadb database
// dbName: ProxyPool
// dbTableName: ip
// select distinct if(exists(select * from ProxyPool.ip where type1='https'),1,0) as a from ProxyPool.ip;
func TestHttps() bool {
	db := database.GetDB()
	err := db.Model(new(IP)).Where(&IP{Type1: "https"}).Find(new(IP))
	if err != nil {
		return false
	}
	return true
}

func FindAll(value string) ([]IP, error) {
	db := database.GetDB()
	list := make([]IP, 0)

	switch value {
	case "http":
		err := db.Model(new(IP)).Where("type1=?", "http").Find(&list)
		if err != nil {
			return list, err.Error
		}
	case "https":
		//test has https proxy on databases or not
		HasHttps := TestHttps()
		if HasHttps == false {
			return list, nil
		}
		err := db.Model(new(IP)).Where("type1=?", "https").Find(&list)
		if err != nil {
			return list, err.Error
		}
	default:
		return list, nil
	}

	return list, nil
}

func Update(ip *IP) {
	db := database.GetDB().Begin()
	ipModel := ip
	ipModel.UpdateTime = util.FormatDateTime()
	err := db.Model(new(IP)).Where("id = ?", 1).Updates(ipModel)
	if err.Error != nil {
		db.Rollback()
		logger.Errorf("[CheckIP] Update IP = %v Error = %v", *ip, err.Error)
	}
	db.Commit()
}
