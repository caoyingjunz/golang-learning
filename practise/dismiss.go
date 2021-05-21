package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxArriveTime = 60 * 60 * 9
	maxspendTime  = 60 * 30

	allPersions = 20
	works       = 8
)

type Person struct {
	ArriveTime int // 到的时间随机，9点到下午6之间
	SpendTime  int //  处理时间随机，0分钟到30分钟之间
	LeaveTime  time.Time
}

func (p *Person) SetTime() {
}

type empty struct{}

func main() {
	// 设置随机数seed
	rand.Seed(time.Now().UnixNano())

	// 初始化人员，可以是 9 点到 18 之间任意时间到来，每个人的处理时间是10分钟到30分钟
	persons := make([]*Person, 0)
	for i := 0; i < allPersions; i++ {
		aTime := rand.Intn(maxArriveTime)
		sTime := rand.Intn(maxspendTime)
		persons = append(persons, &Person{
			ArriveTime: aTime,
			SpendTime:  sTime,
		})
	}

	// 控制并发数量
	ch := make(chan empty, works)
	var wg sync.WaitGroup

	for _, p := range persons {
		ch <- empty{}
		wg.Add(1)
		go func(p *Person) {
			defer wg.Done()
			p.SetTime()
			<-ch
		}(p)
	}
	wg.Wait()

	for _, p := range persons {
		fmt.Println(*p)
	}

}
