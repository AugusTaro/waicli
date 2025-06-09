package ai

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIKey   string `yaml:"api_key"`
	Prompt   string `yaml:"prompt"`
	Model    string `yaml:"model"`
	Endpoint string `yaml:"endpoint"`
}

func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("ホームディレクトリ取得に失敗しました: %v", err)
	}
	configPath := filepath.Join(homeDir, ".wai-cli", "config.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("設定ファイルの解析に失敗しました: %v", err)
	}

	// 必須フィールドチェック
	if cfg.APIKey == "" || cfg.Prompt == "" || cfg.Model == "" || cfg.Endpoint == "" {
		return nil, fmt.Errorf("設定ファイルに必要な値が不足しています")
	}

	return &cfg, nil
}
