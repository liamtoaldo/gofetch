package main

import (
	"fmt"
	"runtime"
)

func main() {
	goversion := runtime.Version()

	gopherASCII := ""
	gopherASCII = (`

	       ,_---~~~~~----._                                       
	 _,,_,*^____       _____^*,_,,_                           `) + goversion + (`
	/ __/ /'     \    /     '\ \__ \ 
       [  @f | @))    |  | @))   | f@  ]  
        \/    \~____ / __ \_____/    \/   
	|            _l__l_           I   
	}           [______]          I  
	]             |_|_|           |  
	]                             |  
	|                             |   
	|                             |   

	`)

	fmt.Println(gopherASCII)
}