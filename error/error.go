package error

import (
	"fmt"
)

const (
	//ExitCodeOK 正常終了
	ExitCodeOK = 0
	//ExitCodeWarn 警告終了
	ExitCodeWarn = 1
	// ExitCodeError 異常終了
	ExitCodeError = 2
)

// ExitError エラーインターフェースを実装した構造体
type ExitError struct {
	exitCode  int    // 終了コードを定義 const化したほうがいい
	messageID string // メッセージIDの定義
	err       error  //エラーをそのまま入れる
}

func (ee *ExitError) Error() string {
	if ee.err == nil {
		return ""
	}
	return fmt.Sprintf("%v", ee.err)
}

// NewExitError Exiterror型の生成
func NewExitError(exitCode int, messageID string, err error) *ExitError {
	return &ExitError{
		exitCode:  exitCode,
		messageID: messageID,
		err:       err,
	}
}

// HandleExit errorのハンドリングを行う
func HandleExit(err error) int {

	// Errorが設定されていない場合
	if err == nil {
		return ExitCodeOK
	}

	// ExitErrorの場合
	if e, ok := err.(*ExitError); ok {
		fmt.Println("errorHandle!! ExitErrorType")

		fmt.Printf("messageID:%v\n", e.messageID)
		fmt.Printf("err:%v\n", e.err)
		fmt.Printf("exitCode:%v\n", e.exitCode)
		return e.exitCode
	}

	// errorの場合
	if e, ok := err.(error); ok {
		fmt.Println("errorHandle!! errorType")
		fmt.Printf("error:%v\n", e)
		return ExitCodeError
	}



	return ExitCodeOK
}
