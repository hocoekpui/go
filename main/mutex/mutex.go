package mutex

import "sync"

var mutex sync.Mutex

var rwMutex sync.RWMutex

func Mutex() {

	mutex.Lock()
	/*注意解锁需要这样写*/
	defer mutex.Unlock()

	/*不可重入，将会造成死锁*/
	mutex.Lock()
}

func RWMutex() {

	rwMutex.RLocker()
	defer rwMutex.RUnlock()

	/*不可升级，将会造成死锁*/
	rwMutex.Lock()
	defer rwMutex.Unlock()
}
