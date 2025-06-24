package contacts

import (
    "io/ioutil"
    "os"
    "testing"
)

// Already exists, but let's improve:
func TestSaveAndLoadContacts(t *testing.T) {
    testFile := "test_contacts.json"
    defer os.Remove(testFile)

    cs := []Contact{{Name: "Alice", Phone: "123"}}
    SaveContacts(testFile, cs)
    loaded := LoadContacts(testFile)
    if len(loaded) != 1 || loaded[0].Name != "Alice" {
        t.Error("Save/Load did not work as expected")
    }
}

// Test loading when file does not exist
func TestLoadContacts_FileNotFound(t *testing.T) {
    loaded := LoadContacts("nonexistentfile.json")
    if len(loaded) != 0 {
        t.Error("Expected empty contacts slice for missing file")
    }
}

// Test loading when file contains invalid JSON
func TestLoadContacts_InvalidJSON(t *testing.T) {
    tmpfile, err := ioutil.TempFile("", "invalid*.json")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())

    tmpfile.Write([]byte("{invalid json"))
    tmpfile.Close()

    loaded := LoadContacts(tmpfile.Name())
    if len(loaded) != 0 {
        t.Error("Expected empty contacts slice for invalid JSON")
    }
}

// Test SaveContacts error path (cannot create file)
func TestSaveContacts_CannotCreate(t *testing.T) {
    // Try to save to a directory (not a file) to cause an error
    dir, err := ioutil.TempDir("", "contactdir")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(dir)
    SaveContacts(dir, []Contact{{Name: "Bob", Phone: "555"}})
    // No way to assert since it just prints, but this line ensures coverage
}
