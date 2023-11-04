# SDK Labsmobile SMS Gateway API

[![GoDoc](https://godoc.org/github.com/fsandov/labsmobile-sdk-go?status.svg)](https://godoc.org/github.com/fsandov/labsmobile-sdk-go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

This is a Go SDK for the [Labsmobile SMS Gateway API](https://www.labsmobile.com).

```bash
go get github.com/fsandov/labsmobile-sdk-go
```

## Documentation
All the documentation of the API can be found [here](https://apidocs.labsmobile.com).

## Usage

### Create a client

1. Create a new client with your credentials. You can find your credentials in the [API Credentials](https://websms.labsmobile.com/SY0204/api) section of your Labsmobile account. Its recommended to use environment variables to store your keys.
2. Create a context, you can use the `context.Background()` function. Also you can use the `context.WithTimeout()` function to set a timeout for the request.
3. Create a struct with data to send to the API.

Example:
```go
import "github.com/fsandov/labsmobile-sdk-go/pkg/sms"

func main() {
    // Create a context
    ctx := context.Background()
    // Create a client
    client := sms.NewClient("your-username", "your-password")
    message := sms.Message{
        Destination: "34600000000",
        Message:     "Hello World",
    }
    response, err := client.Send(ctx, message)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
}
```
You can see more examples in the examples folder of this repository.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](LICENSE)


