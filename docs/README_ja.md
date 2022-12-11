# Fantia DL

[Fantia](https://fantia.jp/)からコンテンツをダウンロードする CLI ツール

## インストール

`fantia-dl`はインストール不要です。
[releases](https://github.com/KleinChiu/fantia-dl/releases)ダウンロードして、そのままパスで使用できます。
あるいは Go のコマンドラインからインストールできます。

```cli
go install github.com/KleinChiu/fantia-dl
```

## 使う方

```sh
Usage:

        fantia-dl <post|backnumber> [arguments]
```

### 投稿

投稿のコンテンツを URL 指定してでダウンロード

```sh
fantia-dl post --url https://fantia.jp/posts/1 --session your_session_string
```

### バックナンバー

バックナンバーのコンテンツを URL 指定してダウンロード

```sh
fantia-dl backnumber --url https://fantia.jp/fanclubs/1/backnumbers?month=201811&plan=1 --session your_session_string
```

> 指定する URL はブラウザの URL 欄にある URL と違う場合があります。右側の「投稿月別」から URL をコピーしてください。
> ![](./docs/backnumber_url.jpg)

## Q&A

### Session ID は何ですか、どこにある？

Session ID はログイン状態を保持する文字列です。
`fantia-dl`はこれを使ってコンテンツをダウンロードします。そして、コンテンツダウンロード以外には使用しません。
一部のコンテンツは会員限定なので、ログインしてからコピーしてください。

Google Chrome / Microsoft Edge の場合、

1. F12 を押して "Developer Tools" を開いてください
2. "Application" タブに移動
3. Cookies > https://fantia.jp から "\_session_id" の値をコピーしてください。

Firefox の場合、

1. F12 を押して "Web Developer Tools" を開いてください
2. "Storage" タブに移動
3. Cookie > https://fantia.jp から "\_session_id" の値をコピーしてください。
