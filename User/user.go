package user

import "time"

type User struct {
	Id          int
	Username    string
	Email       string
	Password    string
	Create_time time.Time
}
