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
	gopkg :=     "GO PACKAGES:  " + fmt.Sprint(pkgd.PackageCounter())

	gopherASCII := ""
	gopherASCII = (`

	       ,_---~~~~~----._                                       
	 _,,_,*^____       _____^*,_,,_                `) + goversion + (`
	/ __/ /'     \    /     '\ \__ \               `) + arch + (`
       [  @f | @))    |  | @))   | f@  ]               `) + os + (`
        \/    \~____ / __ \_____/    \/                `) + gopkg + (`
	|            _l__l_           I   
	}           [______]          I  
	]             |_|_|           |  
	]                             |  
	|                             |   
	|                             |   

	`)

	fmt.Println(gopherASCII)
}

// // Distro ...
// type Distro struct {
// 	Display  string
// 	Release  string
// 	Codename string
// 	Description string
// }

// var dists []distroInfo

// type distroInfo struct {
// 	Display string
// 	Search  []string
// 	Files   []string
// }

// func init() {
// 	dists = []distroInfo{
// 		{"Debian", nil, nil},
// 		{"Ubuntu", nil, nil},
// 		{"openSUSE", []string{"SUSE Linux", "openSUSE project"}, []string{"/etc/SuSE-release"}},
// 		{"Mint Linux", []string{"LinuxMint", "mint"}, nil},
// 		{"Gentoo", nil, []string{"/etc/gentoo-release"}},
// 		{"Fedora", nil, []string{"/etc/fedora-release"}},
// 		{"CentOS", nil, []string{"/etc/centos-release"}},
// 		{"Arch Linux", []string{"archlinux", "archarm", "arch"}, nil},
// 		{"Kali Linux", []string{"kali", "debian kali linux"}, nil},
// 	}
// }

// func getLSB() (map[string]string, bool) {
// 	lsb := make(map[string]string)

// 	proc := exec.Command("lsb_release", "-a")
// 	raw, err := proc.Output()

// 	if err != nil {
// 		return nil, false
// 	}

// 	s := string(raw)

// 	for _, line := range strings.Split(s, "\n") {
// 		if line == "" {
// 			break
// 		}

// 		pair := strings.Split(line, ":")
// 		k := pair[0]
// 		v := pair[1]

// 		k = strings.Trim(k, " \t")
// 		v = strings.Trim(v, " \t")

// 		lsb[k] = v
// 	}

// 	return lsb, true
// }

// func getOSRelease() (map[string]string, bool) {
// 	osmap := make(map[string]string)

// 	raw, err := ioutil.ReadFile("/etc/os-release")

// 	if err != nil {
// 		return osmap, false
// 	}

// 	s := string(raw)

// 	for _, line := range strings.Split(s, "\n") {
// 		if line == "" {
// 			break
// 		}

// 		pair := strings.Split(line, "=")
// 		k := pair[0]
// 		v := pair[1]

// 		v = strings.Trim(v, "\"")

// 		osmap[k] = v
// 	}

// 	return osmap, true
// }

// //GetDist is used to 
// func GetDist() Distro {
// 	var detect string
// 	var release string
// 	var codename string
// 	var description string

// 	var info distroInfo

// 	lsb, lsbExists := getLSB()

// 	if lsbExists {
// 		detect = lsb["Distributor ID"]
// 		release = lsb["Release"]
// 		codename = lsb["Codename"]
// 		description = lsb["Description"]
// 	}

// 	osmap, osMapExists := getOSRelease()

// 	if !lsbExists && osMapExists {
// 		distribID := osmap["DISTRIB_ID"]
// 		if distribID != "" {
// 			detect = distribID
// 		}

// 		name := osmap["NAME"]
// 		if distribID == "" && name != "" {
// 			detect = name
// 		}

// 		version := osmap["VERSION_ID"]
// 		if version != "" {
// 			release = version
// 		}

// 		distribRelease := osmap["DISTRIB_RELEASE"]
// 		if distribRelease != "" {
// 			release = distribRelease
// 		}

// 		distribCodename := osmap["DISTRIB_CODENAME"]
// 		if distribCodename != "" {
// 			codename = distribCodename
// 		}
// 		distribDescription := osmap["DISTRIB_DESCRIPTION"]
// 		if distribDescription != "" {
// 			description = distribDescription
// 		}
// 	}

// 	for _, k := range dists {
// 		dl := strings.ToLower(detect)

// 		if strings.ToLower(k.Display) == dl {
// 			info = k
// 			goto ret
// 		}

// 		if k.Search != nil {
// 			for _, search := range k.Search {
// 				if search == dl {
// 					info = k
// 					goto ret
// 				}
// 			}
// 		}

// 		if k.Files != nil {
// 			for _, file := range k.Files {
// 				if _, err := os.Stat(file); err == nil {
// 					info = k
// 					goto ret
// 				}
// 			}
// 		}
// 	}

// ret:

// 	return Distro{info.Display, release, codename, description}
// }