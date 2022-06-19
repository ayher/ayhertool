package skiplist

import (
	"math/rand"
)

type zskiplist struct {
	header *zskiplistNode
	tail   *zskiplistNode
	level  uint64
	lenght uint64
}

type zrangespec struct {
	min, max     float64
	minex, maxex int
}

func (self *zskiplist) ZslInsert(score float64, obj interface{}) *zskiplistNode {
	var update [ZSKIPLIST_MAXLEVEL]*zskiplistNode //新节点插入位置的前一个节点
	var x *zskiplistNode
	var rank [ZSKIPLIST_MAXLEVEL]uint64 // 链表中的第几个 index
	var i, level int64

	x = self.header

	// 找到新增加的节点的每个层数应该在链表中排第几个，从 header 开始，（rank 相加）
	for i = int64(self.level) - 1; i >= 0; i-- {
		if i != int64(self.level)-1 {
			rank[i] = rank[i+1]
		}

		for x.level[i].forword != nil && (x.level[i].forword.score < score ||
			(x.level[i].forword.score == score && compareStringObjects(x.level[i].forword.robj, obj) < 0)) {
			rank[i] += x.level[i].span
			x = x.level[i].forword
		}

		update[i] = x
	}

	level = int64(zslRandomLevel())
	if level > int64(self.level) {

		for i = int64(self.level); i < level; i++ {
			rank[i] = 0
			update[i] = self.header
			update[i].level[i].span = self.lenght // 没有后续节点就是最长？
		}

		self.level = uint64(level)
	}

	// 创建新节点
	x = zslCreateNode(uint64(level), score, obj)

	for i = 0; i < level; i++ {
		x.level[i].forword = update[i].level[i].forword
		update[i].level[i].forword = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i]) // ？？？
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i = level; i < int64(self.level); i++ {
		update[i].level[i].span++
	}

	if update[0] == self.header {
		x.backword = nil
	} else {
		x.backword = update[0]
	}

	if x.level[0].forword != nil {
		x.level[0].forword.backword = x
	} else {
		self.tail = x
	}

	self.lenght++
	return x
}

// compareStringObjects ???为什么要用他？
func compareStringObjects(a interface{}, b interface{}) int64 {
	return -1
}

const (
	P        = 0.25
	MaxLevel = 32
)

func zslRandomLevel() uint64 {
	level := uint64(1)
	for random() < P && level < MaxLevel {
		level++
	}
	return level
}

func random() float64 {
	return float64(rand.Intn(10)) / 10
}

func (self *zskiplist) ZslGetRank(score float64) (backword *zskiplistNode, forword *zskiplistNode) {
	if self.lenght <= 0 || self.header == nil || self.tail == nil {
		return nil, nil
	}

	backword = self.header
	var lev = self.level - 1
	for {
		if backword == nil {
			return
		}

		if backword.score == score {
			return backword, backword.level[lev].forword
		} else if backword.score > score {
			forword = backword
			backword = backword.backword
		} else if backword.score < score {
			if backword.level[lev].forword == nil {
				if lev != 0 {
					lev--
					backword = backword.level[lev].forword
				} else {
					return backword, backword.level[0].forword
				}
			} else if backword.level[lev].forword.score <= score {
				backword = backword.level[lev].forword
			} else if backword.level[lev].forword.score > score && lev != 0 {
				lev--
				backword = backword.level[lev].forword
			} else {
				return backword, backword.level[lev].forword
			}
		}
	}
}

func (self *zskiplist) ZslGetElementByRank(rank float64) *zskiplistNode {
	var x *zskiplistNode
	var tracersed = uint64(0)

	x = self.header
	for i := self.level - 1; i >= 0; i-- {

		for x.level[i].forword != nil && float64(tracersed+x.level[i].span) <= rank {
			tracersed += x.level[i].span
			x = x.level[i].forword
		}

		if float64(tracersed) == rank {
			return x
		}
	}

	return nil
}

func (self *zskiplist) zslIsInRange(range_var *zrangespec) bool {
	var x *zskiplistNode

	if range_var.min > range_var.max ||
		(range_var.min == range_var.max && (range_var.minex == 0 || range_var.maxex == 0)) {
		return false
	}

	x = self.tail
	if x == nil || !zslValueGteMin(x.score, range_var) {
		return false
	}

	x = self.header.level[0].forword
	if x == nil || !zslValueLteMax(x.score, range_var) {
		return false
	}

	return true
}

func (self *zskiplist) ZslFirstInRange(range_var *zrangespec) *zskiplistNode {
	var x *zskiplistNode
	var i int

	if !self.zslIsInRange(range_var) {
		return nil
	}

	x = self.header
	for i = int(self.level - 1); i >= 0; i-- {
		for x.level[i].forword != nil &&
			!zslValueGteMin(x.level[i].forword.score, range_var) {
			x = x.level[i].forword
		}
	}

	x = x.level[0].forword
	if x == nil {
		return nil
	}

	if !zslValueLteMax(x.score, range_var) {
		return nil
	}
	return x
}

func (self *zskiplist) ZslFree(zsl *zskiplist) {
	// 不需要释放，别抢 gc 的工作。
}

const ZSKIPLIST_MAXLEVEL = 32

func ZslCreate() *zskiplist {
	var zsl = &zskiplist{}

	zsl.level = 1
	zsl.lenght = 0

	zsl.header = zslCreateNode(ZSKIPLIST_MAXLEVEL, 0, nil)
	for i := 0; i < ZSKIPLIST_MAXLEVEL; i++ {
		zsl.header.level[i].forword = nil
		zsl.header.level[i].span = 0
	}
	zsl.header.backword = nil

	zsl.tail = nil

	return zsl
}
