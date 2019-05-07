package common

import (
	"fmt"
	"os/exec"
	"sync"
)

//GoShell 调用command脚本文件
func (snat *SnatValues) GoShell() {
	for _, v := range snat.ShellPath {
		cmd := exec.Command("/bin/bash", "-c", v)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Execute Shell:%s failed with error:%s", v, err.Error())
			return
		}
		snat.GetIPFromSnat(string(output))
	}
}

//PushDataPrometheus  k=ip ks=tag vs=tagvalue
func (snat *SnatValues) PushDataPrometheus() {
	snat.IPFirewallValue.Range(func(k, v interface{}) bool {
		switch t := v.(type) {
		case sync.Map:
			t.Range(func(ks, vs interface{}) bool {
				data := fmt.Sprintf(`%v %v %v %v %v`, snat.PushGateWay, k, ks, vs, ks)
				cmd := exec.Command("/bin/bash", "-c", data)
				output, err := cmd.Output()
				if err != nil {
					fmt.Println("Execute Shell:%s failed with error:%s", output, err.Error())
				}
				return true
			})
		default:
		}
		return true
	})
	fmt.Println("success!")
}

//Warp go加锁
func (snat *SnatValues) Warp(cf func()) {
	snat.wg.Add(1)
	go func() {
		cf()
		snat.wg.Done()
	}()
}
