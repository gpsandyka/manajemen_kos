package services

import (
	"github.com/stretchr/testify/mock"
)

// type UserServiceMock struct {
// 	Mock mock.Mock
// }

type NotificationMock struct {
	Mock mock.Mock
}

// type ProblemMock struct {
// 	Mock mock.Mock
// }

// func (u *UserServiceMock) ProblemExtractorFromNotification(notif *NotificationMock) (*ProblemMock, error) {
// 	arguments := u.Mock.Called(notif)
// 	if arguments.Get(0) == nil {
// 		return nil, arguments.Error(1)
// 	}
// 	return arguments.Get(0).(*ProblemMock), arguments.Error(1)
// }
