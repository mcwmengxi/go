package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	Cond()
	// Pool()
	// SyMap()
	// WaitG()
	// Once()
	// LockOption()
}
func LockOption(){
	lock := &sync.Mutex{}
	read_lock := &sync.RWMutex{}
	go LockFunc(lock)
	go LockFunc(lock)
	go ReadLockFunc(read_lock)
	go ReadLockFunc(read_lock)
	go ReadLockFunc(read_lock)
	go ReadLockFunc(read_lock)
	for{}
}
func LockFunc(lock *sync.Mutex){
	// 写的时候排斥其他的写锁和读锁
	lock.Lock()
	fmt.Println("11")
	time.Sleep((1*time.Second))
	lock.Unlock()
}
func ReadLockFunc(r_lock *sync.RWMutex){
	// 在读取的时候不会阻塞其他的读锁但是会排斥掉写锁
	r_lock.RLock()
	fmt.Println("read-11")
	time.Sleep((1*time.Second))
	r_lock.RUnlock()
}

func Once(){
	once := &sync.Once{}
	for i:=1; i<5; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
	for i:=0; i<5; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}

func WaitG (){
	w := &sync.WaitGroup{}
	w.Add(2)
	go func(){
		time.Sleep(8*time.Second)
		w.Done()
		fmt.Println("delta-1")
	}()
	go func(){
		time.Sleep(6*time.Second)
		w.Done()
		fmt.Println("delta-1")
	}()
	w.Wait()
	time.Sleep(4*time.Second)
}

func SyMap(){
	m := &sync.Map{}
	go func(){
		for{
			m.Store(4, 0)
		}
	}()
	go func(){
		for{
			fmt.Println(m.Load(4))
		}
	}()
	time.Sleep(100)
	m.Store(1,2)
	m.LoadOrStore(3,3)
	m.Delete(1)
	fmt.Println(m.LoadOrStore(3,3))
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k,"==>",v)
		time.Sleep(100)
		return true
	})
}

func Pool(){
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	p.Put(5)
	p.Put(6)
	for i:=0; i<8; i++{
		time.Sleep(1 *time.Second)
		fmt.Println(p.Get())
	}
}
func Cond(){
	co := sync.NewCond(&sync.Mutex{})
	// co.L.Lock()
	// co.Wait()
	// co.Wait()
	// co.Wait()
	// co.L.Unlock()
	// co.Signal()
	// co.Broadcast()

	go func(){
		co.L.Lock()
		fmt.Println("Lock1")
		co.Wait()
		fmt.Println("unlock1")
		co.L.Unlock()
	}()
	go func(){
		co.L.Lock()
		fmt.Println("Lock2")
		co.Wait()
		fmt.Println("unlock2")
		co.L.Unlock()
	}()
	time.Sleep(2*time.Second)
	// co.Broadcast()
	co.Signal()
	time.Sleep(1*time.Second)
}