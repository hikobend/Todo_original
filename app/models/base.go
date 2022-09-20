package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	// データーベースとエラー。ドライバとデーターベース名
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	// コマンドの作成
	// 最後にテーブル名を渡す
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	// コマンドを呼び出し
	Db.Exec(cmdU)
}

// UUID作成
// 返り値をuuidobj
func createUUID() (uuidobj uuid.UUID) {
	// NewUUIDを使用
	uuidobj, _ = uuid.NewUUID()
	// returnで返す
	return uuidobj
}

// パスワード作成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	// returnで返す
	return cryptext
}
