package common

import (
	"fmt"
	"sync"

	"github.com/robfig/cron"
)

type SnatValues struct {
	PushGateWay     string
	ShellPath       []string
	FirewallIPs     []string
	Ips             []string
	IPFirewallValue sync.Map
	Spec            string
	wg              WaitGroupSnat
}
type WaitGroupSnat struct {
	sync.WaitGroup
}

//Corn 定时采集snat
func (snat *SnatValues) Corn() {
	fmt.Println("start successfully!")
	c := cron.New()
	err := c.AddFunc(snat.Spec, func() {
		snat.GoShell()
		snat.PushDataPrometheus()
		//时间间隔执行函数
	})
	if err != nil {
		fmt.Println("faild!", err)
	}
	c.Start()
	select {}
}

//Warp go加锁
func (snat *SnatValues) Warp(cf func()) {
	snat.wg.Add(1)
	go func() {
		cf()
		snat.wg.Done()
	}()
}
