package shell

import (
	"testing"
	"os"
	"strings"
)

func TestNewCommand(t *testing.T) {
	hostnameForOs, _ := os.Hostname()
	session := NewCommand(false, "", "hostname")
	content, err := session.Output()
	if err != nil {
		t.Error(err)
		return
	}
	hostnameForCMD := strings.TrimSpace(string(content))
	t.Logf("hostnameForOs:%s hostnameForCMD:%s", hostnameForOs, hostnameForCMD)
	if hostnameForOs != hostnameForCMD {
		t.Fail()
	}


}
