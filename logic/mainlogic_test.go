package logic

import (
	"fmt"
	"testing"

	"github.com/awaduharatk/go-batch-starter/model"
)

// Mock用の構造体を定義してSublogicinterfaceを満たすように実装を行う
type sublogicMock struct{}

func (sub *sublogicMock) SelectData(string) ([]model.User, error) {
	fmt.Println("SelectData Mock!!!")
	return nil, nil
}

func (sub *sublogicMock) OutputUser(users []model.User) error {
	fmt.Println("OutputUser Mock!!!")
	return nil
}

func (sub *sublogicMock) CreatePanic() error {
	fmt.Println("CreatePanic Mock!!!")
	return nil
}

// sublogicをMockにしてMainlogigを実行する
func TestMock(t *testing.T) {
	fmt.Println("running!! TestMock!!!!")

	// Mockを差し込んだMainlogicを生成
	logic := Mainlogic{&sublogicMock{}}

	// 呼び出し
	logic.Logic([]string{"1234"})

}
