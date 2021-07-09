package lru

type KeyCache struct{
	Value int
	Pre *KeyCache
	Nex *KeyCache
}

type LRUCache struct {
	Head *KeyCache
	Tail *KeyCache
	M map[int]*KeyCache
}


func Constructor(capacity int) LRUCache {
	h:=&KeyCache{capacity,nil,nil}
	t:=&KeyCache{0,h,nil}
	h.Nex=t
	return LRUCache{
		h,
		t,
		map[int]*KeyCache{},
	}
}


func (this *LRUCache) Get(key int) int {
	if v,ok:=this.M[key];ok && v.Value!=-1{
		this.setHead(v)
		return v.Value
	}else{
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if v,ok:=this.M[key];ok && v.Value!=-1{
		v.Value=value
		this.setHead(v)
	}else{
		k:=&KeyCache{Value: value,Pre:this.Head,Nex: this.Head.Nex}
		this.M[key]=k
		this.Head.Nex.Pre=k
		this.Head.Nex=k
		this.Tail.Value++
		if this.Tail.Value>this.Head.Value{
			this.delTail()
		}
	}
}

func (this *LRUCache)delTail(){
	this.Tail.Pre.Pre.Nex=this.Tail
	this.Tail.Pre.Value=-1
	this.Tail.Pre=this.Tail.Pre.Pre
	this.Tail.Value--
}

func (this *LRUCache)setHead(v *KeyCache){
	if this.Head.Nex!=v{
		v.Pre.Nex=v.Nex
		v.Nex.Pre=v.Pre
		v.Nex=this.Head.Nex
		v.Pre=this.Head
		this.Head.Nex.Pre=v
		this.Head.Nex=v
	}
}

