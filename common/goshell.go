package common

import (
	"fmt"
	"os/exec"
	"sync"
)

//GoShell 调用command脚本文件
func (snat *SnatValues) GoShell() {
	/* for _, v := range snat.ShellPath {
		cmd := exec.Command("/bin/bash", "-c", v)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Execute Shell:%s failed with error:%s", v, err.Error())
			return
		}
		snat.GetIPFromSnat(string(output))
	} */

	snat.GetIPFromSnat(`spawn ssh -p22 hillstone@10.211.4.253
	hillstone@10.211.4.253's password: 
	SG6000-E5260-B(B)# sh snat resource detail
	IP:211.151.30.249
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (100 of 64512 current available: 64412), port range: 1024-65535
		Udp port pool:
			pool:0% (23 of 64512 current available: 64489), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.112
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (3 of 64512 current available: 64509), port range: 1024-65535
		Udp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.254
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:38% (25021 of 64512 current available: 39491), port range: 1024-65535
		Udp port pool:
			pool:12% (8124 of 64512 current available: 56388), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.251
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (108 of 64512 current available: 64404), port range: 1024-65535
		Udp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.248
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (141 of 64512 current available: 64371), port range: 1024-65535
		Udp port pool:
			pool:0% (21 of 64512 current available: 64491), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:10.211.4.254
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (1 of 64512 current available: 64511), port range: 1024-65535
		Udp port pool:
			pool:0% (64 of 64512 current available: 64448), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.252
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (110 of 64512 current available: 64402), port range: 1024-65535
		Udp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.113
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:8% (5450 of 64512 current available: 59062), port range: 1024-65535
		Udp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	IP:211.151.30.250
		Status: alloced already
		expanded status: no expanded
		Flags: 0x0
		Tcp port pool:
			pool:0% (97 of 64512 current available: 64415), port range: 1024-65535
		Udp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
		Icmp port pool:
			pool:0% (0 of 64512 current available: 64512), port range: 1024-65535
	----------------------------------------------------------------------------
	SG6000-E5260-B(B)# exit
	Connection to 10.211.4.253 closed by remote host.
	Connection to 10.211.4.253 closed.`)
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
