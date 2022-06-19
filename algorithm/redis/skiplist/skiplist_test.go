package skiplist

import (
	"fmt"
	"testing"
)

func TestSkiplist(t *testing.T) {
	zsl := ZslCreate()
	zsl.ZslInsert(1, "yi")
	zsl.ZslInsert(2, "er")
	zsl.ZslInsert(5, "wu")
	zsl.ZslInsert(6, "liu")
	zsl.ZslInsert(3, "san")
	zsl.ZslInsert(4, "si")
	fmt.Println("===============", zsl)
	fmt.Println(fmt.Sprintf("%+v", zsl.ZslGetElementByRank(1)))
	fmt.Println(fmt.Sprintf("%+v", zsl.ZslFirstInRange(&zrangespec{
		min: 2, max: 3, minex: 0, maxex: 0,
	}).robj))

}
