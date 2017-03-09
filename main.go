package main

import (
	"fmt"
	"github.com/IISmichinaboy/go-toolkits/os"
)

func main()  {
	osInfo := new(os.OsInfo)
	osInfo.Init()
	fmt.Println(osInfo)
}
