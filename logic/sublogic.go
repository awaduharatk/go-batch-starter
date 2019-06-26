package logic

import (
	"fmt"

	"github.com/awaduharatk/go-batch-starter/db"
	errorh "github.com/awaduharatk/go-batch-starter/error"
	"github.com/awaduharatk/go-batch-starter/model"
	"github.com/jinzhu/gorm"
)

// Sublogic interface
type Sublogic interface {
	SelectData(string) ([]model.User, error)
	OutputUser(users []model.User) error
	CreatePanic() error
}

// Sublogic st
type sublogicst struct {
	db *gorm.DB
}

// NewSublogic sublogic constractor
func NewSublogic(db *gorm.DB) Sublogic {
	return &sublogicst{
		db,
	}
}

// SelectData データを取得する
func (sub *sublogicst) SelectData(args string) ([]model.User, error) {
	users := []model.User{}

	if args == "9999" {
		// エラーになるSQLを実行
		err := sub.db.Raw("aaaaaaSELECT * FROM user").Scan(&users).Error
		if err != nil {
			fmt.Println("db errror")
			return nil, err
		}
	} else {
		_, err := db.Transact(sub.db, func(tx *gorm.DB) (interface{}, error) {
			err := sub.db.Raw("SELECT * FROM user").Scan(&users).Error
			if err != nil {
				return nil, err
			}
			return users, nil
		})
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

// OutputUser user情報をファイルに出力
func (sub *sublogicst) OutputUser(users []model.User) error {
	fmt.Println(users)
	return nil
}

// CreatePanic パニック生成
func (sub *sublogicst) CreatePanic() error {

	fmt.Println("createPanic")
	panic(errorh.NewExitError(
		errorh.ExitCodeError,
		"E001",
		nil,
	))
	// panic("panic!!!!")
}
