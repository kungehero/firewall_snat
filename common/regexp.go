package common

import (
	"fmt"
	"regexp"
	"sync"
)

//GetIPFromSnat 根据snat命令,获取ip数组
//((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))
func (snat *SnatValues) GetIPFromSnat(shellReturnValue string) {
	ipregxp := `((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`
	reg := regexp.MustCompile(ipregxp)
	ips := reg.FindAllString(shellReturnValue, -1)
	if len(ips) > 0 {
		snat.GetIPValueFromTCP(ips, shellReturnValue)
	} else {
		fmt.Printf("Execute Shell:ip%s", "is null")
		return
	}
}

//GetIPValueFromTCP 根据ip获取tcp的值
//IP:211.151.30.254+[^)]*
func (snat *SnatValues) GetIPValueFromTCP(ips []string, shellReturnValue string) {
	son := sync.Map{}
	snat.IPFirewallValue = sync.Map{}
	fmt.Println(ips)
	for _, v := range ips {
		fmt.Println(fmt.Sprintf(`IP:%v+[^)]*`, v))
		reg := regexp.MustCompile(fmt.Sprintf(`IP:%v+[^)]*`, v))
		fmt.Println(shellReturnValue)
		arr := reg.FindString(shellReturnValue)
		fmt.Println(arr)
		reg1 := regexp.MustCompile(`\d+`)
		arr1 := reg1.FindAllString(arr, -1)
		fmt.Println(arr1)
		if len(arr1) > 0 {
			son.Store("usedrate", arr1[6])
			son.Store("used", arr1[7])
			son.Store("total", arr1[8])
			son.Store("available", arr1[9])
			snat.IPFirewallValue.Store(v, son)
		} else {
			fmt.Println(len(arr1))
			return
		}

	}
}
