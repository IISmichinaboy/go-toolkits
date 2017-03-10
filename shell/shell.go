package shell

import (
	"runtime"
	"github.com/codeskyblue/go-sh"
)

const (
	cmd   = "c:\\windows\\system32\\cmd.exe"
	shell = "sh"
)

var isWindows bool = runtime.GOOS == "windows"


func NewCommand(login bool, user string, a ...string) *sh.Session {
	args := make([]string, 0)
	var name string

	if isWindows {
		name = cmd
		args = append(args, "/c")
	} else {
		name = shell
		args = append(args, "-c")
		args = append(args, "su")
		// system user
		if user == "" {
			user = "nobody"
		}
		// load env
		if login {
			args = append(args, "-" + user)
		} else {
			args = append(args, "-" + user)
		}
	}

	args = append(args, a...)

	return sh.Command(name, args)
}