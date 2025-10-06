package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/victoratsuta/google_map2whatsapp/config"
)

func Execute(container *config.Container) {

	fmt.Println("\nThis program send messages through ur Whatsapp from Google Maps by provided location")
	fmt.Println("\nLets first auth in ur Whatsapp, scan this QR with ur Whatsapp. This step could be skipped if u already logged in")

	_ = container.GetWhatsAppService().Auth()

	company := inputSearchingCompany()
	location := inputSearchingLocation()

	companies, _ := container.GetCompaniesRepository().GetByLocation(location + ", " + company)

	fmt.Printf("Total found companies:  %d\n", companies.Count())

	message := inputMessage()
	fmt.Println("\nYour input:")
	fmt.Println(message)
	fmt.Println("\nStarting sending messages to whatsapp")

	err := container.GetWhatsAppService().SendToWhatsApp(companies, message)
	if err != nil {
		fmt.Printf("❌ SendToWhatsApp error: %v\n", err)
		return
	}
}

func inputSearchingCompany() string {
	var searchingCompany = "tour company"
	fmt.Printf("Enter company type u are interested in (example: %s): ", searchingCompany)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return searchingCompany
	}

	input = strings.TrimSpace(input)
	if input != "" {
		searchingCompany = input
	}

	return searchingCompany
}

func inputSearchingLocation() string {
	var searchingLocation = "Milan"
	fmt.Printf("Enter location/city you are interested in (example: %s): ", searchingLocation)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return searchingLocation
	}

	input = strings.TrimSpace(input)
	if input != "" {
		searchingLocation = input
	}

	return searchingLocation
}

func inputMessage() string {

	var message = "Hi, how are u?"
	var input string
	fmt.Printf("Enter message to send (default: %s): ", message)

	scanner := bufio.NewScanner(os.Stdin)
	var messageInputLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		messageInputLines = append(messageInputLines, line)
	}

	input = strings.Join(messageInputLines, "\n")

	if input != "" {
		message = input
	}

	return message
}
