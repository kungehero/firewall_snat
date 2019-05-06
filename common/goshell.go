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
			fmt.Printf("Execute Shell:%s failed with error:%s", v, err.Error())
			return
		}
		snat.GetIPFromSnat(string(output))
	}
	//fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//PushDataPrometheus  k=ip ks=tag vs=tagvalue
func (snat *SnatValues) PushDataPrometheus() {
	snat.IPFirewallValue.Range(func(k, v interface{}) bool {
		switch t := v.(type) {
		case sync.Map:
			t.Range(func(ks, vs interface{}) bool {
				exec.Command("/bin/bash", "-c", fmt.Sprintf(`%v %v %v %v`, snat.PushGateWay, k, ks, vs))
				return true
			})
		default:
		}
		return true
	})
}
