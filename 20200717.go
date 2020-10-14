package main

import (
	"fmt"
	"time"
)

// "bufio"

// "errors"

// "fmt"

// "os"

// "os/exec"

// "strings"

// "time"

//--------------------------1
// func main() {
// a := make([]int, 3, 6) //[0 0 0] //第二个参数和第三个参数分别是len和cap,如果没有定义cap，那么俩者则相等
// //容量不能比长度小，长度大小会默认赋予零值
// a=append(a,1)//上面定义了容量为3的，那么此时里面就会有3个零值的，此时再次进行添加数据就是在这个三个末尾追加数字，如果超出cap那么则会扩容，扩容是翻倍的扩容
//扩容达到1024长度的话，不再是双倍扩容而是1.25倍

// fmt.Println(len(a))
// fmt.Println(cap(a))
// fmt.Println(a)
// }

// 输出:
// 4
// 6
// [0 0 0 1]
//--------------------------2
// type ss struct{
// 	A int
// }

// func main() {
// 	fmt.Println( "start")
// 	var f ss
// 	f.A=1
// 	// a:=map[string]ss{"a":f} //不能通过这个赋值，他只是类型// 输出: cannot assign to struct field a["a"].A in map
// 	a:=map[string]*ss{"a":&f} //可以通过指针赋值
//     a["a"].A=2

// 	fmt.Println( a["a"].A)

// 	}
//--------------------------3
//这道题考的是闭包+for range 参数V的,他并不是每次循环都定义的，而是只有一个变量值，然后进行赋值操作
func main() {
	a := make([]*int, 0, 2) //这里存的指针只是陷阱
	b, c := 2, 3
	a = append(a, &b)
	a = append(a, &c)
	for _, v := range a { //这里for循环的变量，是只定义一次，然后
		time.Sleep(1 * time.Second) //如果睡眠放在这里，那么他会立马循环赋值，然后结果是3，3
		fmt.Println(&v)             //每次循环的地址是一样的，证明了是一个值，并不是每次都是重新定义k,v,每次循环都是赋值操作
		go func() {
			fmt.Println(v)  //他的值是不一样的，因为存放的是变量地址指针
			fmt.Println(*v) //对地址取值

		}()
		time.Sleep(1 * time.Second) //如果睡眠放在这里，那么他不会立马循环赋值，然后结果是2，3
	}
	time.Sleep(3 * time.Second)
}

// 输出:
// // 0xc000006028
// // 0xc000006028
// // 3
// // 3
