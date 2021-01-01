// +build windows

package main

import (
	_ "bytes"
	"fmt"
	"gofetch-1/pkgd"
	"os/user"
	"runtime"

	"github.com/fatih/color"
	si "github.com/matishsiao/goInfo"
	"github.com/shirou/gopsutil/mem"
	"github.com/wille/osutil"
)

func main() {
	//go version
	_ = color.New(color.FgHiCyan)

	//detect current user
	currentUser, _ := user.Current()
	user := currentUser.Name
	host := "Host:      " + user + "@" + si.GetInfo().Hostname

	goversion := "GO Version:   " + runtime.Version()
	goarch := "GO Arch:      " + osutil.GetDisplayArch()
	os := "OS:           " + fmt.Sprintf("%v %v", osutil.Name, osutil.GetVersion())
	version := "Version:      " + si.GetInfo().Core
	gopkg := "GO Packages:  " + fmt.Sprint(pkgd.PackageCounter())
	gocompiler := "GO Compiler:  " + fmt.Sprint(runtime.Compiler)
	goroot := "Go Root:      " + runtime.GOROOT()
	spacer := " -------------------------------"
	m, _ := mem.VirtualMemory()
	usedmemory := fmt.Sprint(m.Used/1048575, "MiB")
	totalmemory := fmt.Sprint(m.Total/1048575, "MiB")
	memory := "Memory:       " + usedmemory + " / " + totalmemory
	gopherASCII := (`

	   ,_---~~~~~----._                                       
    _,,_,*^____       _____^*,_,,_         `) + host + (`
   / __/ /'     \    /     '\ \__ \      `) + spacer + (`
  [  @f | @))    |  | @))   | f@  ]        `) + goversion + (`
   \/    \~____ / __ \_____/    \/         `) + goarch + (`
   |            _l__l_           I         `) + gopkg + (`
   }           [______]          I         `) + gocompiler + (`
   ]             |_|_|           |         `) + goroot + (`
   ]                             |         `) + os + (`
   |                             |         `) + version + (`
   |                             |         `) + memory + (`
                                        `) + spacer + (`
`)

	fmt.Println(gopherASCII)
}
