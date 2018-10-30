package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*User 用户*/
type User struct {
	Username           string
	Password           string
	Email              string
	Phone              string
	SponsorMeeting     []string //发起的会议title
	ParticipateMeeting []string //参与的会议title
}

/*Time 时间*/
type Time struct {
	timeID int //时间段，上午两段下午两段，编号1，2，3，4
	day    int
	month  int
}

/*Meeting 会议*/
type Meeting struct {
	Title        string
	Sponsor      string   //发起者
	Paticipators []string //参与者用户名
	MeetingTime  []Time   //可以包括多个时间段
}

/*CurrentUserName 当前登陆用户*/
var CurrentUserName string

/*READUSERS  读取文件*/
func READUSERS() (user []User) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/Users.txt")
	checkerr(err)

	//json转变为对象
	var users []User
	json.Unmarshal(b, &users)

	log.Println("READUSER success")
	return users
}

/*WRITEUSER 写入文件*/
func WRITEUSER(users []User) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(users)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/Users.txt", b, 0777)
	checkerr(err)
	log.Println("WRITEUSER success")
}

/*READMEETINGS 读取会议信息*/
func READMEETINGS() (meetings []Meeting) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/Meetings.txt")
	checkerr(err)
	//json转变为对象
	json.Unmarshal(b, &meetings)
	log.Println("READMEETINGS success")
	return meetings
}

/*WRITEMEETINGS 写入会议信息*/
func WRITEMEETINGS(meetings []Meeting) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(meetings)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/Meetings.txt", b, 0777)
	checkerr(err)
	log.Println("WRITEMEETINGS success")
}

func checkerr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
