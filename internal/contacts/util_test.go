package contacts

import (
    "os"
    "testing"
)

// Test the "windows" branch
func TestClearScreenForOS_Windows(t *testing.T) {
    oldStdout := os.Stdout
    _, w, _ := os.Pipe()
    os.Stdout = w

    clearScreenForOS("windows")

    w.Close()
    os.Stdout = oldStdout
    // Optionally: out, _ := io.ReadAll(r)
}

// Test the non-windows branch (e.g., "linux")
func TestClearScreenForOS_NonWindows(t *testing.T) {
    oldStdout := os.Stdout
    _, w, _ := os.Pipe()
    os.Stdout = w

    clearScreenForOS("linux") // or "darwin", etc.

    w.Close()
    os.Stdout = oldStdout
    // Optionally: out, _ := io.ReadAll(r)
}

// Test the wrapper that calls the real runtime.GOOS (can't branch, but gets coverage)
func TestClearScreen(t *testing.T) {
    oldStdout := os.Stdout
    _, w, _ := os.Pipe()
    os.Stdout = w

    ClearScreen()

    w.Close()
    os.Stdout = oldStdout
    // Optionally: out, _ := io.ReadAll(r)
}

