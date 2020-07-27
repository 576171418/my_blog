package models

import (
	"fmt"
	"time"
)

type User struct {
	Id			int64
	Name		string		`xorm:"name" json:"name"`
	Phone		string		`xorm:"phone" json:"phone"`
	CreatedAt 	time.Time	`xorm:"created_at" json:"created_at"`
	Password	string		`xorm:"password" json:"password"`
	Avatar		string		`xorm:"avatar" json:"avatar"`
}

func (user *User) Create() (err error)  {
	if err = Engine.Sync2(User{}); err != nil {
		return
	}

	_, err = Engine.Insert(&User{
		Id:	       0,
		Name:      user.Name,
		Phone:     user.Phone,
		CreatedAt: time.Time{},
		Password:  user.Password,
		Avatar:	   "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2519824424,1132423651&fm=26&gp=0.jpg",
	})
	fmt.Println(time.Time{})
	if err != nil {
		fmt.Println("")
		return
	}
	return
}

func (user *User) Delete() (err error) {
	if err = Engine.Sync2(User{}); err != nil {
		return
	}
	_, err = Engine.Delete(&user)

	if err != nil {
		return
	}
	return
}

func (user *User) Update() (err error) {
	if err = Engine.Sync2(User{}); err != nil {
		return
	}
	_, err = Engine.Update(&user)
	return
}

func Users() (users []User, err error) {
	if err = Engine.Ping(); err != nil {
		return
	}
	err = Engine.Find(&users)
	return
}

func UserById(id int) (err error, user User) {
	if err = Engine.Ping(); err != nil {
		return
	}
	_, err = Engine.Where("id = ?", id).Get(&user)
	return
}

func AuthUserPassword(phone, password string) (has bool, err error) {
	if err = Engine.Ping(); err != nil {
		return
	}
	has, err = Engine.Where("phone = ? and password = ?", phone, password).Get(&User{})
	return
}

func CheckPhoneExist(phone string) bool {
	if err := Engine.Ping(); err != nil {
		panic(fmt.Sprintf("ping database failed: %s", err.Error()))
	}
	if err := Engine.Sync2(User{}); err != nil {
		panic(fmt.Sprint(err))
	}
	has, err := Engine.Where("phone = ?", phone).Get(&User{})
	if err != nil {
		panic(fmt.Sprint("查询失败, ", err))
	}
	return has
}