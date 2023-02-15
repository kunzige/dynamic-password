package models

import "time"

type User struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Identity  string    `xorm:"VARCHAR(20)"`
	Email     string    `xorm:"VARCHAR(20)"`
	Password  string    `xorm:"VARCHAR(40)"`
	Nickname  string    `xorm:"VARCHAR(30)"`
	Status    int       `xorm:"default 0 comment('0 正常 1被禁用') INT"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
