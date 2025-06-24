package main

import "fmt"
import "os"
import "os/exec"
import "strings"
import "encoding/json"
import "bufio"
import "strconv"

func loadContacts(filename string) []Contact {
	file, err := os.Open(filename)
	if err != nil {
		return []Contact{}
	}
	defer file.Close()

	var contacts []Contact
	json.NewDecoder(file).Decode(&contacts)
	return contacts
}

func saveContacts(filename string, contacts []Contact) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving contacts:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(contacts)
}

func clearScreen(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main(){
	
	file := "contacts.json"
	var contacts [] Contact = loadContacts(file)
	reader := bufio.NewReader(os.Stdin)

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
            clearScreen()
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

				newContact := Contact{Name: name, Phone: phone}
				contacts = append(contacts, newContact)
				saveContacts(file, contacts)
				fmt.Println("Contact added.\n")
			case 2:
				if len(contacts) == 0 {
					fmt.Println("No contacts found.\n")
				} else {
					for index, contact := range contacts {
						fmt.Printf("%d. Name: %s, Phone: %s\n", index+1, contact.Name, contact.Phone)
					}
					fmt.Print("\n")
				}
			case 3:
				fmt.Print("Enter name to search: ")
				searchName, _ := reader.ReadString('\n')
				searchName = strings.TrimSpace(searchName)
				found := false

				for _, contact := range contacts {
					if strings.EqualFold(contact.Name, searchName){
						fmt.Printf("Found: Name: %s, Phone: %s\n", contact.Name, contact.Phone)
						found = true
					}	
				}
				if !found {
					fmt.Println("No contacts found with that name")
				}
				fmt.Print("\n")
			case 4:
				if len(contacts) == 0 {
					fmt.Println("No contacts to delete\n")
				} else {
					for index, contact := range contacts {
						fmt.Printf("%d. Name: %s, Phone: %s\n", index + 1, contact.Name, contact.Phone)
					}
					fmt.Print("Enter the number of the contact to delete: ")
					delStr, _ := reader.ReadString('\n')
					delStr = strings.TrimSpace(delStr)
					delIndex, err := strconv.Atoi(delStr)
					
					if err != nil || delIndex < 1 || delIndex > len(contacts) {
						fmt.Println("Invalid contact number.\n")
					} else {
						contacts = append(contacts[:delIndex - 1], contacts[delIndex:]...)
						saveContacts(file, contacts)

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
        clearScreen()
	}
}
