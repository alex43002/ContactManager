package contacts

import (
    "encoding/json"
    "fmt"
    "os"
)

func LoadContacts(filename string) []Contact {
    file, err := os.Open(filename)
    if err != nil {
        return []Contact{}
    }
    defer file.Close()

    var contacts []Contact
    json.NewDecoder(file).Decode(&contacts)
    return contacts
}

func SaveContacts(filename string, contacts []Contact) {
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error saving contacts:", err)
        return
    }
    defer file.Close()
    json.NewEncoder(file).Encode(contacts)
}

