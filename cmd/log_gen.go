package cmd

import (
	"fmt"
	"os"

	"github.com/AugusTaro/waicli/pkg/ai"
	"github.com/spf13/cobra"
)

// logGenCmd represents the 'log gen' command
var logGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "AIによる日報を生成して保存します",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🧠 日報を生成中...")

		// AI連携処理の呼び出し（仮の関数名）
		outFilePath, err := ai.GenerateNippou()
		if err != nil {
			return fmt.Errorf("日報生成に失敗しました: %v", err)
		}
		content, err := os.ReadFile(outFilePath)
		if err != nil {
			return fmt.Errorf("日報ファイルの読み込みに失敗しました: %v", err)
		}
		fmt.Printf("✅ 日報生成が完了しました: %s\n", outFilePath)
		fmt.Println(string(content))
		return nil
	},
}

func init() {
	logCmd.AddCommand(logGenCmd)
}
