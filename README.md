# hellogpt
CLI application made in Go language to talk to ChatGPT.

## Usage

```sh
API_KEY={your OpenAI API key} go run main.go
```

### Setting the Context for the Conversation

It is also possible to define the context of the conversation in JSON beforehand before executing it.

```json
[
    {"role": "system", "content": "You are a Go programmer"},
    {"role": "user", "content": "Good morning!"},
    {"role": "assistant", "content": "Hi! If you have any problems with the Go language, please ask me anytime!"}
]
```

```sh
API_KEY={your OpenAI API key} go run main.go -c {JSON file path}
```