package os

import (
	"syscall"
	"fmt"
	"errors"
	"github.com/IISmichinaboy/go-toolkits/shell"
	"strings"
	"unsafe"
)

type rtlOsVersionInfoW struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformID        uint32
	CSDVersion        [128]byte
}

var  versions = map[string]string{
	"5.0":  "Windows 2000 Professional / Windows 2000 Server ",
	"5.1":  "Windows XP ",
	"5.2":  "Windows XP Professional x64 / Windows Server 2003 ",
	"6.0":  "Windows Vista / Windows Server 2008 ",
	"6.1":  "Windows 7 / Windows Server 2008 R2 ",
	"6.2":  "Windows 8 / Windows Server 2012 ",
	"6.3":  "Windows 8.1 / Windows Server 2012 R2 ",
	"10.0": "Windows 10 / Windows Server 2016 ",
}

func (oi *OsInfo) init() error {
	version, err := getVersion()
	if err != nil {
		return err
	}
	oi.Version = version
	oi.Name,_  = getName(version)
	oi.Pretty  = oi.Name + "Build(" + version+ ")"
	oi.Kernel  = version
	oi.Platform = "Windows"
	if oi.Bit, err = getBit(); err != nil {
		return err
	}
	if oi.Arch, err = getArchitecture(); err != nil {
		return err
	}
	return nil
}

// window version history (https://en.wikipedia.org/wiki/List_of_Microsoft_Windows_versions)
func getName(version string) (string, error) {
	idx := strings.LastIndex(version, ".")
	majorVersion := version[0:idx]
	if name, ok :=  versions[majorVersion]; ok {
		return name, nil
	}
	return "Unkown Windows Version", nil
}

func getArchitecture() (string, error)  {
	session := shell.NewCommand(false, "", "echo %PROCESSOR_ARCHITECTURE%")
	content, err := session.Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error get architecture: %v", err))
	}
	return strings.TrimSpace(string(content)), nil
}

func getBit() (string, error) {
	session := shell.NewCommand(false, "", "wmic cpu get addresswidth")
	content, err := session.Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error get bit: %v", err))
	}
	bit :=  "32"
	if strings.Contains(string(content), "64") {
		bit = "64"
	}
	return bit, nil
}

func getVersion() (string, error) {
	dll := syscall.NewLazyDLL("ntdll.dll")
	proc := dll.NewProc("RtlGetVersion")
	info := &rtlOsVersionInfoW{
		OSVersionInfoSize: 148,
	}
	ret, _, err := proc.Call(uintptr(unsafe.Pointer(info)))
	if ret > 0xC0000000 {
		fmt.Println(errors.New("RtlGetVersion failed: " + err.Error()))
	}
	osVersion := fmt.Sprintf("%d.%d.%d", info.MajorVersion, info.MinorVersion, info.BuildNumber)
	return osVersion, nil
}




