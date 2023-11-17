package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Config 構造体は設定ファイルの形式を定義します。
type Config struct {
	Token     string `json:"Token"`
	ChannelID string `json:"ChannelID"`
}

func main() {
	// 設定ファイルを読み込み
	configFile, err := os.Open("../../config.json")
	if err != nil {
		fmt.Println("設定ファイル読み込みエラー: ", err)
		return
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println("設定ファイル解析エラー: ", err)
		return
	}

	// Discordセッションを作成
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("エラーが発生しました: ", err)
		return
	}

	// Discordに接続
	err = dg.Open()
	if err != nil {
		fmt.Println("接続エラー: ", err)
		return
	}
	defer dg.Close()

	// 標準入力からテキストを読み取り
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// 読み取ったテキストを一つのメッセージとして結合
	message := strings.Join(lines, "\n")

	// メッセージをDiscordに送信
	_, err = dg.ChannelMessageSend(config.ChannelID, message)
	if err != nil {
		fmt.Println("メッセージ送信エラー: ", err)
		return
	}
}
