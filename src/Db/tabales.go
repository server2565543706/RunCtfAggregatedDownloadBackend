package Db

import "time"

type InformationTable struct {
	Id      int64
	Name    string
	Connect string    `xorm:"varchar(200)"`
	Notes   string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
type AdministratorInformation struct {
	Id       int64
	Name     string
	Username string    `xorm:"varchar(200)"`
	Password string    `xorm:"varchar(200)"`
	Email    string    `xorm:"varchar(200)"`
	Token    string    `xorm:"varchar(200)"`
	Access   string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
