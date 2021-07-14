package lru

import "github.com/ayher/ayhertool/fmt"

type Item struct {
	Value int
	time int
}

type LRUCache struct {
	M map[int]*Item
	Capacity int
	RealUse int
}

func (this *LRUCache)out(){
	for k,v:=range this.M{
		fmt.Debug(k,"-->",v)
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Capacity: capacity,RealUse: 0,M: map[int]*Item{}}
}


func (this *LRUCache) Get(key int) int {
	if v,ok:=this.M[key];ok{
		this.updateLRUCache(key)
		return v.Value
	}else{
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	this.RealUse++
	if v,ok:=this.M[key];ok{
		v.Value=value
		this.updateLRUCache(key)
	}else{
		this.M[key]=&Item{Value: value,time: 0}
		this.updateLRUCache(key)
		if this.RealUse>this.Capacity{
			this.delLRUCache()
		}
	}
}

func (this *LRUCache)updateLRUCache(key int){
	for k,v:=range this.M{
		if k==key{
			v.time=0
		}else{
			v.time++
		}
	}
}

func (this *LRUCache)delLRUCache()  {
	var maxKey,maxTime int
	for k,v:=range this.M{
		if v.time>maxTime{
			maxTime=v.time
			maxKey=k
		}
	}
	if maxKey!=0 && maxTime!=0{
		delete(this.M,maxKey)
		this.RealUse--
	}
}