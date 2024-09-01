# cmvalidator

`cmvalidator` is a custom validation library for Go that extends the popular `go-playground/validator/v10` package. This library allows developers to easily specify custom error messages for struct fields using struct tags.

## Features
- Custom Error Messages: Specify custom validation messages directly in struct tags.
- Seamless Integration: Built on top of `go-playground/validator/v10`, making it easy to integrate with existing Go projects.

## Installation
```bash
go get github.com/yudai2929/cmvalidator
```

## Usage
### 1. Define Your Struct
Add custom validation messages to your struct fields using the `customMessage` tag:
```go
package main

import (
    "github.com/yudai2929/cmvalidator"
)

type User struct {
    FirstName string `validate:"required" customMessage:"First Name is required!"`
    LastName  string `validate:"required" customMessage:"Last Name is required!"`
    Age       int    `validate:"gte=0,lte=130" customMessage:"Age must be between 0 and 130!"`
    Email     string `validate:"required,email" customMessage:"Email is invalid!"`
}
```

### 2. Validate Your Struct
Use the `cmvalidator` package to validate your struct:
```go
package main

import (
    "errors"
    "github.com/yudai2929/cmvalidator"
)

func main() {

    // Create a new instance of the validator
    validate := cmvalidator.New()

    // Create a new instance of the struct to validate
    newUser := User{
        FirstName: "John",
        LastName:  "Doe",
        Age:       200, // invalid age
        Email:     "john@example.com",
    }

    // Validate the struct
    if err := validate.Struct(newUser); err != nil {
        var cmvalidateErrors cmvalidator.CFValidateErrors
        if errors.As(err, &cmvalidateErrors) {
            // Handle the error
            for _, cmvalidateError := range cmvalidateErrors {
                // Print the custom message
                println("Custom Message:", cmvalidateError.CustomMessage())
            }
        } else {
            // Handle the error
        }
    }
}
```
[validator.go](validator.go)
### 3. Error Handling
The custom validation errors can be accessed and managed using the `CustomMessage` method:
```go
for _, cmvalidateError := range cmvalidateErrors {
    println("Custom Message:", cmvalidateError.CustomMessage())
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
