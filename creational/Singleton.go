package creational
import "sync"

type Singleton struct {
}

var instance *Singleton

//1、懒汉单例模式-非线程安全
func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}   // <--- NOT THREAD SAFE
	}
	return instance
}
//2、带线程锁的单例模式
var once sync.Once

func GetInstance_()*Singleton{
	once.Do(func(){
		instance=&Singleton{}
	})
	return instance
}