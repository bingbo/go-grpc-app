package dao

import "testing"

func TestUserDao_GetUsers(t *testing.T) {
	userDao := NewUserDao()
	users, err := userDao.GetUsers()
	t.Log(users, err)
}
