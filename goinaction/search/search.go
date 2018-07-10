package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	// 阻止go应用退出
	// go的main函数返回时，应用就会退出了。这个和Java不一样
	var waitGroup sync.WaitGroup
	// 设置waitGroup的值是需要处理的feed数。每个feed有一个goroutine处理
	// 类似于Java的CountDownLatch
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 运行goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "Matcher")

	matchers[feedType] = matcher
}
