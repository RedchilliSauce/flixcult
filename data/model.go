package data

//Visiblity - determines the visibility of an entity
type Visiblity string

const (
	//Public - the endpoint and its properties are visible to everybody
	Public Visiblity = "public"

	//GroupPrivate - the endpoint and its properties are visible only to
	// owner's groups
	GroupPrivate Visiblity = "group_private"

	//Private - the endpoint and its properties are visible only to owner
	Private Visiblity = "private"
)

//User - struct representing the user of the service
type User struct {
	Name       string `json:"name" db:"user_name"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	Email      string `json:"email" db:"email"`
}

//Group - is a set of users, who can share permissions
type Group struct {
	GroupID     string    `json:"groupID" db:"group_id"`
	Name        string    `json:"groupName" db:"name"`
	Owner       string    `json:"groupOwner" db:"owner"`
	Description string    `json:"groupDesc" db:"description"`
	Visiblity   Visiblity `json:"visibility" db:"visibility"`
}

//Show -
type Show struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Year int    `json:"year" db:"year"`
}

//Movie -
type Movie struct {
	Show
}

//Series -
type Series struct {
	Show
}
