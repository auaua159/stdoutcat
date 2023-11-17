# stdoutcat
This is the bot send stdout as message to discord channel.
## Getting Started
please clone

```
git clone git@github.com:auaua159/stdoutcat.git
```

create `stdoutcat/config.json` and write like below

```
{
    "Token": "YOUR_BOT_TOKEN",
    "ChannelID": "YOUR_CHANNEL_ID"
}
```
build on `stdoutcat/cmd/app`
```
go build -o stdoutcat
```
### make cli command
hard code `var config Config` in `main.go` like below and build
```
config := Config{
		Token:     "YOUR_BOT_TOKEN",
		ChannelID: "YOUR_CHANNEL_ID",
	}
```
move bin file
```
mv stdoutcat /usr/local/bin
```

## Usage
use command with piped stdout 

example
```
echo "Hello, Discord!" | ./stdoutcat
```

when maked cli command
```
echo "Good Bye, Discord!" | stdoutcat
```