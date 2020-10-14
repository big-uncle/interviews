package main

import "fmt"

//--------------------------
// type User struct {
// 	name  string
// 	email string
// }

// func (u User) not() {
// 	fmt.Println("Email is ", u.name, u.email)
// }

// func (u User) changeEmail(name, email string) { //这个题考的是指针的用法,   go中的方法绑定就是主要看接受者接受的类型的
// 	//取决于这里的接收者是指针还是变量，，变量是不会变的，相当于cp
// 	u.name = name
// 	u.email = email

// }
// func main() {
// 	m := User{"罗盼", "lpj_1996@163.com"}
// 	m.changeEmail("静静", "11947969678@qq.com") //调用方法时，不管你是对象还是指针，他都会帮你转成接受者想要的类型的，这个是看接收者，
// 	m.not()
// }

//--------------------------
//还有一道面试题，就是数组赋值，被赋值的数组改变，赋值的数组也会变（扩容除外），首先明确一下数组和切片的区别：数组是带大小的不可扩容的，比如str1 := [3]int{1, 2, 3}，而	str1 := []int{1, 2, 3}这个就是切片是可以append的
//数组与切片一样，只是扩容问题，本质都是数组
// func main() {
// 	str1 := []int{1, 2, 3}
// 	str2 := str1[1:] //初始化切片str2,是切片str1的引用，   他俩不管改变谁，另外一个都会改变（只是索引对应部分），因为他俩用的是一个内存地址，（因为堆内存中数组a的地址给了b数组，他们在堆内存中是同一片空间。b数组的改变就是a数组的改变）
// 	str1[1] = 8
// 	fmt.Println(str1)
// 	fmt.Println(str2)
// 	fmt.Println("------------------") //扩容之前还是之前的引用，扩容之后就是互不影响的
// 	str1 = append(str1, 9, 10)
// 	str1[1] = 0
// 	fmt.Println(str1)
// 	fmt.Println(str2)
// }

//--------------------------1
//下面这个考的是接口和类型的转换，
type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() { //这里是这个结构体的指针实现了这个接口并不是这个结构体实现了，   go里面只是提供了，在调用时允许 a.add() 或者  &a.add()   而且在用对象实现，那么可以用返回对象的指针，用实现的接口来接受

}

//接口由type和data组成，如果俩者都为nil，那么==nil条件才成立，
// 你给他赋值实现者为nil，那么这仅仅代表这个接口的data为nil，而打印则是输出接口的data
func live() People {
	var stu *Student
	return stu
	// 输出： BBBBBB < nil >
	// 	true
	// return nil
	// 输出： AAAAAA <nil>
	// true
}
func main() {
	if live() == nil {
		fmt.Println("AAAAAA", live())

	} else {
		fmt.Println("BBBBBB", live())
	}
	var s *Student
	fmt.Println(s == nil)
}
