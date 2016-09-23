package models

import (
    "github.com/astaxie/beego/orm"
	"fmt"
	"github.com/ajaybodhe/orm_test/conf"
	 _ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id          int        
    Name        string     
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Ayala       int
}

type Profile struct {
    Id          int
    Age         int16   
    User        *User   `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(User), new(Profile))
	
	orm.RegisterDriver("mysql", orm.DRMySQL)
	
	orm.RegisterDataBase("default", "mysql", conf.OrmNewConfig.DB.ConnID, 
		conf.OrmNewConfig.DB.MaxIdle, conf.OrmNewConfig.DB.MaxConn)
	
	// Database alias.
	name := "default"
	
	// Drop table and re-create.
	force := false
	
	// Print log.
	verbose := true
	
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
}

