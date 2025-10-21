package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/victoratsuta/google_map2whatsapp/config"
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

func Execute(container *config.Container) {
	greetings()
	whatsappAuth(container)

	company := inputSearchingCompany()
	location := inputSearchingLocation()

	companies, err := container.GetCompaniesRepository().GetByLocation(location + ", " + company)

	if err != nil {
		color.Red("❌ Get from Google maps error: %v\n", err)
		return
	}

	printCompaniesDetails(companies)

	message := inputMessage()

	color.Blue("\nStarting sending messages to whatsapp")

	err = container.GetWhatsAppService().SendToWhatsApp(companies, message)
	if err != nil {
		color.Yellow("\nSome sendings failed.")
		color.Yellow("❌ SendToWhatsApp error: %v\n", err)
		return
	}

	color.Green("\nWe are done!")
}

func greetings() {
	color.Green("Hi!")
	color.Blue("\nThis program send messages through ur Whatsapp from Google Maps by provided location")
	color.Green("---------------------------")
}

func whatsappAuth(container *config.Container) {
	color.Blue("Lets first auth in ur Whatsapp, scan this QR with ur Whatsapp. This step will be skipped if u already logged in")

	err := container.GetWhatsAppService().Auth()
	if err != nil {
		color.Red("❌ Auth error: %v\n", err)
	}

	color.Green("Auth in Whatsapp is done!\n")
	color.Green("---------------------------")
}

func printCompaniesDetails(companies entity.CompanyCollection) {
	color.Green("Total found %d companies\n", companies.Count())
	color.Green("Here is list of found companies\n")

	for _, company := range companies.Get() {
		color.Yellow(company.Name())
		color.Yellow(company.PhoneNumber())
		color.Yellow("_________________")
	}

	color.Green("---------------------------")
}

func inputSearchingCompany() string {
	var searchingCompany = "tour company"
	color.Blue("Enter company type u are interested in (example: %s): ", searchingCompany)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
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
	color.Blue("Enter location/city you are interested in (example: %s): ", searchingLocation)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
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
	color.Blue("Enter message to send (default: %s):\n", message)
	color.Blue("(Press Ctrl+D or Ctrl+Z on Windows to finish)")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	input := strings.Join(lines, "\n")
	if input != "" {
		message = input
	}

	color.Green("---------------------------")

	return message
}
