package main

import (
	"context"
	"fmt"
	"time"
)

// var wg sync.WaitGroup

// func doTask(n int) {
// 	time.Sleep(time.Duration(n))
// 	fmt.Printf("Task %d done\n", n)
// 	wg.Done()
// }

// func main() {
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go doTask(i + 1)
// 	}
// 	wg.Wait()
// 	fmt.Println("All Task Done")
// }

// 等待所有子协程任务全部完成
// Task 3 done
// Task 2 done
// Task 1 done
// All Task Done

// var stop chan bool

// func reqTask(name string) {
// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("stop", name)
// 			return
// 		default:
// 			fmt.Println(name, "send request")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }
// func main() {
// 	stop = make(chan bool)
// 	go reqTask("worker1")
// 	time.Sleep(3 * time.Second)
// 	stop <- true
// 	time.Sleep(3 * time.Second)
// }

// select + chan
// worker1 send request
// worker1 send request
// worker1 send request
// stop worker1

// 更复杂 context.WithCancel
type Options struct{ Interval time.Duration }

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			op := ctx.Value("options").(*Options)
			fmt.Println(name, "send request")
			time.Sleep(op.Interval * time.Second)
		}

	}
}
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	vCtx := context.WithValue(ctx, "options", &Options{1})
	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")
	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(5 * time.Second)
}
