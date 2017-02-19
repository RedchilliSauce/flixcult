package service

import "github.com/labstack/echo"

//TODO : Versioning
//Route - Attributes of a route to be registered with router
type Route struct {
	Methods         []string
	Path            string
	HandlerFunc     echo.HandlerFunc
	MiddlewareFuncs []echo.MiddlewareFunc
}

//Routes - Stores all routes(each of type Route)
var Routes = []Route{
	Route{
		Methods:         []string{"POST"},
		Path:            "/login",
		HandlerFunc:     LoginREST,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "/users",
		HandlerFunc:     GetAllUsers,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "/users/:name",
		HandlerFunc:     GetUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "/users/filtersearch",
		HandlerFunc:     GetUserByFilter,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"POST"},
		Path:            "/users/:name",
		HandlerFunc:     CreateUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"PUT"},
		Path:            "/users/:name",
		HandlerFunc:     UpdateUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"DELETE"},
		Path:            "/users/:name",
		HandlerFunc:     DeleteUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "/groups",
		HandlerFunc:     GetAllGroups,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "/groups/:groupID",
		HandlerFunc:     GetGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"POST"},
		Path:            "/groups/:groupID",
		HandlerFunc:     CreateGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"PUT"},
		Path:            "/groups/:groupID",
		HandlerFunc:     UpdateGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},

	Route{
		Methods:         []string{"DELETE"},
		Path:            "/groups/:groupID",
		HandlerFunc:     DeleteGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"PUT"},
		Path:            "/groups/:groupID/users/add",
		HandlerFunc:     AddUserToGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"PUT"},
		Path:            "/groups/:groupID/users/remove",
		HandlerFunc:     RemoveUserFromGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "groups/:groupID/users",
		HandlerFunc:     GetUsersInGroup,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "groups/:groupID/users/filtersearch",
		HandlerFunc:     GroupHasUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
	Route{
		Methods:         []string{"GET"},
		Path:            "users/:name/groups",
		HandlerFunc:     GetGroupsForUser,
		MiddlewareFuncs: DefaultMWFuncs,
	},
}
