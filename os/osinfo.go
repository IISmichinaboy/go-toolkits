package os

import "fmt"

type OsInfo struct {
	Platform     string // 平台
	Kernel       string // 内核
	Arch         string //
	Version      string
	Distribution string
	Pretty       string
}

func (o *OsInfo) String() string {
	return fmt.Sprintf("Platform: %s\nKernel: %s\nArch: %s\nDistribution: %s\nVersion: %s\nPretty: %s", o.Platform, o.Kernel, o.Arch, o.Distribution, o.Version, o.Pretty)
}

