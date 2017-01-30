package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/RedchilliSauce/flixcult/data"
	"github.com/labstack/echo"
)

//ReqBodyLimit - Req Body Size limit in bytes
const ReqBodyLimit int64 = 1048576 //1024*1024

//UsernameExists - Struct containing info on whether username exists or not
type UsernameExists struct {
	UserName string `json:"name"`
	Exists   bool   `json:"exists"`
}

//GroupIDExists - Struct containing info on whether group exists or not
type GroupIDExists struct {
	GroupID string `json:"groupID"`
	Exists  bool   `json:"exists"`
}

//EmailExists - Struct containing info on whether email(of a user) exists or not
type EmailExists struct {
	Email  string `json:"email"`
	Exists bool   `json:"exists"`
}

//Temporary only
const httpBadStatus int = http.StatusBadRequest

//GetAllUsers - Get list of users- Implemented with Echo JSON Streaming to support large JSON objects
func GetAllUsers(c echo.Context) error {
	db := data.GetStore()
	users, err := db.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(users)
}

//GetUser - Get user based on username
func GetUser(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")

	user, err := db.GetUser(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, user)
}

//GetUserWithEmail - Get user based on email
func GetUserWithEmail(c echo.Context) error {
	db := data.GetStore()
	email := c.Param("email")

	user, err := db.GetUser(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, user)
}

//UserExists - Check if user exists based on username
func UserExists(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")

	exists, err := db.UserExists(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, UsernameExists{UserName: uName, Exists: exists})
}

//UserExistsWithEmail - Check if user exists based on email
func UserExistsWithEmail(c echo.Context) error {
	db := data.GetStore()
	email := c.Param("email")

	exists, err := db.UserExistsWithEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, EmailExists{Email: email, Exists: exists})
}

//CreateUser - Create User
func CreateUser(c echo.Context) error {
	db := data.GetStore()
	var user data.User

	body, err := ioutil.ReadAll(io.LimitReader(c.Request().Body, ReqBodyLimit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
		//TODO some common error function? Or just propogate?
	}
	if err = json.Unmarshal(body, user); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	if err = db.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, user)
}

//UpdateUser - Update User
func UpdateUser(c echo.Context) error {
	//TODO
	return nil
}

//TODO : Shall we return the deleted User details?

//DeleteUser - Delete user based on username
func DeleteUser(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")

	user, err := db.GetUser(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	err = db.DeleteUser(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	//TODO What to return here?
	return c.JSON(http.StatusOK, user)
}

//GetAllGroups - Get all groups
func GetAllGroups(c echo.Context) error {
	db := data.GetStore()

	groups, err := db.GetAllGroups()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(groups)
}

//GetGroup - Get group based on groupID
func GetGroup(c echo.Context) error {
	db := data.GetStore()
	groupID := c.Param("groupID")

	group, err := db.GetUser(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, group)
}

//GroupExists - Check if group exists based on groupID
func GroupExists(c echo.Context) error {
	db := data.GetStore()
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	return c.JSON(http.StatusOK, GroupIDExists{GroupID: groupID, Exists: exists})
}

//CreateGroup - Create Group
func CreateGroup(c echo.Context) error {
	db := data.GetStore()
	var group data.Group

	body, err := ioutil.ReadAll(io.LimitReader(c.Request().Body, ReqBodyLimit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
		//TODO some common error function? Or just propogate?
	}
	if err = json.Unmarshal(body, group); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	if err = db.CreateGroup(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, group)
}

//UpdateGroup - Update Group
func UpdateGroup(c echo.Context) error {
	//TODO
	return nil
}

//TODO : Shall we return the deleted Group details?

//DeleteGroup - Delete Group based on groupID
func DeleteGroup(c echo.Context) error {
	db := data.GetStore()
	groupID := c.Param("groupID")

	err := db.DeleteGroup(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	//TODO What to return here?
	return c.JSON(http.StatusOK, "WOOHOO")
}

//AddUserToGroup - Add a user to group
func AddUserToGroup(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "GroupID doesn't exist")
	}

	if exists, err = db.UserExists(uName); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	if err = db.AddUserToGroup(uName, groupID); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}

	//TODO What to return here?
	return c.JSON(http.StatusOK, "WOOHOO")
}

//RemoveUserFromGroup - Delete User from Group based on username, groupID
func RemoveUserFromGroup(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "GroupID doesn't exist")
	}

	if exists, err = db.UserExists(uName); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	if err = db.RemoveUserFromGroup(uName, groupID); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}

	//TODO What to return here?
	return c.JSON(http.StatusOK, "WOOHOO")
}

//GetUsersInGroup - Get list of users in a group
func GetUsersInGroup(c echo.Context) error {
	db := data.GetStore()
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "GroupID doesn't exist")
	}

	users, err := db.GetUsersInGroup(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(users)
}

//GroupHasUser - Check if user exists in a group
func GroupHasUser(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "GroupID doesn't exist")
	}

	exists, err = db.UserExists(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	hasUser, err := db.GroupHasUser(groupID, uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}

	//TODO What to return here?
	return c.JSON(http.StatusOK, hasUser)
}

//GetGroupsForUser - Get list of groups that a user is a part of
func GetGroupsForUser(c echo.Context) error {
	db := data.GetStore()
	uName := c.Param("name")

	exists, err := db.UserExists(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	groups, err := db.GetGroupsForUser(uName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(groups)
}
