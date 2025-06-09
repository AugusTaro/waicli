# wai-cli

🎯 **wai-cli** は、日々の作業ログを手軽に記録し、AIを活用して自動で日報を生成する CLI ツールです。  
テキストメモ → Markdown日報 までをシンプルなコマンド操作で完結できます。

---

## 🚀 機能一覧

- `log write`：作業メモを記録（リアルタイム追記）
- `log gen`：AIで日報をMarkdown形式で自動生成
- `init`：設定ファイル（`~/.wai-cli/config.yaml`）のテンプレートを作成

---

## 🛠 インストール

```bash
git clone https://github.com/yourname/wai-cli.git
cd wai-cli
go build -o wai-cli
