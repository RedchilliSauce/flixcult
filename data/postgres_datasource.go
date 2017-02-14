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
	queryStr := `DELETE FROM orek_user WHERE user_name = ?`
	_, err = pg.Exec(queryStr, userName)
	return err
}

//GetAllGroups -
func (pg *PostgresStore) GetAllGroups() (groups []*Group, err error) {
	queryStr := `SELECT * FROM fc_group ORDER BY group_id`
	groups = make([]*Group, 0, 100)
	err = pg.Select(&groups, queryStr)
	return groups, err
}

//GetGroup -
func (pg *PostgresStore) GetGroup(groupID string) (group *Group, err error) {
	queryStr := `SELECT * FROM fc_group WHERE group_id = ?`
	group = &Group{}
	err = pg.Get(group, queryStr, groupID)
	return group, err
}

//GroupExists -
func (pg *PostgresStore) GroupExists(groupID string) (exists bool, err error) {
	query := `SELECT 1 FROM fc_group WHERE group_id = ? LIMIT 1`
	exists, err = pg.checkExists(query, groupID)
	return exists, err
}

//CreateGroup -
func (pg *PostgresStore) CreateGroup(group *Group) (err error) {
	queryStr := `INSERT INTO fc_group(
		group_id,
		name,
		owner,
		description,
		visibility
	) VALUES (
		:group_id,
		:name,
		:owner,
		:description
		:visibility 
	)`
	_, err = pg.NamedExec(queryStr, group)
	return err
}

//UpdateGroup -
func (pg *PostgresStore) UpdateGroup(group *Group) (err error) {
	queryStr := `UPDATE fc_user_group SET
			name = :name,
			owner = :owner,
			description = :description,
			visibility = :visibility
		WHERE group_id = :group_id`
	_, err = pg.NamedExec(queryStr, group)
	return err
}

//DeleteGroup -
func (pg *PostgresStore) DeleteGroup(groupID string) (err error) {
	queryStr := `DELETE FROM fc_user_group WHERE group_id = ?`
	_, err = pg.Exec(queryStr, groupID)
	return err
}

//AddUserToGroup -
func (pg *PostgresStore) AddUserToGroup(userName, groupID string) (err error) {
	queryStr := `INSERT INTO fc_user_to_group( 
		group_id,
		user_name
	) VALUES (
		?,
		?
	)`
	_, err = pg.Exec(queryStr, groupID, userName)
	return err
}

//UserExistsInGroup -
func (pg *PostgresStore) UserExistsInGroup(
	userName, groupID string) (exists bool, err error) {
	query := `SELECT 1 FROM fc_user_to_group WHERE 
		user_name = ? AND
		group_id = ? LIMIT 1`
	exists, err = pg.checkExists(query, userName, groupID)
	return exists, err
}

//RemoveUserFromGroup -
func (pg *PostgresStore) RemoveUserFromGroup(userName, groupID string) (err error) {
	queryStr := `DELETE FROM fc_user_to_group 
		WHERE group_id = ? AND user_name = ?`
	_, err = pg.Exec(queryStr, groupID, userName)
	return err
}

//GetUsersInGroup -
func (pg *PostgresStore) GetUsersInGroup(
	groupID string) (users []*User, err error) {
	queryStr := `SELECT * FROM fc_user WHERE user_name IN(
		SELECT user_name FROM fc_user_to_group WHERE group_id = ?
	)`
	users = make([]*User, 0, 100)
	err = pg.Select(&users, queryStr, groupID)
	return users, err
}

//GetGroupsForUser -
func (pg *PostgresStore) GetGroupsForUser(
	userName string) (groups []*Group, err error) {
	queryStr := `SELECT * FROM fc_user_group WHERE group_id IN (
		SELECT group_id FROM fc_user_to_group WHERE user_name = ?
	)`
	groups = make([]*Group, 0, 100)
	err = pg.Select(&groups, queryStr, userName)
	return groups, err
}
