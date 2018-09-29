package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 演示Go中并发控制的几种方式

// 使用WaitGroup
func Test_waitgroup_1(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("大家都干完了，可以继续了")
}

// 使用select + chan
func Test_select_chan_1(t *testing.T) {

	stop := make(chan bool) // 无缓冲的chan

	// 起一个新的goroutine
	go func() {
		for {
			select {
			case <-stop: // 只有当stop chan中有数据时，case才会命中
				fmt.Println("监控退出,停止了...")
				return // return返回整个函数，goroutine结束
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true // 向stop chan中放数据
	fmt.Println("已通知监控停止")

	time.Sleep(30 * time.Second)
}

func Test_context_1(t *testing.T) {
	// context.Background()返回一个空的Context，一般用于一个Context树的根节点
	// context.WithCancel创建一个可取消的子Context
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 当调用了cancel()后，ctx.Done()的chan有数据
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(30 * time.Second)
}

func Test_context_2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[监控1]")
	go watch(ctx, "[监控2]")
	go watch(ctx, "[监控3]")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(30 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("name"), ", 监控退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value("name"), ", goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

// context.WithValue()的简单用法
func Test_context_value_1(t *testing.T) {
	valueCtx := context.WithValue(context.Background(), "name", "tom")

	go func(ctx context.Context) {
		for {
			fmt.Println("my name is: ", ctx.Value("name"))
			time.Sleep(3 * time.Second)
		}
	}(valueCtx)

	time.Sleep(30 * time.Second)
}

// context.WithCancel, context.WithValue
func Test_context_cancel_value_1(t *testing.T) {
	// 创建了一个可以Cancel的context
	ctx, cancel := context.WithCancel(context.Background())

	// 从上一步中的cancel context，创建一个有Value的子context
	// 添加了一个Key=name，Value=jerry
	valueCtx := context.WithValue(ctx, "name", "jerry")
	go watch2(valueCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(30 * time.Second)
}

func Test_context_deadline_1(t *testing.T) {
	// 创建一个有Deadline的context
	// 可以手动调用cancel方法来停止context
	// 当到达deadline时，会自动停止context
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Local().Add(10*time.Second))
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 当调用了cancel()后，ctx.Done()的chan有数据
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(30 * time.Second)
}

func Test_context_timeout_1(t *testing.T) {
	// 创建一个有Timeout的context
	// 可以手动调用cancel方法来停止context
	// 当到达Timeout时，会自动停止context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go func(ctx context.Context) {
		// 可以在这里加一个defer，defer里面做一些清理的工作
		defer func() {
			// do some cleanup work...
		}()
		for {
			select {
			case <-ctx.Done(): // 当调用了cancel()后，ctx.Done()的chan有数据
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(30 * time.Second)
}
