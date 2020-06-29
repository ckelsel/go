// 函数参数 os.Args
package main

import "fmt"
import "os"

// main(int argc, char **argv)
func main() {
    // cmdline :=""
    // seq :=" "

    // for (int i = 0; i < 10; i++) {}
    // for i:=1; i<len(os.Args); i++ {
    //     cmdline += os.Args[i] + seq
    // }
    //
    // fmt.Println(cmdline)

    for _, arg := range os.Args[0:] {
        // fmt.Println(index)
        fmt.Println(arg)
    }
}
