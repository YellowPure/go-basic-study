package main

import (
	// "github.com/veggiedefender/torrent-client/torrentfile"
	// "log"
	// "os"
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "strings"
	// "time"
)

// type animal interface {
// 	description() string
// }
// type Image struct {
// }

// type person struct {
// 	name   string
// 	age    int
// 	gender string
// 	Type   string
// }

// func (c person) description() string {
// 	return fmt.p("Name: %s", c.name)
// }

// type IPAddr [4]byte

// // TODO: 给 IPAddr 添加一个 "String() string" 方法
// func (ip IPAddr) String() string {
// 	var s = strings.Replace(strings.Trim(fmt.Sprint(ip), "[]"), " ", ".", -1)
// 	return fmt.Sprintf("%v", s)
// }

type student struct {
	Name string
	Age  int8
}

func main() {
	// message := "hello world"
	// fmt.Printf(message)
	// number2 := []int{1, 2, 3, 4}
	// fmt.Println(number2[:1])

	// ap := &number2
	// fmt.Print(ap)

	// var p animal
	// p = person{name: "Bob", age: 42, gender: "Male"}
	// fmt.Println(p.description())

	// go c()
	// fmt.Println("I am 1")
	// time.Sleep(time.Second * 2)
	// m := Image{}
	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }
	// c := [4]byte{8, 8, 8, 8}
	// var after = strings.Replace(strings.Trim(fmt.Sprint(c), "[]"), " ", ".", -1)
	// fmt.Printf("%v", after)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Geektutu")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		role := c.Query("role")
		c.String(http.StatusOK, "Hello %s , is %s", name, role)
	})

	r.POST("/user", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		id := c.Query("id")

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"username": username,
			"password": password,
		})
	})

	r.GET("/dyma", func(c *gin.Context) {
		c.String(http.StatusOK, "hot reload")
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	V1Logger := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			t := time.Now()
			c.Set("time", t)
			c.Next()
			late := time.Since(t)
			log.Print(late)
		}
	}
	v1 := r.Group("/v1/api")
	v1.Use(V1Logger())
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "Lily", Age: 20}
	stu2 := &student{Name: "Hanmeimei", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Run(":9999")
}

func defaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path": c.Request.URL.Path,
	})
}

// func c() {
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("I am 2")
// }
