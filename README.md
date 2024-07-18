# SAYT (Simple AI Yield Tool)
![Project Icon](asset/icon.svg)

[中文](asset/README_cn.md)

SAYT is a command-line tool designed to interact with Language Model Models (LLMs) effortlessly. It allows users to communicate with AI assistants via simple commands in the terminal.

## Installation

To install SAYT, use the following command:

```sh
go install github.com/Disdjj/sayt@latest
```

## Initialization

After installation, initialize SAYT with the following command:

```sh
sayt init
```
This command sets up the necessary configurations for SAYT to interact with LLMs.


## Config

**open file `~/sayt/config.toml`**

**input your `api-host`/`api-key`/`model`**

## Init Repo

```shell
sayt repo add basic https://github.com/Disdjj/sayt-basic-prompts.git
```

## List Repo

list Repo to know how to use

```shell
sayt repo list 
sayt repo list category
```

## Usage


```sh
sayt use ai 'hello'
```

```sh
git diff | sayt use summarize_git_diff
```
