# json_to_env

The `jsontoenv` package is a Go library designed to simplify the process of setting environment variables from JSON objects. It provides a flexible way to read a JSON object and convert its key-value pairs into environment variables, with options for customization such as omitting specific keys and converting keys to uppercase.

## Features

- **Read JSON Objects**: Easily convert JSON object key-value pairs into environment variables.
- **Omit Specific Keys**: Ability to specify keys that should not be converted into environment variables.
- **Case Customization**: Option to convert all keys to uppercase before setting them as environment variables.

## Getting Started

### Installation

To start using `json_to_env`, install Go and run `go get`:

```bash
go get -u github.com/hummusim/jsontoenv
```

This will retrieve the library.

### Usage

```go
package main

import (
	"log"

	"github.com/hummusim/jsontoenv"
)

func main() {
	// Example JSON data
	jsonData := []byte(`{"username": "admin", "password": "secret"}`)

	// Create a new env instance with options
	env := jsontoenv.New(jsontoenv.Opts{
		UseUpperCase: true, // Convert keys to uppercase
	})

	// Omitting sensitive keys
	env.OmitKeys("password")

	// Setting environment variables from JSON
	if err := env.FromBytes(jsonData); err != nil {
		log.Fatalf("Error setting environment variables: %s", err)
	}
}
```

This example demonstrates how to use the `jsontoenv` library to set environment variables from a JSON object, while omitting sensitive information and converting keys to uppercase.

## Documentation

For more details on the API and its capabilities, refer to the GoDoc documentation.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

`jsontoenv` is licensed under the MIT License. See the LICENSE file for more details.