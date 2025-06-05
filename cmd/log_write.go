package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/AugusTaro/waicli/pkg/logstore"
	"github.com/spf13/cobra"
)

var logWriteCmd = &cobra.Command{
	Use:   "write",
	Short: "ログを書く",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ディレクトリ作成
		file, err := logstore.PrepareTodayLogFile()
		if err != nil {
			return fmt.Errorf("ファイル準備に失敗: %v", err)
		}
		defer file.Close()

		// ファイル作成
		fmt.Println("ファイルを作成する処理をここに書く")

		// 終了を検知する
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-signalChan
			fmt.Println("\n記録を終了しました。")
			os.Exit(0)
		}()
		fmt.Println("終了します")

		// 入力をループで受け取り、ファイルに書き込む
		fmt.Println("ファイルに書き込む処理をここに書く")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				continue
			}

			timestamp := time.Now().Format("15:04:05")
			line := fmt.Sprintf("[%s] %s\n", timestamp, input)
			if _, err := file.WriteString(line); err != nil {
				fmt.Println("書き込みエラー:", err)
			}
		}

		return nil
	},
}

func init() {
	// logCmd に対してサブコマンドとして追加
	logCmd.AddCommand(logWriteCmd)
}
