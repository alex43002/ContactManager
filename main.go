package main

import "fmt"
import "os"
import "os/exec"
import "strings"

type Contact struct {
	Name string
	Phone string
}

func clearScreen(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main(){
	
	var contacts [] Contact

	for {
		fmt.Println("1) Add Contact")
		fmt.Println("2) List Contacts")
		fmt.Println("3) Search Contact")
		fmt.Println("4) Exit\n")
		fmt.Print("Choose an option: ")
		

		var choice int
		fmt.Scanln(&choice)

		switch choice {

			case 1:
				var name, phone string
				fmt.Print("Enter name: ")
				fmt.Scanln(&name)

				fmt.Print("Enter phone: ")
				fmt.Scanln(&phone)

				newContact := Contact{Name: name, Phone: phone}
				contacts = append(contacts, newContact)
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
				var searchName string
				fmt.Print("Enter name to search: ")
				fmt.Scanln(&searchName)
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
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Invalid option.\n")



		}

		fmt.Println("Press enter to continue...")
		fmt.Scanln()
		clearScreen()
	}
}
