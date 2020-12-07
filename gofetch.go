// +build !windows

package main

import (
	_ "bytes"
	"fmt"
	"gofetch-1/pkgd"
	"os/user"
	"runtime"

	"github.com/fatih/color"
	si "github.com/matishsiao/goInfo"
	"github.com/wille/osutil"
	_"github.com/zcalusic/sysinfo"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	//go version
	_ = color.New(color.FgHiCyan)

	/*\033[34m etc are ANSI COLOR CODES*/

	//detect current user
	currentUser, _ := user.Current()
	user := currentUser.Name
	host := /*blue*/"\033[1;33m" + "Host:      " + "\033[37m" + user + "@" + si.GetInfo().Hostname + "\033[1;34m"

	goversion := "\033[1;33m" + "GO Version:   " + "\033[37m" + runtime.Version() + "\033[1;34m"
	goarch :=    "\033[1;33m" + "GO Arch:      " + "\033[37m" + osutil.GetDisplayArch() + "\033[1;34m"
	os := 		 "\033[1;33m" + "Distro:       " + "\033[37m" + pkgd.GetDist().Description + "\033[1;34m"
	osbased :=   "\033[1;33m" + "Based on:     " + "\033[37m" + osutil.GetDisplay() + "\033[1;34m"
	kernel := 	 "\033[1;33m" + "Kernel:       " + "\033[37m" + si.GetInfo().Core + "\033[1;34m"
	gopkg := 	 "\033[1;33m" + "GO Packages:  " + "\033[37m" + fmt.Sprint(pkgd.PackageCounter()) + "\033[1;34m"
	gocompiler :="\033[1;33m" + "GO Compiler:  " + "\033[37m" + fmt.Sprint(runtime.Compiler) + "\033[1;34m"
	goroot :=    "\033[1;33m" + "Go Root:      " + "\033[37m" + runtime.GOROOT() + "\033[1;34m"
	spacer :=    /*green*/"\033[1;32m" + " -------------------------------" + "\033[1;34m"
	m, _ := mem.VirtualMemory()
	usedmemory := fmt.Sprint(m.Used/1048575, "MiB")
	totalmemory := fmt.Sprint(m.Total/1048575, "MiB")
	memory :=    "\033[1;33m" + "Memory:       " + "\033[37m" + usedmemory +" / " + totalmemory
	colors1 :=   "\033[1;41m   \033[1;0m\033[1;42m   \033[1;0m\033[1;43m   \033[1;0m\033[1;44m   \033[1;0m\033[1;45m   \033[1;0m\033[1;46m   \033[1;0m\033[1;47m   \033[0;0m"
	colors2 :=   "\033[1;41m   \033[0;0m\033[1;42m   \033[0;0m\033[1;43m   \033[0;0m\033[1;44m   \033[0;0m\033[1;45m   \033[0;0m\033[1;46m   \033[0;0m\033[1;47m   \033[0;0m"

	gopherASCII := (/*yellow/orange*/"\033[1;34m" + `

	   ,_---~~~~~----._                                       
    _,,_,*^____       _____^*,_,,_         `) + host + (`
   / __/ /'     \    /     '\ \__ \      `) + spacer + (`
  [  @f | @))    |  | @))   | f@  ]        `) + goversion + (`
   \/    \~____ / __ \_____/    \/         `) + goarch + (`
   |            _l__l_           I         `) + gopkg + (`
   }           [______]          I         `) + gocompiler + (`
   ]             |_|_|           |         `) + goroot + (`
   ]                             |         `) + os + (`
   |                             |         `) + osbased + (`
   |                             |         `) + kernel + (`
   |				 |	   `) + memory + (`
                                         `) + spacer + (`
                                               `) + colors1 + (`
                                               `) + colors2 + (`
`)

	fmt.Println(gopherASCII)
}
