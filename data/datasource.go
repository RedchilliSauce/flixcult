package data

import "log"

var db Store

//SetStore sets the application wIDe data store to be used
func SetStore(dbInst Store) error {
	db = dbInst
	return nil
}

//GetStore - give the application data store
func GetStore() Store {
	if db == nil {
		log.Fatal("DataStore is Nil!!!")
	}
	return db
}

//Store - interface declares the operation that will be exposed by a
//application data store
type Store interface {
	GetAllUsers() (users []*User, err error)
	GetUser(userName string) (user *User, err error)
	GetUserWithEmail(email string) (user *User, err error)
	UserExists(userName string) (exists bool, err error)
	UserExistsWithEmail(email string) (exists bool, err error)
	CreateUser(user *User) (err error)
	UpdateUser(user *User) (err error)
	DeleteUser(userName string) (err error)

	GetAllGroups() (groups []*Group, err error)
	GetGroup(GroupID string) (group *Group, err error)
	GroupExists(GroupID string) (exists bool, err error)
	CreateGroup(Group *Group) (err error)
	UpdateGroup(Group *Group) (err error)
	DeleteGroup(GroupID string) (err error)
	AddUserToGroup(userName, groupID string) (err error)
	UserExistsInGroup(userName, groupID string) (exists bool, err error)
	RemoveUserFromGroup(userName, groupID string) (err error)
	GetUsersInGroup(groupID string) (userInGroup []*User, err error)
	GroupHasUser(groupID, userName string) (has bool, err error)
	GetGroupsForUser(userName string) (groupsForUser []*Group, err error)
}
