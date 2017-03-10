package os

import "fmt"

type OsInfo struct {
	Name         string //
	Platform     string // 平台 uname : Linux、Windows
	Kernel       string // 内核 uname -r
	Arch         string // CPU架构 architecture  ;uname -p; %PROCESSOR_ARCHITECTURE%
	Version      string // 版本 os-release VERSION_ID
	// distribution  os-release NAME : pretty PRETTY_NAME
}

func (o *OsInfo) String() string {
	return fmt.Sprintf("Platform: %s\nKernel: %s\nArch: %s\nVersion: %s", o.Platform, o.Kernel, o.Arch, o.Version)
}

