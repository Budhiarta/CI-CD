package mock

import "github.com/stretchr/testify/mock"

type LibMock struct {
	mock.Mock
}

func (l *LibMock) GenerateToken(id int) (string, error) {
	args := l.Called()

	return args.Get(0).(string), args.Error(1)
}
