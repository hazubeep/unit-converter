# Unit Converter

A web-based unit converter application built with Go that allows users to convert measurements across length, weight, and temperature units. The application provides a simple and intuitive HTML interface for real-time conversion.

Project reference: https://roadmap.sh/projects/unit-converter

## Features

- Convert between metric length units (meter, kilometer, centimeter, millimeter, etc.)
- Convert between metric weight units (kilogram, gram, milligram, etc.)
- Convert between temperature scales (Celsius, Fahrenheit, Kelvin)
- Web-based interface with HTML templates
- Real-time conversion results
- Comprehensive unit testing for all conversion functions

## Supported Conversions

| Category | Units |
|----------|-------|
| Length | meter, kilometer, centimeter, millimeter, decameter, hectometer |
| Weight | kilogram, hectogram, decagram, gram, decigram, centigram, milligram |
| Temperature | Celsius, Fahrenheit, Kelvin |

## Requirements

- Go 1.21 or later

## Installation

Clone the repository:

```bash
git clone https://github.com/hazubeep/unit-converter.git
cd unit-converter
```

Install dependencies (if any):

```bash
go mod tidy
```

## Usage

Run the application with:

```bash
go run main.go
```

The application will start a web server on `http://localhost:8080`. Open this URL in your browser to access the unit converter interface.

Alternatively, build a binary first:

```bash
go build -o bin/unit-converter .
./bin/unit-converter
```

## Testing

Run the test suite to verify all conversion functions:

```bash
go test ./...
```

For verbose test output:

```bash
go test -v ./...
```

## Project Structure

```
unit-converter/
├── converter/
│   ├── length.go       # Length conversion logic
│   ├── length_test.go  # Length conversion tests
│   ├── weight.go       # Weight conversion logic
│   ├── weight_test.go  # Weight conversion tests
│   ├── temperature.go  # Temperature conversion logic
│   └── temperature_test.go # Temperature conversion tests
├── templates/
│   ├── layout.html     # Base HTML template
│   ├── length.html     # Length converter interface
│   ├── weight.html     # Weight converter interface
│   └── temperature.html # Temperature converter interface
├── main.go             # Application entry point
├── handler.go          # HTTP request handlers
├── go.mod              # Go module definition
└── README.md
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
