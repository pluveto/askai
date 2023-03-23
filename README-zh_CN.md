# AskAI

![preview_askai](https://user-images.githubusercontent.com/50045289/227156465-ca30161d-4b62-4c7d-bd67-b43dca32c228.gif) 


这是一个名为AskAI的命令行界面（CLI）工具，它使用OpenAI API让您与GPT-3聊天。使用AskAI，您可以提问、寻求帮助或进行交谈。
## 先决条件

在使用AskAI之前，您需要拥有来自OpenAI的API密钥。如果您已经拥有API密钥，可以将其设置为名为`OPENAI_API_KEY`的环境变量。或者，您可以在AskAI可执行文件所在的同一目录中创建一个名为`askai_config.yaml`的文件，并按以下方式添加您的API密钥：

```makefile
api_key: YOUR_API_KEY
```


## 用法

要开始与GPT-3聊天，请运行以下命令：

```bash
./askai
```



这将开始一个交互式会话，在此会话中，您可以输入您的问题或陈述。输入`help`以查看可用命令列表。

您还可以通过将提示作为参数来使用简单模式的AskAI：

```bash
./askai "What is the meaning of life?"
```



这将将GPT-3的响应打印到标准输出。
## 许可证

该工具根据[MIT许可证](https://github.com/sashabaranov/go-openai/blob/master/LICENSE) 获得许可。
