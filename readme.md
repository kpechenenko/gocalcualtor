# gocalcualtor

The first AI calculator written in golang.

## Usage

```golang
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/kpechenenko/gocalculator"
    "github.com/openai/openai-go"
    "github.com/openai/openai-go/option"
)

func main() {
    lmStudioClient := openai.NewClient(option.WithBaseURL("http://localhost:1234/v1"))
    
    calc := gocalculator.New(&lmStudioClient, "meta-llama-3.1-8b-instruct")

    res, _ := calc.Calculate(context.Background(), "100 * 2 + 150 - 100 / 2")
    fmt.Printf("Result: %s\n", res)
}
```
