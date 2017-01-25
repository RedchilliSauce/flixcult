package data

//GetAllUsers -
func (pg *PostgresStore) GetAllUsers() (users []*User, err error) {
	return users, err
}

//GetUser -
func (pg *PostgresStore) GetUser(userName string) (user *User, err error) {
	return user, err
}

//GetUserWithEmail -
func (pg *PostgresStore) GetUserWithEmail(email string) (user *User, err error) {
	return user, err
}

//UserExists -
func (pg *PostgresStore) UserExists(userName string) (exists bool, err error) {
	return exists, err
}

//UserExistsWithEmail -
func (pg *PostgresStore) UserExistsWithEmail(email string) (exists bool, err error) {
	return exists, err
}

//CreateUser -
func (pg *PostgresStore) CreateUser(user *User) (err error) {
	return err
}

//UpdateUser -
func (pg *PostgresStore) UpdateUser(user *User) (err error) {
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
