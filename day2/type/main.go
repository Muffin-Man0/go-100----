// 结构体

package main

import "time"
import "fmt"
type User struct {
	id 	int
	Score float32
	name,addr string
	enrollment time.Time
}

// 匿名结构体
// var user struct{Name string; Age int}


//  在User原型链加hello方法
func (user User) hello(man string) string {
	return "hello "+ man + ",i'm "+user.name	
}



func hello(man string,user User) string  {
	return "hello "+ man + ",i'm "+user.name	
}

func main()  {
	wx := User{Score:100,name:"zcy"}
	// wx.Score = 93.5
	// wx.enrollment = time.Now()
	
	// a:=wx.id+24
	// fmt.Printf("a=%d\n",a)

	fmt.Println(wx.hello("zcy"))
	fmt.Println(hello("zcy",wx))
}
