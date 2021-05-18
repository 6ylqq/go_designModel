package singleton

import "sync"

type singleton struct{}

var instance *singleton

// GetInstanceUnsafe unSafeSingleton
func GetInstanceUnsafe() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

var mu sync.Mutex

// GetInstanceLock 加锁改串行的方式
func GetInstanceLock() *singleton {
	// 这里锁的粒度还是比较粗，考虑挪到if后形成一个check lock check 结构
	mu.Lock()
	// 释放锁
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func GetInstanceLock2() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

var once sync.Once

// GetInstanceBest 使用go的Once包来实现
func GetInstanceBest() *singleton {
	once.Do(func() {
		// 这里的初始化是安全的
		instance = &singleton{}
	})
	return instance
}
