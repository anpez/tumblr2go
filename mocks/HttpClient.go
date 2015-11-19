package mocks

import "github.com/stretchr/testify/mock"

type HttpClient struct {
	mock.Mock
}

func (_m *HttpClient) Get(url string, out interface{}) error {
	ret := _m.Called(url, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(url, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
