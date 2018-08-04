package main

import "sync"

var userList = &UserList{
	Data:make(map[int64]*User),
	rw:&sync.RWMutex{},
}

type UserList struct {
	Data map[int64]*User
	rw   *sync.RWMutex
}

func (this *UserList) Set(userid int64, user *User) {

	this.rw.Lock()
	defer this.rw.Unlock()

	this.Data[userid] = user
}

func (this *UserList) Get(userid int64) *User {

	this.rw.RLock()
	defer this.rw.RUnlock()

	user, ok := this.Data[userid]
	if !ok {
		return nil
	}

	return user
}

// 获取 user 所有数据
func (this *UserList) List() map[int64]*User {

	this.rw.RLock()
	defer this.rw.RUnlock()

	return this.Data
}
