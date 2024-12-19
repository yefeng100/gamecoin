package configvip

import (
	"github.com/topfreegames/pitaya/v2"
	"project/base/dao/daobase"
	"project/base/entity"
	"sync"
	"time"
)

var (
	_confVip *ConfigVip
	_once    sync.Once
)

type ConfigVip struct {
	mu      sync.RWMutex
	vipList []*entity.ConfigVip
}

func NewConfigVip() {
	_once.Do(func() {
		_confVip = &ConfigVip{}
		_confVip.init()
	})
}

func Ins() *ConfigVip {
	if _confVip == nil {
		NewConfigVip()
	}
	return _confVip
}

func (t *ConfigVip) init() {
	t.initData()
	//每2分钟执行一次
	pitaya.NewTimer(time.Minute*2, t.initData)
}

func (t *ConfigVip) initData() {
	vipList := daobase.GetConfigVipList()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.vipList = make([]*entity.ConfigVip, 0)
	t.vipList = append(t.vipList, vipList...)
}

func (t *ConfigVip) GetVipList() []*entity.ConfigVip {
	t.mu.RLock()
	defer t.mu.RUnlock()
	vipList := make([]*entity.ConfigVip, len(t.vipList))
	//拷贝数据, 如果直接返回t.vipList，那么返回出去的是指针，在外面对这些数据进行读取或操作，都不会有锁保护
	copy(vipList, t.vipList)
	return vipList
}

func (t *ConfigVip) GetVipExp(level int32) int64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, v := range t.vipList {
		if v.Level == level {
			return v.Exp
		}
	}
	return 0
}
