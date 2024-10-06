```markdown
# 🤖 OpenAI Go Module

Welcome to the **OpenAI Go Module**! This Go module allows you to easily interact with any OpenAI model and view your API balance. It's designed to be simple, fast, and easy to integrate into your existing Go projects. 🚀

## Features

- 🎯 **Interact with OpenAI Models**: Send requests to OpenAI models like `GPT-4`, `DALL-E`, or `Whisper` for language generation, image generation, or audio transcription.
- 💸 **Check Your OpenAI API Balance**: Keep track of your usage and view your available balance programmatically.
- ⚡ **Fast and Efficient**: Built with Go's concurrency model for optimal performance.

## Installation

To install the module, use `go get`:

```bash
    go get github.com/StepOne-ai/chatgpt_golang
```

## Usage

### Interacting with OpenAI Models

To interact with OpenAI models, simply import the module and use the relevant functions:

```go
package main

import (
	"fmt"
	openai "github.com/StepOne-ai/chatgpt_golang"
	"bufio"
	"os"
	"strings"
)


func main() {
	client := openai.CreateClient(API_KEY)

	for {
		var prompt string

		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			prompt = strings.Join(words, " ")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		resp, err := openai.GenerateResponse(client, prompt, "gpt-4o-mini")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("> " + resp)
	}
}
```

### Checking Your OpenAI API Balance

You can also check your API balance using this module:

```go
balance, err := openai.GetBalance(API_KEY)
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Printf("Your current balance is: $%.2f\n", balance.balance)
```

## Models

Here is the list of all available models:

![image](https://github.com/user-attachments/assets/318267d6-adc9-4331-ba8b-8be732b23e06)

For more information access https://proxyapi.ru

## Contributing

Feel free to open issues or submit pull requests if you'd like to contribute to the module. Contributions are always welcome! 🎉

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Made with ❤️ and Go
```

This `README.md` includes installation instructions, usage examples, and details about features, helping developers integrate the module easily into their projects. Let me know if you'd like any changes!
