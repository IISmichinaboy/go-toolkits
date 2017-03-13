package os

import "fmt"

type OsInfo struct {
	Name	     string // 简称 /etc/os-release Centos、Windows、Ubuntu。。。。
	Platform     string // 平台 uname : Linux、Windows
	Pretty       string // 全称 /etc/os-release PRETTY_NAME
	Kernel       string // 内核版本 uname -r
	Arch         string // 架构 architecture  ;linux uname -p;windos echo %PROCESSOR_ARCHITECTURE%
	Bit          string // 系统位数 64 或 32
	Version      string // 系統版本 /etc/os-release VERSION_ID
}

func NewOsInfo() (*OsInfo, error) {
	oi := &OsInfo{}
	err := oi.init()
	if err != nil {
		return nil, err
	}
	return oi, nil
}

func (o *OsInfo) String() string {
	return fmt.Sprintf("Name: %s\nPlatform: %s\nPretty: %s\nKernel: %s\nArch: %s\nBit: %s\nVersion: %s",
		o.Name, o.Platform, o.Pretty, o.Kernel, o.Arch, o.Bit, o.Version)
}
