// service_mock_test.go
package main

import (
	"time"
)

type MockDb struct {
	// mock.Mock
	storage map[string]string
}

func NewMockDb() *MockDb {
	var g MockDb
	g.storage = make(map[string]string)
	return &g
}

func (m *MockDb) Exists(login string) (int64, error) {
	// _ = m.Called()
	_, is_in := m.storage[login]
	if is_in {
		return 1, nil
	}
	return 0, nil
}

func (m *MockDb) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	_, is_in := m.storage[key]
	if is_in {
		return false, nil
	}
	m.storage[key] = "1"
	return true, nil
}
