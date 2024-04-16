package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type operation func()

func ExecutNtimes(n int, op operation) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			op()
		}()
	}
	wg.Wait()
}
func GenerateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)

	for i := range s {
		randindex := rand.Intn(len(letters))
		s[i] = letters[randindex]
	}
	return string(s)
}

func TestUrlr(url string, body interface{}) {
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)
	fmt.Printf("请求内容: %+v, 响应时间: %v, 状态码: %d\n", body, duration, resp.StatusCode)
}

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	NewPassword     string `json:"newpassword"`
}

func TestLogin() {
	url := "http://localhost:8080/user/login"
	user := User{
		Username: readUsernameFromFile(),
		Password: "123456",
	}
	TestUrlr(url, user)
}

func TestRegister() {
	filename := "users.txt"
	psw := "123456"
	user := User{
		Username:        GenerateRandomString(10),
		Password:        psw,
		ConfirmPassword: psw,
	}
	saveUserToFile(filename, user.Username)

	url := "http://localhost:8080/user/register"
	TestUrlr(url, user)
}

func TestMpwd() {

	username := readUsernameFromFile()
	password := "123456"
	newPwd := "654321"
	user := User{
		Username:    username,
		Password:    password,
		NewPassword: newPwd,
	}

	url := "http://localhost:8080/user/modifypassword"
	TestUrlr(url, user)
}

func TestUeserlist() {
	url := "http://localhost:8080/user/list"
	TestUrlr(url, nil)
}

type Page struct {
	MyPage int `json:"counts"`
	Offset int `json:"offset"`
}

func TestUeserpagelist() {
	url := "http://localhost:8080/user/pagelist"
	page := Page{
		MyPage: rand.Intn(10),
		Offset: rand.Intn(10),
	}
	TestUrlr(url, page)
}

func TestUeserdelete() {
	username := readUsernameFromFile()
	user := User{
		Username: username,
	}
	url := "http://localhost:8080/user/delete"
	TestUrlr(url, user)
}

type Moment struct {
	UserID     uint64 `json:"user_id"`
	Content    string `json:"content"`
	ImagePaths string `json:"image_paths"`
}
type MomentsByUserID struct {
	UserID   int `json:"user_id"`
	MomentId int `json:"moment_id"`
}

func TestSentMoment() {
	url := "http://localhost:8080/moment/send"
	moment := Moment{
		UserID:     uint64(rand.Intn(50)),
		Content:    GenerateRandomString(10),
		ImagePaths: GenerateRandomString(10),
	}
	TestUrlr(url, moment)
}
func TestMomentlist() {
	url := "http://localhost:8080/moment/list"
	TestUrlr(url, nil)
}
func TestMomentpagelist() {
	url := "http://localhost:8080/moment/pagelist"
	page := Page{
		MyPage: rand.Intn(10),
		Offset: rand.Intn(10),
	}
	TestUrlr(url, page)
}
func Testusermoments() {
	url := "http://localhost:8080/moment/usermoments"
	page := Page{
		MyPage: rand.Intn(10),
		Offset: rand.Intn(10),
	}
	TestUrlr(url, page)
}
func Testmomentdelete() {
	url := "http://localhost:8080/moment/delete"
	Moment := MomentsByUserID{
		MomentId: rand.Intn(10),
	}
	TestUrlr(url, Moment)
}
func Testmomentlike() {
	url := "http://localhost:8080/moment/like"
	Moment := MomentsByUserID{
		UserID:   rand.Intn(500),
		MomentId: 240 + rand.Intn(20),
	}
	TestUrlr(url, Moment)
}

func Testmomentunlike() {
	url := "http://localhost:8080/moment/unlike"
	Moment := MomentsByUserID{
		UserID:   rand.Intn(500),
		MomentId: rand.Intn(100),
	}
	TestUrlr(url, Moment)
}

func Testmomentcountlikes() {
	url := "http://localhost:8080/moment/countlikes"
	Moment := MomentsByUserID{
		UserID:   rand.Intn(100),
		MomentId: rand.Intn(100),
	}
	TestUrlr(url, Moment)
}

type JsonUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func saveUserToFile(filename, username string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", username))
	if err != nil {
		return err
	}
	return nil
}

func readUsernameFromFile() string {
	content, _ := ioutil.ReadFile("users.txt")
	lines := strings.Split(string(content), "\n")
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(lines))
	randomLine := lines[randomIndex]
	return randomLine
}

func main() {
	// ExecutNtimes(400, TestRegister)
	// ExecutNtimes(20, TestSentMoment)
	ExecutNtimes(50, Testmomentlike)
}
