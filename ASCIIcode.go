// ASCIIcode prints its command-line arguments' unicode.
// ASCIIcode 输出字符对应的ascii值
package main

import (
	"os"
	"fmt"
)

var ascii = `                            常用ASCII值码点表
HEX   DEC   OCT  Char      HEX   DEC   OCT  Char     HEX   DEC   OCT  Char
000   0     00    \0       111   73    49    I       145   101   65    e
177   127   7F    DEL      112   74    4A    J       146   102   66    f
011   9     09    \t       113   75    4B    K       147   103   67    g
012   10    0A    \n       114   76    4C    L       150   104   68    h
060   48    30    0        115   77    4D    M       151   105   69    i
061   49    31    1        116   78    4E    N       152   106   6A    j
062   50    32    2        117   79    4F    O       153   107   6B    k
063   51    33    3        120   80    50    P       154   108   6C    l
064   52    34    4        121   81    51    Q       155   109   6D    m
065   53    35    5        122   82    52    R       156   110   6E    n
066   54    36    6        123   83    53    S       157   111   6F    o
067   55    37    7        124   84    54    T       160   112   70    p
070   56    38    8        125   85    55    U       161   113   71    q
071   57    39    9        126   86    56    V       162   114   72    r
101   65    41    A        127   87    57    W       163   115   73    s
102   66    42    B        130   88    58    X       164   116   74    t
103   67    43    C        131   89    59    Y       165   117   75    u
104   68    44    D        132   90    5A    Z       166   118   76    v
105   69    45    E        141   97    61    a       167   119   77    w
106   70    46    F        142   98    62    b       170   120   78    x
107   71    47    G        143   99    63    c       171   121   79    y
110   72    48    H        144   100   64    d       172   122   7A    z`


func init(){
    if len(os.Args) != 2 {
        fmt.Println("Usage1: " + os.Args[0] + " xxx")
        fmt.Println("Usage2: " + os.Args[0] + " --all | -a")
        os.Exit(1)
    }
}

// main输出对应字符得unicode码点值
func main() {
    // 输出常用ascii值
    if os.Args[1] == "--all" || os.Args[1] == "-a" {
        fmt.Println(ascii)
        os.Exit(1)
    }

	for i := 0; i < len(os.Args[1:]); i++ {
		for _, c := range os.Args[i+1] {
			fmt.Printf("%q\t%d\n", c, c)
		}
	}
}
