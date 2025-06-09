package logstore

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// GetLatestTextLogPath returns the path to the latest .txt log file in ~/.wai-cli/logs/text/
func GetLatestTextLogPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ホームディレクトリ取得失敗: %v", err)
	}

	logDir := filepath.Join(homeDir, ".wai-cli", "logs", "text")

	entries, err := os.ReadDir(logDir)
	if err != nil {
		return "", fmt.Errorf("ログディレクトリ読み込み失敗: %v", err)
	}

	var logFiles []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".txt" {
			logFiles = append(logFiles, entry)
		}
	}

	if len(logFiles) == 0 {
		return "", fmt.Errorf("ログファイルが見つかりませんでした")
	}

	// ファイル名から日付をパースしてソート（名前は YYYY-MM-DD.txt 前提）
	sort.Slice(logFiles, func(i, j int) bool {
		ti, err1 := time.Parse("2006-01-02.txt", logFiles[i].Name())
		tj, err2 := time.Parse("2006-01-02.txt", logFiles[j].Name())
		if err1 != nil || err2 != nil {
			// パースに失敗した場合、ファイル名の文字列比較
			return logFiles[i].Name() > logFiles[j].Name()
		}
		return ti.After(tj)
	})

	latest := logFiles[0]
	fullPath := filepath.Join(logDir, latest.Name())
	return fullPath, nil
}
