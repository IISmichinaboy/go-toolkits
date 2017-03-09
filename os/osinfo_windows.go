package os

import (
	"fmt"
	"strings"
	"os/exec"
)

func (o *OsInfo) Init()  {
	out, err := exec.Command("C:\\Windows\\system32\\cmd.exe cmd ver").Output()
	if err != nil {
		fmt.Println(err)
	}
	o.Arch = strings.TrimSpace(string(out))
}


