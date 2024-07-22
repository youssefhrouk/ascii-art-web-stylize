# Ascii-Art-Web

## Description
Ascii-Art-Web is a web application that converts text input into ASCII art. Designed to explore web development concepts using Go, users can input text, choose a banner style, and generate ASCII art through a simple web interface.

## Features
- Convert input text to ASCII art
- Multiple banner styles: standard, shadow, and thinkertoy
- Web-based user interface
- Error handling for invalid inputs and requests

## Technologies Used
- Backend: Go (Golang)
- Frontend: HTML
- HTTP Server: Go's net/http package
- Templates: Go's html/template package

<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Ascii-Art-Web Project Structure</title>
<style>
    pre {
        font-family: monospace;
        font-size: 14px;
    }
</style>
<body>
    <h1>Ascii-Art-Web Project Structure</h1>
    <pre>
ascii-art-web/
├── main.go
├── Handlers/
│   ├── AsciiHandler.go
│   └── IndexHandler.go
├── ascii/
│   ├── PrintAndSplit.go
│   ├── Converter.go
│   ├── FileReader.go
│   ├── GetBannerPath.go
│   ├── HandleNewLines.go
│   ├── Printer.go
│   └── ValidateInput.go
└── templates/
    ├── index.html
    └── ascii-art.html
    </pre>
</body>

## Usage
1. Run the application:
2. Open a web browser and visit `http://localhost:8080`
3. Enter text in the input field, select a banner style, and click "Generate ASCII Art" to create ASCII art

## API Endpoints
- GET `/`: Displays the main page with the input form
- POST `/ascii-art`: Generates and displays the ASCII art

## Error Handling
The application includes error handling for various scenarios:
- 400 Bad Request: For invalid form inputs
- 404 Not Found: For non-existent pages
- 405 Method Not Allowed: For incorrect HTTP methods
- 500 Internal Server Error: For server-side issues

## Authors
- @asoudri
- @yhrouk