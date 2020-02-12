package common

import (
	"errors"
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"log"
	"sync"
)

//全局用户账号管理
var NumberHolder *NumberPollHolder

type NumberPollHolder struct {
	m sync.Map
}

type NumberPool struct {
	NO   int64 `gorm:"column:NO"`
	Used int   `gorm:"used"`
}

func (p *NumberPollHolder) NumberPollLoad() {
	numbers := make([]*NumberPool, 0)
	if err := igorm.DB.Model(NumberPool{}).Where("used = 0").Scan(&numbers).Error; err != nil {
		log.Fatalf("load number failed. Err:%v", err)
	}
	//将mysql 中的未使用的账号加载到Map中
	for _, v := range numbers {
		p.m.Store(v.NO, false)
	}
}

//获取账号
func (p *NumberPollHolder) GetNumber() (int64, error) {

	//加redis锁
	if err := iredis.RLock("NUMBER_LOCK", 0); err != nil {
		return 0, errors.New("Redis lock. ")
	}
	defer iredis.RUnlock("NUMBER_LOCK")

	var NO int64
	p.m.Range(func(key, value interface{}) bool {
		NO = key.(int64)
		return true
	})
	if NO != 0 {
		p.m.Delete(NO)
		if err := igorm.DB.Model(NumberPool{}).Where("NO = ?", NO).Update(&NumberPool{Used: 1}).Error; err != nil {
			return 0, err
		}
	}
	return NO, nil
}

//生成账号 1-10000 号
func InitNum() {
	for i := 0; i < 10000; i++ {
		igorm.DB.Model(NumberPool{}).Create(&NumberPool{
			Used: 0,
		})
	}
}
