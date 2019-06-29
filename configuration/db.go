package configuration

import (
	"time"

	"github.com/jinzhu/gorm"
)

// NewDB dbconnection
func NewDB() *gorm.DB {
	// prop := config.GetProperties()

	// parseTime=trueを指定しないとdatetime→time.Timeへの変更でエラーが発生する。
	// CONNECT := prop.User + ":" + prop.Pass + "@" + prop.Protocol + "/" + prop.Dbname + "?parseTime=true" + "&readTimeout=10s"
	// db, err = gorm.Open(prop.Dbms, CONNECT)

	CONNECT := "user" + ":" + "password" + "@" + "tcp(localhost:3306)" + "/" + "sampledb" + "?parseTime=true" + "&readTimeout=10s"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
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
	return db
}
