# stdoutcat
discord チャンネルに標準出力を流し込む bot
## Getting Started
クローンする

```
git clone git@github.com:auaua159/stdoutcat.git
```

ルートディレクトリに `config.json` を追加し、その中身を以下で書き換える(変数の中身は適宜変更)

```
{
    "Token": "YOUR_BOT_TOKEN",
    "ChannelID": "YOUR_CHANNEL_ID"
}
```
`stdoutcat/cmd/app` でビルド
```
go build -o stdoutcat
```
### コマンド化したい場合
`main.go` 内の `var config Config` をトークンとチャンネル ID を埋め込むハードコーディングで書き換えてからビルドしてください
```
config := Config{
		Token:     "YOUR_BOT_TOKEN",
		ChannelID: "YOUR_CHANNEL_ID",
	}
```
バイナリを移動
```
mv stdoutcat /usr/local/bin
```

## Usage
標準出力をパイプして流します

例
```
echo "Hello, Discord!" | ./stdoutcat
```

コマンド化した場合
```
echo "Good Bye, Discord!" | stdoutcat
```