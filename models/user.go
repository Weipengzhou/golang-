package models

import (
	"errors"
	"strconv"
	"time"
		"io/ioutil"
  "net/http"	
  "fmt"


)

var (
	UserList map[string]*User

)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}
func GetInformation(server_id string , areaid string,server_name string,page string,query_order string,kindid string,view_loc string,count string)(string){
	fmt.Printf("https://recommd.xyq.cbg.163.com/cgi-bin/recommend.py?act=recommd_by_role&server_id="+server_id+"&areaid="+areaid+"&server_name="+server_name+"&page="+page+"&query_order="+query_order+"&kindid="+kindid+"&view_loc="+view_loc+"&count="+count)
	resp, err := http.Get("https://recommd.xyq.cbg.163.com/cgi-bin/recommend.py?act=recommd_by_role&server_id="+server_id+"&areaid="+areaid+"&server_name="+server_name+"&page="+page+"&query_order="+query_order+"&kindid="+kindid+"&view_loc="+view_loc+"&count="+count)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		return err.Error()
	 }
	return string(body)
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
