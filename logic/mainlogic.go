package logic

import (
	"fmt"

	errorh "github.com/awaduharatk/go-batch-starter/error"
	"github.com/awaduharatk/go-batch-starter/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Mainlogicinterface aa
type Mainlogicinterface interface {
	Logic([]string) error
}

// Mainlogic st
type Mainlogic struct {
	sublogic Sublogicinterface
}

// NewMainlogic constractor
func NewMainlogic(db *gorm.DB) Mainlogicinterface {
	return &Mainlogic{
		NewSublogic(db),
	}
}

// Logic 業務ロジックを記載
func (main *Mainlogic) Logic(args []string) error {
	// 引数チェック
	if len(args) == 0 {
		return errorh.NewExitError(
			errorh.ExitCodeWarn,
			"0000",
			errors.New("引数0件エラー"),
		)
	}

	// sublogicのインスタンス生成
	// sub := NewSublogic(db.GetDB())
	var err error

	var users []model.User
	// SelectDataを呼び出す
	// users, err = main.sublogic.SelectData(args[0])
	users, err = main.sublogic.SelectData(args[0])
	if err != nil { // エラーが発生している場合,、ExitErrorを返却
		fmt.Println("error catch")
		return errorh.NewExitError(
			errorh.ExitCodeError,
			"E001",
			err,
		)
	}

	// err = main.sublogic.OutputUser(users)
	err = main.sublogic.OutputUser(users)
	if err != nil { // エラーが発生している場合,、ExitErrorを返却
		fmt.Println("error catch")
		return errorh.NewExitError(
			errorh.ExitCodeError,
			"E002",
			err,
		)
	}

	return nil
}
