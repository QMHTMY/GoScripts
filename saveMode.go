package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var i int = 0x12345678
    const size int = int(unsafe.Sizeof(i))
    ps := (*[size]byte)(unsafe.Pointer(&i))

    fmt.Printf("数字：0x%x\n", i)
    fmt.Printf("类型：%T\n", ps)
    fmt.Printf("其值：%v\n", *ps)

    fmt.Println("地址一：", &ps[0])
    fmt.Println("地址二：", &ps[1])
    fmt.Println("地址三：", &ps[2])
    fmt.Println("地址四：", &ps[3])

    if ps[0] == 0x78 {
        fmt.Println("底层模式：小端模式")
    } else {
        fmt.Println("底层模式：大端模式")
    }
}
