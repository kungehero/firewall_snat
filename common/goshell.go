package common

import (
	"fmt"
	"os/exec"
	"strings"
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
		str := k.(string)
		fw := strings.Split(str, `,`)
		switch t := v.(type) {
		case sync.Map:
			t.Range(func(ks, vs interface{}) bool {
				data := fmt.Sprintf(`%v %v %v %v %v %v %v`, snat.PushGateWay, fw[0], ks, vs, ks, fw[0], fw[1])
				cmd := exec.Command("/bin/bash", "-c", data)
				output, err := cmd.Output()
				if err != nil {
					fmt.Println("Execute Shell:%s failed with error:%s", output, err.Error())
				}
				return true
			})
		default:
			fmt.Println("不一致啊！")
		}
		return true
	})
}

//Warp go加锁
func (snat *SnatValues) Warp(cf func()) {
	snat.wg.Add(1)
	go func() {
		cf()
		snat.wg.Done()
	}()
}
