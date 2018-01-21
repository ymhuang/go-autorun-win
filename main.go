package main

import (
    "./winutil"
    "fmt"
    "os"
    "path/filepath"
)

func main() {

    var prog string = ""

    path, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if (err != nil) {
        fmt.Println(err.Error())
        winutil.WaitForEnter()
        return
    }
    prog = fmt.Sprintf("%s\\%s", path, os.Args[0])
    winutil.CreateAutoRun(prog)
    winutil.WaitForEnter()
    winutil.Reboot()
}
