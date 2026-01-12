package bot

import (
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
думал ограничить вызовы, но задержки ТГ апи впринципе съедают всё ненужное,
так эта штука оказалась не нужна, пусть просто полежит.
*/
func NewThrottler(q int) *Throttler {
	var maxQ byte
	if q >= 1 && q <= 255 {
		maxQ = byte(q)
	} else {
		maxQ = 5
	}

	t := &Throttler{
		lock: sync.RWMutex{},
		ttl:  20 * time.Second,
		maxQ: maxQ,
		list: make(map[int64]*userThrottleInfo),
	}

	ticker := time.NewTicker(2 * time.Minute)

	go func() {
		for {
			<-ticker.C
			t.cleanUp()
		}
	}()

	return t
}

type userThrottleInfo struct {
	q        byte
	liveTime time.Time
}

type Throttler struct {
	lock        sync.RWMutex
	list        map[int64]*userThrottleInfo
	maxQ        byte
	ttl         time.Duration
	cleanupFreq time.Duration
}

func (t *Throttler) cleanUp() {
	t.lock.RLock()
	defer t.lock.RUnlock()
	toClean := make([]int64, 0, len(t.list))
	for key, info := range t.list {
		if info.liveTime.Before(time.Now()) {
			toClean = append(toClean, key)
		}
	}

	go func(keyList []int64) {
		t.lock.Lock()
		defer t.lock.Unlock()
		for _, key := range keyList {
			delete(t.list, key)
		}
	}(toClean)

}

func (t *Throttler) CheckUser(user tgbotapi.User) bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	v, ok := t.list[user.ID]
	if !ok {
		info := &userThrottleInfo{
			q:        1,
			liveTime: time.Now().Add(t.ttl),
		}
		t.list[user.ID] = info
		return true
	} else {
		if v.liveTime.After(time.Now()) {
			v.q = 1
			v.liveTime = time.Now().Add(t.ttl)
			t.list[user.ID] = v
			return true
		}
		v.q += 1
		v.liveTime = time.Now().Add(t.ttl)
		return !(v.q > t.maxQ)
	}
}
