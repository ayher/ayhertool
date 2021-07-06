package search_test

import (
	"github.com/ayher/ayhertool/fmt"
	"testing"
	"github.com/ayher/ayhertool/algorithm/search"
)

func TestBinarySearch(T *testing.T)  {
	fmt.Println(search.BinarySearch([]int{-1,0,3,5,9,12},9))
	fmt.Println(search.BinarySearch([]int{-1,0,3,5,9,12},2))
	fmt.Println(search.BinarySearch([]int{-1,0,3,5,9,12},13))
	fmt.Println(search.BinarySearch([]int{-1,0,3,5,9,12},-1))
}
