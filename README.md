# AskAI

![preview_askai](https://user-images.githubusercontent.com/50045289/227156465-ca30161d-4b62-4c7d-bd67-b43dca32c228.gif)


This is a command-line interface (CLI) tool called AskAI that lets you chat with GPT-3 using OpenAI API. With AskAI, you can ask questions, get help, or just have a conversation.
## Prerequisites

Before using AskAI, you need to have an API key from OpenAI. If you already have an API key, you can set it as an environment variable named `OPENAI_API_KEY`. Alternatively, you can create a file named `askai_config.yaml` in the same directory as the AskAI executable and add your API key as follows:

```makefile
api_key: YOUR_API_KEY
```


## Usage

To start chatting with GPT-3, run the following command:

```bash
./askai
```



This will start an interactive session where you can type in your questions or statements. Type `help` to see a list of available commands.

You can also use AskAI in a simple mode by providing a prompt as an argument:

```bash
./askai "What is the meaning of life?"
```



This will print the response from GPT-3 to stdout.
## License

This tool is licensed under the [MIT License](https://github.com/sashabaranov/go-openai/blob/master/LICENSE) .
