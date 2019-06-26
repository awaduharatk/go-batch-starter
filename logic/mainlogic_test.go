package logic

import "github.com/awaduharatk/go-batch-starter/model"

type sublogicMock struct{}

func (sub *sublogicMock) SelectData(string) ([]model.User, error) {}
func (sub *sublogicMock) OutputUser(users []model.User) error     {}
func (sub *sublogicMock) CreatePanic() error                      {}
