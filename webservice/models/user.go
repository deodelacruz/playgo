package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID ! = 0 {
		return User{}, errors.New("New User must not include ID.")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int)  (User,error) {
	for _,u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{},fmt.Errorf("user with id '%v' not found.",u.ID)
}

func RemoveUserBydID(id int) error {
	for i,i:= range users {
		if u.ID == id {
			user = append(users[:i],users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found",id)
}
