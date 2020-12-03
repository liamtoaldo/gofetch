package main

import (
	_ "bytes"
	"fmt"
	"runtime"
	_"github.com/zcalusic/sysinfo"
	"github.com/fatih/color"
	"github.com/wille/osutil"
	_"os/user"
	"gofetch-1/pkgd"
)

func main() {
	//go version
	_ = color.New(color.FgHiCyan)
	goversion := "GO VERSION:   " + runtime.Version()
	arch :=      "GO ARCH:      " + osutil.GetDisplayArch()
	os :=        "OS:           " + pkgd.GetDist().Description
	osbased :=   "BASED ON:     " + osutil.GetDisplay()
	gopkg :=     "GO PACKAGES:  " + fmt.Sprint(pkgd.PackageCounter())

	gopherASCII := ""
	gopherASCII = (`

	       ,_---~~~~~----._                                       
	 _,,_,*^____       _____^*,_,,_                `) + goversion + (`
	/ __/ /'     \    /     '\ \__ \               `) + arch + (`
       [  @f | @))    |  | @))   | f@  ]               `) + os + (`
        \/    \~____ / __ \_____/    \/                `) + osbased + (`
	|            _l__l_           I                `) + gopkg + (`
	}           [______]          I  
	]             |_|_|           |  
	]                             |  
	|                             |   
	|                             |   

	`)

	fmt.Println(gopherASCII)
}