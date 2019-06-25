package logic

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/temp-go-dev/sample-batch/db"
	errorh "github.com/temp-go-dev/sample-batch/error"
	"github.com/temp-go-dev/sample-batch/model"
)

// Mainlogicinterface aa
type Mainlogicinterface interface {
	Logic([]string) error
}

// Mainlogic st
type Mainlogic struct {
}

// Logic 業務ロジックを記載
func (main *Mainlogic) Logic(args []string) error {
	if len(args) == 0 {
		return errorh.NewExitError(
			errorh.ExitCodeWarn,
			"0000",
			errors.New("引数0件エラー"),
		)
	}

	sub := &SublogicRegister{db: db.GetDB()}
	var err error

	var users []model.User
	// SelectDataを呼び出す
	users, err = sub.SelectData(args[0])
	if err != nil { // エラーが発生している場合,、ExitErrorを返却
		fmt.Println("error catch")
		return errorh.NewExitError(
			errorh.ExitCodeError,
			"E001",
			err,
		)
	}

	err = sub.OutputUser(users)
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
