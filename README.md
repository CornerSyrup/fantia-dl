# Fantia DL

[日本語](./docs/README_ja.md)

CLI tool to download content from [Fantia](https://fantia.jp/)

## Installation

`fantia-dl` does not require any installation.
You can download the executable from [releases](https://github.com/KleinChiu/fantia-dl/releases) and use it directly by specifying the path.
Otherwise, you can install it via go cli

```cli
go install github.com/KleinChiu/fantia-dl
```

## Usage

```sh
Usage:

        fantia-dl <post|backnumber> [arguments]
```

### Post

You can download all contents of a post with its URL

```sh
fantia-dl post --url https://fantia.jp/posts/1 --session your_session_string
```

### Back number

You can download all contents of a back number with its URL

```sh
fantia-dl backnumber --url https://fantia.jp/fanclubs/1/backnumbers?month=201811&plan=1 --session your_session_string
```

> The url to download backnumber is a bit different to the url shown in the URL field sometime. You can copy the url from the right hand side.
> ![](./docs/backnumber_url.jpg)

## Q&A

### Where is the Session ID, how do I find it?

The session ID is a string that represents your login state.
`fantia-dl` needs it to download content from Fantia, and will only use it to download content.
Some contents are member only, please copy the value after you logged in.

For Google Chrome / Microsoft Edge,

1. Press F12 to open "Developer Tools"
2. Switch to the "Application" tab
3. Copy the value of "\_session_id" under Cookies > https://fantia.jp

For Firefox,

1. press F12 to open "Web Developer Tools"
2. Switch to the "Storage" tab
3. Copy the value of "\_session_id" under Cookie > https://fantia.jp
