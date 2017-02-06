package data

import (
	"database/sql"
	"fmt"
)

//checkExists - checks if a row is selected by the given query
func (pg *PostgresStore) checkExists(
	query string, args ...interface{}) (exists bool, err error) {
	query = fmt.Sprintf("SELECT exists (%s)", query)
	rows := pg.QueryRowx(query, args...)
	if rows.Err() == nil {
		exists = true
	} else if rows.Err() == sql.ErrNoRows {
		exists = false
	} else {
		exists = false
		err = rows.Err()
	}
	return exists, err
}

//GetAllUsers -
func (pg *PostgresStore) GetAllUsers() (users []*User, err error) {
	users = make([]*User, 0, 20)
	err = pg.Select(&users, "SELECT * FROM fc_user ORDER BY user_name")
	return users, err
}

//GetUser -
func (pg *PostgresStore) GetUser(userName string) (user *User, err error) {
	user = &User{}
	queryStr := `SELECT * FROM fc_user WHERE user_name = ?`
	err = pg.Get(user, queryStr, userName)
	return user, err
}

//GetUserWithEmail -
func (pg *PostgresStore) GetUserWithEmail(email string) (user *User, err error) {
	user = &User{}
	queryStr := `SELECT * FROM orek_user WHERE email = ?`
	err = pg.Get(user, queryStr, email)
	return user, err
}

//UserExists -
func (pg *PostgresStore) UserExists(userName string) (exists bool, err error) {
	query := `SELECT 1 FROM fc_user WHERE user_name = '?' LIMIT 1`
	exists, err = pg.checkExists(query, userName)
	return exists, err
}

//UserExistsWithEmail -
func (pg *PostgresStore) UserExistsWithEmail(email string) (exists bool, err error) {
	query := `SELECT 1 FROM fc_user WHERE email = ? LIMIT 1`
	exists, err = pg.checkExists(query, email)
	return exists, err
}

//CreateUser -
func (pg *PostgresStore) CreateUser(user *User) (err error) {
	queryStr := `INSERT INTO fc_user( 
		user_name,  
		first_name, 
		second_name,
		email      
	) VALUES (
		:user_name,
		:first_name,
		:second_name,
		:email
	)`
	_, err = pg.NamedExec(queryStr, user)
	return err
}

//UpdateUser -
func (pg *PostgresStore) UpdateUser(user *User) (err error) {
	queryStr := `UPDATE fc_user SET
		first_name = :first_name,
		second_name = :second_name,
		email = :email
		WHERE user_name = :user_name
	`
	_, err = pg.NamedExec(queryStr, user)
	return err
}

//DeleteUser -
func (pg *PostgresStore) DeleteUser(userName string) (err error) {
	return err
}

//GetAllGroups -
func (pg *PostgresStore) GetAllGroups() (groups []*Group, err error) {
	return groups, err
}

//GetGroup -
func (pg *PostgresStore) GetGroup(GroupID string) (group *Group, err error) {
	return group, err
}

//GroupExists -
func (pg *PostgresStore) GroupExists(GroupID string) (exists bool, err error) {
	return exists, err
}

//CreateGroup -
func (pg *PostgresStore) CreateGroup(Group *Group) (err error) {
	return err
}

//UpdateGroup -
func (pg *PostgresStore) UpdateGroup(Group *Group) (err error) {
	return err
}

//DeleteGroup -
func (pg *PostgresStore) DeleteGroup(GroupID string) (err error) {
	return err
}

//AddUserToGroup -
func (pg *PostgresStore) AddUserToGroup(userName, groupID string) (err error) {
	return err
}

//UserExistsInGroup -
func (pg *PostgresStore) UserExistsInGroup(
	userName, groupID string) (exists bool, err error) {
	return exists, err
}

//RemoveUserFromGroup -
func (pg *PostgresStore) RemoveUserFromGroup(userName, groupID string) (err error) {
	return err
}

//GetUsersInGroup -
func (pg *PostgresStore) GetUsersInGroup(
	groupID string) (userInGroup []*User, err error) {
	return userInGroup, err
}

//GroupHasUser -
func (pg *PostgresStore) GroupHasUser(
	groupID, userName string) (has bool, err error) {
	return has, err
}

//GetGroupsForUser -
func (pg *PostgresStore) GetGroupsForUser(
	userName string) (groupsForUser []*Group, err error) {
	return groupsForUser, err
}
