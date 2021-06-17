package infrastructure

import (
	"bufio"
	"encoding/json"
	"os"
)

func (c* Context)rewriteFileWithoutLine(filePath string, numberOfLine int){
	fileMutex.Lock()
	defer fileMutex.Unlock()

	tempFilePath := filePath + "temp"

	fileForRead, err:= os.OpenFile(c.FileName, os.O_RDONLY, os.ModePerm)
	if err!= nil{
		panic(err.Error())
	}

	fileForWrite, err := os.OpenFile(tempFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err!= nil{
		panic(err.Error())
	}

	scanner := bufio.NewScanner(fileForRead)
	iter := 0
	for scanner.Scan(){
		iter++
		if iter-1 == numberOfLine{
			continue
		}
		_, err = fileForWrite.Write(scanner.Bytes())
		if err != nil {
			panic(err.Error())
		}
		_, err = fileForWrite.WriteString("\n")
		if err != nil {
			panic(err.Error())
		}
	}
	err = fileForRead.Close()
	if err != nil {
		panic(err.Error())
	}

	err = fileForWrite.Close()
	if err != nil {
		panic(err.Error())
	}

	err = os.Remove(filePath)
	if err != nil {
		panic(err.Error())
	}
	err = os.Rename(tempFilePath, filePath)
	if err != nil {
		panic(err.Error())
	}
}


/*
Возвращает пользователя по id
(пользователь, его номер строки)
*/
func (c *Context) getUserById(id string) (*User, int){
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err:= os.OpenFile(c.FileName, os.O_RDONLY, os.ModePerm)
	if err!= nil{
		panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	iter := 0
	for scanner.Scan(){
		var user User
		err = json.Unmarshal(scanner.Bytes(), &user)
		if err!=nil{
			panic(err.Error())
		}
		if user.Id == id{
			return &user, iter
		}
		iter++
	}
	return nil, -1
}


func (c *Context) addUserToEndOfFile(user *User){
	fileMutex.Lock()
	defer fileMutex.Unlock()

	bytes, err := json.Marshal(&user)
	if err!=nil{
		panic(err.Error())
	}
	file, err := os.OpenFile(c.FileName, os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		panic(err.Error())
	}
	_, err = file.WriteString("\n")
	if err != nil {
		panic(err.Error())
	}
}

func (c *Context) readLines(skip, take int) []User{
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err:= os.OpenFile(c.FileName, os.O_RDONLY, os.ModePerm)
	if err!= nil{
		panic(err.Error())
	}
	defer file.Close()

	users:=make([]User, 0, take)

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		//var bytes = scanner.Bytes()
		skip--
		if skip >= 0{
			continue
		}
		if take <= 0{
			break
		}
		take--

		var user User
		err = json.Unmarshal(scanner.Bytes(), &user)
		if err!=nil{
			panic(err.Error())
		}
		users = append(users, user)
	}

	return users
}

func (c *Context) getCountOfRecords() int{
	fileMutex.Lock()
	defer fileMutex.Unlock()
	file, err:= os.OpenFile(c.FileName, os.O_RDONLY, os.ModePerm)
	if err!= nil{
		panic(err.Error())
	}
	defer file.Close()

	count:=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		count++
	}
	return count
}
