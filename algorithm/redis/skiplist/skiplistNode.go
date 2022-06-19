package skiplist

type zskiplistNode struct {
	// 后退指针
	backword *zskiplistNode

	// 分数索引
	score float64

	// 值
	robj interface {}

	// 层数
	level []levelStruct
}

func zslCreateNode(level uint64,score float64,obj interface{}) *zskiplistNode {
	return &zskiplistNode{
		backword: nil,
		score: score,
		robj: obj,
		level: make([]levelStruct,level),
	}
}

type levelStruct struct{
	// 前进指针
	forword *zskiplistNode
	// 跨度？？为什么需要这个？
	span uint64
}


