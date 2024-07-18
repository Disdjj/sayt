# SAYT（say it）

![Project Icon](asset/icon.svg)

[English](README.md)

SAYT 是一个命令行工具，旨在轻松与语言模型（LLMs）进行交互。它允许用户通过终端中的简单命令与AI助手进行通信。

## 安装

要安装SAYT，请使用以下命令：

```sh
go install github.com/Disdjj/sayt@latest
```

## 初始化

安装后，使用以下命令初始化SAYT：

```sh
sayt init
```

此命令为SAYT与LLMs交互设置必要的配置。

## 配置

**打开文件 `~/sayt/config.toml`**

**输入您的 `api-host`/`api-key`/`model`**

## 初始化仓库

```shell
sayt repo add basic https://github.com/Disdjj/sayt-basic-prompts.git
```

## 列出仓库

列出仓库以了解如何使用

```shell
sayt repo list 
sayt repo list category
```

## 使用方法

```sh
sayt use ai 'hello'
```

```sh
git diff | sayt use summarize_git_diff
```