## GIN API with OpenAI Integration

### Installation

1. **Clone the repository**: Clone the project repository to your local machine.
2. **Install Dependencies**: Navigate to the project directory and run:

```bash
go get github.com/gin-gonic/gin
go get github.com/tmc/langchaingo/llms/openai
```

### Build

```bash
go build
```

### Note: You need an OpenAI API Key to make this work

**Set OpenAI API Key**: Export the OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=Your_OpenAI_API_Key
```
