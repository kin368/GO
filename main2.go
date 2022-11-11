/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import(
    "html/template"
    "log"  //用來輸出程式目前狀態
    "net/http"  //運行網頁用
    "errors"    //驗證帳號密碼
)
    
type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "KinGo"
	data.Content = "KinGoTask"
	tmpl.Execute(w, data)
}
func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/index", test)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("KinGoHtml: ", err)
	}

//寫死正確帳密
var UserData map[string]string

func init() {
	UserData = map[string]string{
		"AAAA": "123456",
	}
}
//驗證帳號密碼
func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("密碼錯誤,該吃砒霜搂")
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("帳號錯誤,該吃銀杏搂")
	}
}


//登入畫面
func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("請輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("請輸入密碼名稱"),
		})
		return
	}
	//檢查帳密是否正確
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
}