package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/victoratsuta/google_map2whatsapp/config"
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

func Execute(container *config.Container) {

	fmt.Println("\nThis program send messages through ur Whatsapp from Google Maps by provided location")
	fmt.Println("\nLets first auth in ur Whatsapp, scan this QR with ur Whatsapp. This step could be skipped if u already logged in")

	err := container.GetWhatsAppService().Auth()
	if err != nil {
		fmt.Printf("❌ Auth error: %v\n", err)
		return
	}
	company := inputSearchingCompany()
	location := inputSearchingLocation()

	companies, _ := container.GetCompaniesRepository().GetByLocation(location + ", " + company)

	printCompaniesDetails(companies)
	fmt.Printf("Total found companies:  %d\n", companies.Count())

	message := inputMessage()
	fmt.Println("\nYour input:")
	fmt.Println(message)
	fmt.Println("\nStarting sending messages to whatsapp")

	err = container.GetWhatsAppService().SendToWhatsApp(companies, message)
	if err != nil {
		fmt.Printf("❌ SendToWhatsApp error: %v\n", err)
		return
	}
}

func printCompaniesDetails(companies entity.CompanyCollection) {
	for _, company := range companies.Get() {
		fmt.Println(company.Name())
		fmt.Println(company.PhoneNumber())
		fmt.Println("_________________")
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

	message := "Hi, how are u?"
	fmt.Printf("Enter message to send (default: %s):\n", message)
	fmt.Println("(Press Ctrl+D or Ctrl+Z on Windows to finish)")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	input := strings.Join(lines, "\n")
	if input != "" {
		message = input
	}
	return message

}
