//package main
//
//import (
//	"fmt"
//	"github.com/Seven11Eleven/go-pod-kapot/leetcode"
//	"runtime"
//	"sync"
//)
//
//const goCount = 5555555

//func main() {
//	myMut := mutex_impl.NewMutexWithChan()
//	cnt := 0
//	wg := sync.WaitGroup{}
//
//	wg.Add(goCount)
//
//	for i := 0; i < goCount; i++ {
//		go func() {
//			defer wg.Done()
//			myMut.Lock()
//			cnt++
//			myMut.Unlock()
//		}()
//	}
//
//	wg.Wait()
//
//	fmt.Println("Final count:", cnt)
//}

//func main() {
//	items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
//	queries := []int{1, 2, 3, 4, 5, 6}
//	vam := leetcode.MaximumBeauty(items, queries)
//	fmt.Println(vam)
//
//}

//runtime.GOMAXPROCS(runtime.NumCPU())
//fmt.Println(runtime.NumCPU())
//wg := new(sync.WaitGroup)
//cnt := 0
//
//for i := 0; i < workersCnt; i++ {
//	wg.Add(1)
//
//	go func() {
//		defer wg.Done()
//		cnt++
//	}()
//}
//wg.Wait()
////wg.Done()
//fmt.Print(cnt)

package main

import (
	"fmt"
	"github.com/Seven11Eleven/go-pod-kapot/grep"
)

const workersCnt = 1000000

//func main() {
//	id := "5f2a6c69e1d7a4e0077b4e6b"
//	validId := vgo.(id)
//	fmt.Println(validId) // true
//}

func main() {
	pat := "hello"
	flag := make([]string, 0)
	flagik := "-i"
	flagik2 := "-n"
	flagik3 := "-x"
	flag = append(flag, flagik)
	flag = append(flag, flagik2)
	flag = append(flag, flagik3)
	filepath := "grep/meme.txt"
	filepath1 := "grep/meme2.txt"
	filepathes := make([]string, 0)
	filepathes = append(filepathes, filepath)
	filepathes = append(filepathes, filepath1)

	stroka := grep.Search(pat, flag, filepathes)
	fmt.Println(stroka)
}
