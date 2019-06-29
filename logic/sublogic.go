package logic

import (
	"fmt"

	"github.com/awaduharatk/go-batch-starter/common"
	errorh "github.com/awaduharatk/go-batch-starter/error"
	"github.com/awaduharatk/go-batch-starter/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Sublogicinterface interface
type Sublogicinterface interface {
	SelectData(string) ([]model.User, error)
	OutputUser(users []model.User) error
	CreatePanic() error
}

// Sublogic st
type sublogic struct {
	db *gorm.DB
}

// NewSublogic sublogic constractor
func NewSublogic(db *gorm.DB) Sublogicinterface {
	return &sublogic{
		db,
	}
}

// SelectData データを取得する
func (sub *sublogic) SelectData(args string) ([]model.User, error) {
	users := []model.User{}

	if args == "9999" {
		// エラーになるSQLを実行
		err := sub.db.Raw("aaaaaaSELECT * FROM user").Scan(&users).Error
		if err != nil {
			return nil, errors.WithStack(err)
			// return nil, err
		}
	} else {
		_, err := common.Transact(sub.db, func(tx *gorm.DB) (interface{}, error) {
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
func (sub *sublogic) OutputUser(users []model.User) error {
	fmt.Println(users)
	return nil
}

// CreatePanic パニック生成
func (sub *sublogic) CreatePanic() error {

	fmt.Println("createPanic")
	panic(errorh.NewExitError(
		errorh.ExitCodeError,
		"E001",
		nil,
	))
	// panic("panic!!!!")
}
