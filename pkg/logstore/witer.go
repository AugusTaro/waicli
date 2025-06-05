package logstore

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// PrepareTodayLogFile は今日の日付のログファイルを準備して返す
func PrepareTodayLogFile() (*os.File, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(home, ".wai-cli", "logs", "text")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	filename := time.Now().Format("2006-01-02") + ".txt"
	path := filepath.Join(dir, filename)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("ファイルオープン失敗: %w", err)
	}

	return file, nil
}
