package cases

import (
	"log"
	"sync"
	"time"
)

var currentTime = time.Now().UnixMilli()

type Cases struct{}

func Execute() {
	ca := Cases{}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		ca.case01()
		ca.case02()
		ca.case03()
	}()

	go func() {
		defer wg.Done()
		ca.case04()
		ca.case05()
		ca.case06()
	}()

	go func() {
		defer wg.Done()
		ca.case07()
		ca.case08()
	}()

	wg.Wait()

	log.Println("中宣部认证系统测试用例执行完毕")
}
