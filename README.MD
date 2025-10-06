# Google Maps to WhatsApp Sender

## 🚀 Overview

Google Maps to WhatsApp Sender is a powerful CLI tool that automates the process of finding businesses on Google Maps and sending them WhatsApp messages. This tool is perfect for travelers, sales representatives, or anyone who needs to contact multiple businesses efficiently.

### Key Features

- 🔍 Search for businesses on Google Maps
- 💬 Send personalized WhatsApp messages to multiple businesses
- 🏗 Built with Clean Architecture principles
- 🔒 Secure authentication with WhatsApp Web
- ⚡ Fast and efficient batch processing

## 🛠 Technical Stack

- **Language**: Go 1.21+
- **Architecture**: Clean Architecture
- **Key Dependencies**:
  - [Google Maps API](https://developers.google.com/maps/documentation/places/web-service/overview) for business search
  - [whatsmeow](https://github.com/tulir/whatsmeow) for WhatsApp Web interaction
- **Project Template**: This project is based on the [go-clean-template](https://github.com/evrone/go-clean-template) by Evrone

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- Google Maps API key with Places API enabled
  - Get your API key from [Google Cloud Console](https://console.cloud.google.com/)
  - Enable the following APIs:
    - Places API
  - Set up billing (some API calls may incur charges), it is 10k calls per month for free
- WhatsApp Web access on your device

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/google_map2whatsapp.git
   cd google_map2whatsapp
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Copy the example environment file and update with your credentials:
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file and add your Google Maps API key.

## 🚀 Usage

### Basic Usage

```bash
go run main.go
```

The application will guide you through the process:
1. Authenticate with WhatsApp Web by scanning the QR code
2. Enter the type of business you're looking for (e.g., "tour company")
3. Specify the location (city or area)
4. Enter your message
5. The app will find matching businesses and send them your message

### Example

```bash
This program send messages through ur Whatsapp from Google Maps by provided location

Lets first auth in ur Whatsapp, scan this QR with ur Whatsapp. This step could be skipped if u already logged in
█████████████████████████████████████████████████████████████████
█████████████████████████████████████████████████████████████████
████ ▄▄▄▄▄ ███ ▀ ▄█▀▄     █ ▄█ █▀ ▄█    ▄ ▄▀▀▄ █▄█▄ ██ ▄▄▄▄▄ ████
████ █   █ █▄  █▀███▄██ █▄▄█▄█▄▄█▀ ██ ███ ▄▀█▄███▀▄ ██ █   █ ████
████ █▄▄▄█ ██ ████▄█▄█▄ ▀ ▀█▄▀ ▄▄▄ ▀▀▄ █▀▄▄█▄ ▄▀▀ ▀▄██ █▄▄▄█ ████
████▄▄▄▄▄▄▄█ █▄█ ▀ ▀▄█▄█ ▀▄▀▄▀ █▄█ █▄█ █ █ ▀ █▄▀ ▀ █ █▄▄▄▄▄▄▄████
████   ▀▄▄▄  ▀ ▀▄▄ ██  ██▄▄█▄▄▄▄▄▄▄█▄ ▄█▄▄▄▄▄█▄  ▄▄ █▄▄█▀▀▄██████
████▄██  ▀▄████▄█ ▀▄██▄█  ▀ ▄▄▄▄ █ █ █ ▀▀█▀▄ ▄  ▄▄█▀▀▄█ ▀██  ████
████████▄▄▄██ ▄ ▀ █▄██▄█▄ █▄█▀▄  █▀▄▄█▄ ▄▄█  ▄█▄▄█▄███▄  █▄▀▀████
████ ▀▀▄██▄ ▀▄▀█ █▄    ██   ▀ █▀██▄▀██▀█   ▀▀▀ ██▀▀ █▀▀▀ ██▄▄████
████ ▀▄▀█▄▄▀▀▄▄█▀▄▄▄█▀ ██  ▄█ █ ▀█▄▄  █▄  ▄ ▄▀█▄ ▄▄▄▄██▄▄█▀▄ ████
████▄▀█▄█ ▄▀██▄ ██▄█▀▀▀▄▄▀▄▀▀▀▄ ██▄█  █▀  ██▀▀▀ ▀█▀ █▄█▄  ▄▄ ████
████▄▄▀ ▀▀▄▄▄▀▀ ▄█▄█ █  ▄▄ █ ▄▀  █▄ ▄██▄ ▀▄ ▄█▄  █▄ ███▀▄▀██▀████
████▄▀▀█▄▄▄█▄▄▀▀▀▄▀█▀▄█ ▄█▄█▀█▀▀▄█▀▀█▄▀ █▄█▀▄██▀▀ ▄ ▄▀█▀▀▄██▄████
████ ▀▄▀▄▄▄ █▀██▀█ ▄█▄▄▀██ ▄██   ▄█ ▀█▀  ▄▄  ▄▄  █▄▄█▄▄▄ ▀▄ █████
██████   ▄▄▄  ▀▄▀ ██  ▀▄ █▄ █▄ ▄▄▄ ▄▀▀▄█▀▀█ ▄ ▀▀██▄▀ ▄▄▄ █▄  ████
████ ▄▄█ █▄█ ██▀ ▀██▄ █▄██▄▀██ █▄█ ▀  ▄ ▄▀▄▄▄▀▄▄ ▄▄▀ █▄█ ▀█▄▄████
████     ▄▄▄ ▀█▄██▀ ▀▄▀█▀▀ ▀ █▄ ▄   █▄▀█▀▀▀  ▄▀▀▀ ▀█     ▄▄▄▄████
████▀ ▀█▄▀▄▀█▀▄▀█████▄  █ ▄██▀▄▄██▄   ▄▀ █▄▄ ▄█▄ █▄█▄███▄▄▀ ▀████
████▄  ▄▀█▄ ▄▄ ▄▄█▄▀█ ▄▀▄█▄▄█ █▀██▄█ ▀█▀ ▄▄█▄ █  █▄▄ █  ▀ ██▄████
█████▀ ▀▄▄▄  █▄▀▀▀▀ ▄  █    ▄ ▀ ▄▀▀▄▄█▄▀ ▀█▄ ▄█▄▄▀▄▀ █▄▄▄ ▀▀ ████
█████▀▀█▀█▄█  ▀▀█▄█ ▄▀██▀▀█ ▄▀▀ ▄█ ▄▄▀██▀▄▀▄  █▄ ██▀▀ ▄ ▄ ▄ █████
█████ ▄ ▄█▄ ▄▄▄▀▄█▀██   █▄▄ █▀ ▄██▄█▄ ▄ ▄▄▄▀ ██▄▄█▄█▄ ███ ▀▀ ████
████▄▀█▄█ ▄▄ █▄█▀▄█▀▀█▄▄█ ▀▀██  ▄▄ ▀▄▄▄▀▄█ █▄ █▀█▄█▀▀▄▄▄▀▀█▄▄████
████ █▀▀  ▄▄ █▀ ▀▀▄█▄█▄█▄ █▄▄  ▄██▀▄▄▀█▄ ▄█ ▄██▄▄▀▄▀▀  ██ █▀ ████
████ ▀ ▀▀▄▄█ ▀▄▄▄█▀██▄ ▄█▄▀▄▄▀▀▀   █▀▄▀  ▄█▀ ▄█▄▄▄█   ▄██ █▄▄████
██████████▄█▀▀▀▄▀▄ █▄▄ ▄▄▄ █▄█ ▄▄▄   █▄ ▄▄█▄▄██▄▄▄█▀ ▄▄▄  ▀█▄████
████ ▄▄▄▄▄ █ ▄▀█ █▀█▄ ▄▀█▄▀▀▄  █▄█ █▀  ▀▄ █ ███▄ ▀▀▀ █▄█ █▄ ▄████
████ █   █ █▄ ▄▄█ ▀█ ▄  ▄▄ █ █   ▄  ▄█▄ ▄█▄  ▄▄ ▄▄██ ▄▄ ▄▀▄▀▄████
████ █▄▄▄█ █ ▄█▄▀▄▄█▀ █ ▄▀ █▀▀▀▄ █▀▀▄▄ █ ▄ █ ▄██▀▀█▄▄▀▄▄ ▄█▀▄████
████▄▄▄▄▄▄▄█▄█▄███▄▄▄▄▄██▄▄██▄▄█▄█▄▄████▄▄█▄▄███▄▄▄██▄██▄▄█▄▄████
█████████████████████████████████████████████████████████████████
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
Login event: success
Enter company type u are interested in (example: tour company): Tour Agency
Enter location/city you are interested in (example: Milan): Barcelona
Total found companies:  10
Enter message to send (default: Hi, how are u?): Hi!
I’m looking for a travel agency that can arrange private tours for me over three days in different areas.
Could you please tell me what options you offer and how much it would cost?

Thank you!
Your input:
Hi!
I’m looking for a travel agency that can arrange private tours for me over three days in different areas.
Could you please tell me what options you offer and how much it would cost?

Starting sending messages to whatsapp
Message sent successfully to Italy Travel Pro!
Message sent successfully to La Scala Experiences!
Message sent successfully to Navigli Excursions!
Message sent successfully to Milan Tour Experts!
Message sent successfully to Lombardy Adventures!
Message sent successfully to Galleria Guides!
Message sent successfully to Sforza Castle Tours!
Message sent successfully to Duomo Tours!
Message sent successfully to Da Vinci Explorers!
Message sent successfully to Last Supper Visits!

```

## 🏗 Project Structure

```
.
├── cmd/              # CLI commands
├── config/          # Configuration management
├── internal/        # Core application code
│   ├── entity/      # Business entities
│   ├── repository/  # Data access layer
│   └── service/     # Business logic
├── .env.example    # Example environment variables
└── main.go         # Application entry point
```

## 🔒 Environment Variables

Create a `.env` file in the root directory with the following variables:

```
# Environment (test/prod)
ENV=prod

# App
APP_NAME=google_map2whatsapp
APP_VERSION=1.0.0

# Logger
LOG_LEVEL=error

# Google Maps
GOOGLE_MAPS_API_KEY=your_api_key_here
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

This project is based on the [go-clean-template](https://github.com/evrone/go-clean-template) and uses [whatsmeow](https://github.com/tulir/whatsmeow) for WhatsApp Web interaction. Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 Changelog

### [1.0.0] - 2025-10-06
- Initial release
- Basic functionality for finding businesses and sending WhatsApp messages

## 📋 TODO

- [ ] Add tests
- [ ] Check if we do pagination for search results, because now its maximum 2o companies per page
- [ ] Add progress bar for message sending
- [ ] Implement proper error handling and retries

## 🙏 Acknowledgments

- Clean Architecture example by [@evrone](https://github.com/evrone/go-clean-template)
- Google Maps API documentation
- WhatsApp Web API documentation
