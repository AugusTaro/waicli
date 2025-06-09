package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/AugusTaro/waicli/pkg/logstore"
)

// GenerateNippou reads latest log, sends it to OpenAI, and saves the response as Markdown
func GenerateNippou() (string, error) {
	// 設定読み込み
	cfg, err := LoadConfig()
	if err != nil {
		return "", err
	}

	// 最新ログファイル読み込み
	logPath, err := logstore.GetLatestTextLogPath()
	if err != nil {
		return "", err
	}
	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		return "", fmt.Errorf("ログファイルの読み込み失敗: %v", err)
	}

	// メッセージ構築
	messages := []map[string]string{
		{"role": "system", "content": cfg.Prompt},
		{"role": "user", "content": string(logBytes)},
	}

	bodyMap := map[string]interface{}{
		"model":    cfg.Model,
		"messages": messages,
	}

	jsonBody, _ := json.Marshal(bodyMap)

	req, _ := http.NewRequest("POST", cfg.Endpoint, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("APIリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("APIエラー (%d): %s", resp.StatusCode, string(body))
	}

	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("レスポンスパース失敗: %v", err)
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("レスポンスに出力が含まれていません")
	}

	// 出力ディレクトリとファイルパス
	home, _ := os.UserHomeDir()
	outDir := filepath.Join(home, ".wai-cli", "logs", "nippou")
	os.MkdirAll(outDir, 0755)

	outFile := filepath.Join(outDir, time.Now().Format("2006-01-02")+".md")
	if err := os.WriteFile(outFile, []byte(res.Choices[0].Message.Content), 0644); err != nil {
		return "", fmt.Errorf("ファイル書き込み失敗: %v", err)
	}
	return outFile, nil
}
