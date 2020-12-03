package pkgd

import (
	"os/exec"
	//"strings"
)

//PackageCounter counts every Go package you have installed
func PackageCounter() int {
	proc := exec.Command("go", "list", "...")
	raw, _ := proc.Output()
	s := string(raw)

	var i int = 0
	var j int = 0
	for i = 0; i < len(s); i++ {
		//counts every package
		if raw[i] == 10 {
			j++
		}
	}
	
	return j
}
