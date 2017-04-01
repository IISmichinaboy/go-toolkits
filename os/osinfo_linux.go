package os

import (
	"github.com/IISmichinaboy/go-toolkits/shell"
)

func (oi *OsInfo) init() error {
	version, err := getVersion()
	if err != nil {
		return err
	}
	oi.Version = version
	oi.Name, _ = getName(version)
	oi.Pretty = oi.Name + "Build(" + version + ")"
	oi.Kernel, _ = getKernel()
	oi.Platform = "Linux"
	if oi.Bit, err = getBit(); err != nil {
		return err
	}
	if oi.Arch, err = getArchitecture(); err != nil {
		return err
	}
	return nil
}

func getVersion() (string, error) {
	return "", nil
}

func getKernel() (string, error) {
	session := shell.NewCommand(false, "", "uname -r")
	content, err := session.Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error get kernel: %v", err))
	}
	return strings.TrimSpace(string(content)), nil
}

func getArchitecture() (string, error) {
	session := shell.NewCommand(false, "", "uname -p")
	content, err := session.Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error get architecture: %v", err))
	}
	return strings.TrimSpace(string(content)), nil
}
