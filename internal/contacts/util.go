package contacts

import (
    "os"
    "os/exec"
    "runtime"
)

func clearScreenForOS(goos string) {
    var cmd *exec.Cmd
    if goos == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// Production use
func ClearScreen() {
    clearScreenForOS(runtime.GOOS)
}
