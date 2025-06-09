package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初期設定ファイル (.wai-cli/config.yaml) を生成します",
	RunE: func(cmd *cobra.Command, args []string) error {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("ホームディレクトリの取得に失敗しました: %v", err)
		}

		configDir := filepath.Join(homeDir, ".wai-cli")
		configPath := filepath.Join(configDir, "config.yaml")

		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("⚠️ 設定ファイルはすでに存在します: %s\n", configPath)
			return nil
		}

		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("ディレクトリ作成失敗: %v", err)
		}

		template := `api_key: sk-xxxxx  # OpenAIのAPIキーをここに入力
model: gpt-4
endpoint: https://api.openai.com/v1/chat/completions
prompt: |
  以下は作業ログです。これをもとにMarkdown形式で日報を作成してください。
`

		if err := os.WriteFile(configPath, []byte(template), 0644); err != nil {
			return fmt.Errorf("設定ファイルの作成に失敗しました: %v", err)
		}

		fmt.Printf("✅ 設定ファイルを作成しました: %s\n", configPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
