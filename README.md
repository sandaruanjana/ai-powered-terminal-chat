# Ai Powered Terminal Chat

This is a simple go-based cli program using OpenAI's GPT-3 API to create a chatbot that can be used in a terminal. 


## Run Locally

Clone the project

```bash
  git clone https://github.com/sandaruanjana/ai-powered-terminal-chat.git

```

Go to the project directory

```bash
  cd ai-powered-terminal-chat
```

Setup OpenAI Key

```bash
    inside the main.go file, set the OPENAI_KEY variable to your OpenAI API key
```


Run the application
    
```bash
  go run main.go
```

Build the application

```bash
  go build main.go
```

There Two ways to use the application

1. Run the application directly

```bash
  ./main "Your question here"
```

2. Run the application with the command line arguments

```bash
  ./main
  Ask Luna: "Your question here"
```

## Usage/Examples

```bash
  Ask Luna: Sugget good movies to watch
    1. The Notebook
    2. The Shawshank Redemption
    3. Forrest Gump
    4. The Green Mile
    5. The Rainmaker
    6. The English Patient
    7. The Pianist
    8. The Intouchables
    9. The Fault in Our Stars
    10. The Book Thief
```

```bash
  Ask Luna: tell me how to make http request in golang
  
  There are several ways to make HTTP requests in built-in `net/http` package. Here's an example of how to make a GET request to an API endpoint:
```
``` go
package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
)

func main() {
        // Define the URL to make the request to
        url := "https://api.example.com/endpoint"

        // Make the GET request
        resp, err := http.Get(url)
        if err != nil {
                fmt.Println("Error making request:", err)
                return
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Println("Error reading response body:", err)
                return
        }

        // Print the response body
        fmt.Println(string(body))
}
```

