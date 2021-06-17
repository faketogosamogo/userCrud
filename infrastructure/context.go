package infrastructure

import (
	"os"
	"sync"
)

/*
А вот благодаря этому вылетать не должно
 */
var fileMutex = &sync.Mutex{}
/*
Как понял, в одном экземляре будет только нормально работаь
 */



type Context struct{
	FileName string
}

func ContextNew() *Context{
	filename:= "db.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
	return &Context{FileName: filename}
}

func (c *Context) AddUser(user* User){
	c.addUserToEndOfFile(user)
}
func (c *Context) GetUser(id string) *User{
	user, _ := c.getUserById(id)
	return user
}

func (c *Context) EditUser(user *User) *User{
	_, numberOfLine := c.getUserById(user.Id)
	if numberOfLine == -1{
		return nil
	}
	c.rewriteFileWithoutLine(c.FileName, numberOfLine)
	c.addUserToEndOfFile(user)
	return user
}

func (c *Context) RemoveUser(user *User) *User{
	_, numberOfLine := c.getUserById(user.Id)
	if numberOfLine == -1{
		return nil
	}
	c.rewriteFileWithoutLine(c.FileName, numberOfLine)
	return user
}

func (c *Context) IndexUser(perPage, page int) []User{
	return c.readLines(perPage*(page-1), perPage)
}

func (c *Context) CountOfRecords()int{
	return c.getCountOfRecords()
}