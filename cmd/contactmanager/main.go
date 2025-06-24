package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/alex43002/ContactManager/internal/contacts"
)

func main() {
    file := "contacts.json"
    contactList := contacts.LoadContacts(file)
    reader := bufio.NewReader(os.Stdin)

    contacts.ClearScreen()
    for {
        fmt.Println("1) Add Contact")
        fmt.Println("2) List Contacts")
        fmt.Println("3) Search Contact")
        fmt.Println("4) Delete Contact")
        fmt.Println("5) Exit\n")
        fmt.Print("Choose an option: ")

        choiceStr, _ := reader.ReadString('\n')
        choiceStr = strings.TrimSpace(choiceStr)
        choice, err := strconv.Atoi(choiceStr)
        if err != nil {
            fmt.Println("Please enter a valid number.")
            fmt.Println("Press enter to continue...")
            reader.ReadString('\n')
            contacts.ClearScreen()
            continue
        }

        switch choice {
        case 1:
            fmt.Print("Enter name: ")
            name, _ := reader.ReadString('\n')
            name = strings.TrimSpace(name)

            fmt.Print("Enter phone: ")
            phone, _ := reader.ReadString('\n')
            phone = strings.TrimSpace(phone)

            newContact := contacts.Contact{Name: name, Phone: phone}
            contactList = append(contactList, newContact)
            contacts.SaveContacts(file, contactList)
            fmt.Println("Contact added.\n")

        case 2:
            if len(contactList) == 0 {
                fmt.Println("No contacts found.\n")
            } else {
                for index, contact := range contactList {
                    fmt.Printf("%d. Name: %s, Phone: %s\n", index+1, contact.Name, contact.Phone)
                }
                fmt.Print("\n")
            }

        case 3:
            fmt.Print("Enter name to search: ")
            searchName, _ := reader.ReadString('\n')
            searchName = strings.TrimSpace(searchName)
            found := false

            for _, contact := range contactList {
                if strings.EqualFold(contact.Name, searchName) {
                    fmt.Printf("Found: Name: %s, Phone: %s\n", contact.Name, contact.Phone)
                    found = true
                }
            }
            if !found {
                fmt.Println("No contacts found with that name")
            }
            fmt.Print("\n")

        case 4:
            if len(contactList) == 0 {
                fmt.Println("No contacts to delete\n")
            } else {
                for index, contact := range contactList {
                    fmt.Printf("%d. Name: %s, Phone: %s\n", index+1, contact.Name, contact.Phone)
                }
                fmt.Print("Enter the number of the contact to delete: ")
                delStr, _ := reader.ReadString('\n')
                delStr = strings.TrimSpace(delStr)
                delIndex, err := strconv.Atoi(delStr)
                if err != nil || delIndex < 1 || delIndex > len(contactList) {
                    fmt.Println("Invalid contact number.\n")
                } else {
                    contactList = append(contactList[:delIndex-1], contactList[delIndex:]...)
                    contacts.SaveContacts(file, contactList)
                    fmt.Println("Contact deleted.\n")
                }
            }

        case 5:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option.\n")
        }

        fmt.Println("Press enter to continue...")
        reader.ReadString('\n')
        contacts.ClearScreen()
    }
}

