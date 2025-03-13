# QR Code Generator

A simple command-line tool to generate QR codes from text input.

## Features
- Quickly generate QR codes from command line arguments
- Creates a PNG image file with the encoded data
- Uses the reliable go-qrcode library

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/qr-generator.git

# Navigate to the project directory
cd qr-generator

# Build the project
go build -o qrgen
```

## Usage

```bash
./qrgen "Your text here"
```

This will generate a file named `qrcode.png` in the current directory containing your encoded text.

## Examples

```bash
# Generate a QR code with a URL
./qrgen "https://example.com"

# Generate a QR code with contact information
./qrgen "MECARD:N:Smith,John;TEL:123456789;EMAIL:john@example.com;;"
```

## Dependencies
- [github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode)

## License
MIT
