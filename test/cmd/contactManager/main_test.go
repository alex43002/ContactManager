package contactManager

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "testing"
    "io"
    "runtime"
)

func TestEndToEndContactManager(t *testing.T) {
    // Path to main.go (relative to project root)
    appPath := filepath.Join("..", "..", "..", "cmd", "contactmanager", "main.go")

    // Build the binary
    binPath := filepath.Join(os.TempDir(), "contactmanager_test_bin")
    if runtime.GOOS == "windows" {
        binPath += ".exe"
    }

    buildCmd := exec.Command("go", "build", "-o", binPath, appPath)
    buildCmd.Stdout = os.Stdout
    buildCmd.Stderr = os.Stderr
    if err := buildCmd.Run(); err != nil {
        t.Fatalf("Failed to build test binary: %v", err)
    }
    defer os.Remove(binPath)

    // Prepare test input (simulate a user adding, listing, then exiting)
    testInput := strings.Join([]string{
        "1",           // Add Contact
        "Test User",   // Name
        "555-5555",    // Phone
        "",            // Press enter to continue
        "2",           // List contacts
        "",            // Press enter to continue
        "5",           // Exit
    }, "\n")

    // Remove contacts.json if it exists to ensure a clean test
    os.Remove("contacts.json")
    defer os.Remove("contacts.json")

    // Start the process
    cmd := exec.Command(binPath)
    stdin, err := cmd.StdinPipe()
    if err != nil {
        t.Fatalf("Failed to get stdin: %v", err)
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        t.Fatalf("Failed to get stdout: %v", err)
    }
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        t.Fatalf("Failed to start binary: %v", err)
    }

    // Write test input in a goroutine (simulate typing)
    go func() {
        io.WriteString(stdin, testInput)
        stdin.Close()
    }()

    // Capture the output
    out, _ := io.ReadAll(stdout)
    cmd.Wait()

    output := string(out)

    // Assert: Did the output contain what we expect?
    if !strings.Contains(output, "Contact added.") {
        t.Error("Did not find 'Contact added.' in output:\n" + output)
    }
    if !strings.Contains(output, "Test User") || !strings.Contains(output, "555-5555") {
        t.Error("Did not find contact in list output:\n" + output)
    }
}
