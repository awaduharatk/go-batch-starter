package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Init DB初期化
func Init() {
	// prop := config.GetProperties()

	// parseTime=trueを指定しないとdatetime→time.Timeへの変更でエラーが発生する。
	// CONNECT := prop.User + ":" + prop.Pass + "@" + prop.Protocol + "/" + prop.Dbname + "?parseTime=true" + "&readTimeout=10s"
	// db, err = gorm.Open(prop.Dbms, CONNECT)

	CONNECT := "user" + ":" + "password" + "@" + "tcp(localhost:3306)" + "/" + "sampledb" + "?parseTime=true" + "&readTimeout=10s"
	db, err = gorm.Open("mysql", CONNECT)

	if err != nil {
		//　err発生時の処理要検討
		panic(err.Error())
	}
	// DBデバッグログの出力設定
	db.LogMode(true)
	// db.SetLogger(config.GetLogger())

	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetimeは再利用され得る最長時間を設定
	db.DB().SetConnMaxLifetime(time.Hour)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Transact トランザクション実行
func Transact(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	// 無名関数にBeginしたDBを渡して実行する
	data, err = txFunc(tx)
	return data, err
}
