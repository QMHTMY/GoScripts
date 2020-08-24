// ASCIIcode prints its command-line arguments' unicode.
// ASCIIcode 输出字符对应的ascii值
package main

import (
	"fmt"
	"log"
	"os"
)

// init提前检测参数
func init() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: " + os.Args[0] + " xxx")
	}
}

// main输出对应字符得unicode码点值
func main() {
	for i := 0; i < len(os.Args[1:]); i++ {
		for _, c := range os.Args[i+1] {
			fmt.Printf("%q\t%d\n", c, c)
		}
	}
}
