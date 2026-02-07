package service

import (
	"fmt"
	"sync"
)

type HandleFunc func(any) error

type EventBus struct {
	locker      sync.Locker
	subscribers map[string][]HandleFunc
}

func (b *EventBus) Subscribe(name string, callback HandleFunc) {
	b.locker.Lock()
	defer b.locker.Unlock()
	b.subscribers[name] = append(b.subscribers[name], callback)
}

func (b *EventBus) Publish(payload any, name string) {
	b.locker.Lock()
	cbs, ok := b.subscribers[name]
	b.locker.Unlock()

	if !ok {
		return
	}

	errs := make([]error, 0, len(cbs))
	for _, callback := range cbs {
		if err := callback(payload); err != nil {
			errs = append(errs, err)
		}
	}

	for _, e := range errs {
		fmt.Println(e.Error()) // в нормальный логер обернуть
	}

}

func newEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]HandleFunc),
		locker:      &sync.Mutex{},
	}
}
