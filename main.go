package main

import (
//   "os"
  "net/http"
  "fmt"
  "github.com/PuerkitoBio/goquery"
//   "ScrapeFile/functions"
//   "strings"
//   "net/http"
  "io/ioutil"
  "net/url"
  "net/http/cookiejar"



)
const (
	baseURL = "http://epu.tailieu.vn"
)
var h1_url[] string
var links[] string 
var srcs [] string
var filepath[] string
var (
	hdnSchoolCode = "epu"
	username = "daonamanh1"
  password = "123456"

)
type App struct {
	Client *http.Client
}
type Image struct {
	Name string
}
func (app *App) login() {
	client := app.Client
	loginURL := "http://epu.tailieu.vn/tai-khoan/dang-nhap.html"
	data := url.Values{
		"hdnSchoolCode" : {hdnSchoolCode},
		"txtLoginUsername":        {username},
		"txtLoginPassword":     {password},
		
	}

	response, _ := client.PostForm(loginURL, data)
	
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)

	
}

func (app *App) getImage() []Image {
	imagesURL := "http://epu.tailieu.vn/download/document/MTEwNDQwMjUxNjIwOA==.NTE2MjA4.html"
	client := app.Client

	response, _ := client.Get(imagesURL)

	defer response.Body.Close()

	document, _ := goquery.NewDocumentFromReader(response.Body)
	

	var images []Image

	document.Find("form#frmDownload").Each(func(i int, s *goquery.Selection) {
		name, _:= s.Find("input#MemberID").Attr("value")
		image := Image{
			Name: name,
		}

		images = append(images, image)
	})

	return images
}


func main(){
	jar, _ := cookiejar.New(nil)
	app := App{
		Client: &http.Client{Jar: jar},
	}
	app.login()
	
	fmt.Println("app :" ,app)

	images := app.getImage()
	fmt.Println("Image : ", images)
	for index, image := range images {
		fmt.Printf("%d: %s\n", index+1, image.Name)
	}
}