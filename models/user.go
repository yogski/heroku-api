package models

type User struct {
	Phone    string     `json:"phone"`
	Name     string  `json:"name"`
	Email   float32 `json:"email"`
	Subscription interface{} `json:"subscription,omitempty"`
}

func GetUsers() ([]*User, error) {
	users := []*User{}
	err = db.Find(&users).Error
    return users, err
}

func GetUser(phone string) (*User, error){
	user := &User{}
	err = db.Where("phone = ?", phone).First(&user).Error
	return user, err
}

func CreateUser(user *User) (*User, error){
	ph := user.Phone
	u := User{Phone: ph}
	db.Find(&u)
	if u == nil {
		err = db.Create(user).Error
		return user, err
	}

}

func UpdateUser(user *User) (*User, error){
	err = db.Save(user).Error
	return user, err
}

func DeleteUser(phone string) error {
	err = db.Where("phone = ?",phone).Delete(&User{}).Error
	return err
}

