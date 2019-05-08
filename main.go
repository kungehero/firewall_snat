package main

import (
	"firewall_snat/common"
	"flag"
)

func init() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	common.ParseConfig(*cfg)
}

func main() {
	snat := common.SnatValues{PushGateWay: common.Config().PushGateWay, Spec: common.Config().Spec, ShellPath: common.Config().ShellPath, SlaveValues: common.Config().Slave, Filepath: common.Config().Filepath}
	snat.Corn()
}
