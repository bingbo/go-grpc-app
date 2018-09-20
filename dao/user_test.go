package dao

import "testing"

func TestUserDao_GetUsers(t *testing.T) {
	userDao := NewUserDao()
	users, err := userDao.GetUsers()
	t.Log(users, err)
}

func TestUserDao_GetUsersV2(t *testing.T) {
	userDao := NewUserDao()
	users, err := userDao.GetUsersV2()
	t.Log(users, err)
}

func TestUserDao_GetUsersV3(t *testing.T) {
	userDao := NewUserDao()
	users, err := userDao.GetUsersV3()
	t.Log(users, err)
}
