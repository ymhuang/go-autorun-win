package winutil

import (
    "fmt"
    "os"
    "os/exec"
    "bufio"
    "time"
    "strings"
    "golang.org/x/sys/windows/registry"
//    "golang.org/x/crypto/ssh/terminal"
)

func WaitForEnter() {
    fmt.Println("Press ENTER to continue")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func IsExpired(startTime, endTime string) bool {
    start, _ := time.Parse( time.RFC822, startTime )
    end, _ := time.Parse( time.RFC822, endTime )
    now := time.Now()
/*
    fmt.Println(start)
    fmt.Println(end)
    fmt.Println(now)
*/
    if ( false == (now.After(start) && now.Before(end)) ) {
        return true
    } else {
        return false
    }
}

func CheckPassword(defaultPassword, inputPassword string) (string, bool) {
    var retVal bool = false
    if ( inputPassword == "" ) {
        fmt.Print("Enter password: ")

        keyboardReader := bufio.NewScanner( os.Stdin )
        for ( keyboardReader.Scan() ) {
            inputPassword = keyboardReader.Text()
            if ( strings.Compare( inputPassword, defaultPassword ) == 0 ) {
                retVal = true
            } else {
                retVal = false
            }
            break
        }

/*
        password, err := terminal.ReadPassword(0)
        if err != nil {
            fmt.Println(err.Error())
            retVal = false
        } else {
            inputPassword = string(password)
            if ( strings.Compare( inputPassword, defaultPassword ) == 0 ) {
                retVal = true
            } else {
                retVal = false
            }
        }
*/
    } else {
        if ( strings.Compare( inputPassword, defaultPassword ) == 0 ) {
            retVal = true
        } else {
            retVal = false
        }
    }

    return inputPassword, retVal
}

func CreateAutoRun( prog string ) {
    key, err := registry.OpenKey(registry.LOCAL_MACHINE,
        "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\RunOnce",
        registry.ALL_ACCESS|registry.WOW64_64KEY)
    if (err != nil) {
        fmt.Println(err.Error())
        os.Exit(0)
    }
    defer key.Close()

    err = key.SetStringValue("AbtAutoRun", prog )
    if (err != nil) {
        fmt.Println(err.Error())
        os.Exit(0)
    }
}

func Reboot() {
    _, err := exec.Command("cmd", "/C shutdown /r -t 00").Output()
    if (err != nil) {
        fmt.Println(err.Error())
    }
}

