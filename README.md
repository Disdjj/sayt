# SAYT (Simple AI Yield Tool)
![Project Icon](asset/icon.svg)

[中文](asset/README_cn.md)

SAYT is a command-line tool designed to interact with Language Model Models (LLMs) effortlessly. It allows users to communicate with AI assistants via simple commands in the terminal.

## Installation

To install SAYT, use the following command:

```sh
go install github.com/Disdjj/sayt
```

## Initialization

After installation, initialize SAYT with the following command:

```sh
sayt init
```
This command sets up the necessary configurations for SAYT to interact with LLMs.


## Config
open file `~/sayt/config.toml`
input your `api-host`/`api-key`/`model`

## Usage

To use SAYT and interact with an AI assistant, use the following command format:

```sh
sayt use assistant 'your message here'
```

For example:

```sh
sayt use assistant 'hello'
```

This command will send the message "hello" to the AI assistant and return the response.
