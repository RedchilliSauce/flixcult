package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/RedchilliSauce/flixcult/data"
	"github.com/RedchilliSauce/flixcult/middleware"
	"github.com/labstack/echo"
)

//ReqBodyLimit - Req Body Size limit in bytes
const ReqBodyLimit int64 = 1048576 //1024*1024

//Temporary only
const httpBadStatus int = http.StatusBadRequest

//GroupUser - groupID user name map
type GroupUser struct {
	Name    string `json:"name"`
	GroupID string `json:"groupID"`
}

//LoginREST - function to login. Expects POST request and returns JWT token
func LoginREST(c echo.Context) error {
	var credential middleware.Credential

	body, err := ioutil.ReadAll(io.LimitReader(c.Request().Body, ReqBodyLimit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad request")
	}
	if err = json.Unmarshal(body, credential); err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect JSON format")
	}

	isValid, isAdmin := middleware.BasicAuthenticate(credential.Name, credential.Password)
	if !isValid {
		return c.JSON(http.StatusUnauthorized, "Username-password combination invalid. Try again.")
	}

	tokenString := middleware.GenerateJWT(middleware.ClaimsInp{IsAdmin: isAdmin, Name: credential.Name})
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

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

//GetUserByFilter - Get user based on attributes apart from name
//Valid inputs - email
func GetUserByFilter(c echo.Context) error {

	db := data.GetStore()
	params := c.QueryParams()

	if len(params) == 0 {
		return c.JSON(http.StatusBadRequest, "No filter provided")
	}

	email := params.Get("email")
	if email != "" {
		user, err := db.GetUserWithEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
			//TODO some common error function? Or just propogate?
		}
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusBadRequest, "Empty filter value for email")
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
		return c.JSON(http.StatusBadRequest, "Wrong JSON format")
	}
	if err = db.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, user)
}

//UpdateUser - Update User
func UpdateUser(c echo.Context) error {
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
	if err = db.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, user)
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
	if err = db.UpdateGroup(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, group)
}

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
	groupID := c.Param("groupID")

	exists, err := db.GroupExists(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		c.JSON(http.StatusBadRequest, "GroupID doesn't exist")
	}

	body, err := ioutil.ReadAll(io.LimitReader(c.Request().Body, ReqBodyLimit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
		//TODO some common error function? Or just propogate?
	}

	var groupUser GroupUser
	if err = json.Unmarshal(body, groupUser); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	uName := groupUser.Name
	if exists, err = db.UserExists(uName); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		return c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	if err = db.AddUserToGroup(uName, groupID); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}

	return c.JSON(http.StatusOK, groupUser)
}

//RemoveUserFromGroup - Delete User from Group based on username, groupID
func RemoveUserFromGroup(c echo.Context) error {
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

	body, err := ioutil.ReadAll(io.LimitReader(c.Request().Body, ReqBodyLimit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
		//TODO some common error function? Or just propogate?
	}

	var groupUser GroupUser
	if err = json.Unmarshal(body, groupUser); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	uName := groupUser.Name
	if exists, err = db.UserExists(uName); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}
	if exists == false {
		return c.JSON(http.StatusBadRequest, "Username doesn't exist")
	}

	if err = db.RemoveUserFromGroup(uName, groupID); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
		//TODO some common error function? Or just propogate?
	}

	return c.JSON(http.StatusOK, groupUser)
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

//GroupHasUser - Check if user exists in a group. User can be searched by name or email
func GroupHasUser(c echo.Context) error {
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

	params := c.QueryParams()

	var uName, email string
	exists = false
	for key, val := range params {
		if key == "name" {
			exists = true
			uName = val[0] //echo supports one value per param
		} else if key == "email" {
			exists = true
			email = val[0]
		}
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, "No valid search filters provided")
	}

	hasUser := false
	if uName != "" {
		exists, err = db.UserExists(uName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
			//TODO some common error function? Or just propogate?
		}
		if exists == false {
			c.JSON(http.StatusBadRequest, "Username doesn't exist")
		}

		hasUser, err = db.UserExistsInGroup(uName, groupID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
			//TODO some common error function? Or just propogate?
		}
	} else if email != "" {
		user, err := db.GetUserWithEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
			//TODO some common error function? Or just propogate?
		}
		if user == nil || user.Name == "" {
			c.JSON(http.StatusBadRequest, "User doesn't exist")
		}

		uName = user.Name
		hasUser, err = db.UserExistsInGroup(uName, groupID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
			//TODO some common error function? Or just propogate?
		}
	}

	userExistsInGroup := "false"
	if hasUser {
		userExistsInGroup = "true"
	}
	return c.JSON(http.StatusOK,
		map[string]string{"groupID": groupID, "name": uName, "userExistsInGroup": userExistsInGroup})
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
