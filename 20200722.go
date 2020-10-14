package main

import (
	"fmt" // alias3
	"time"
)

//切片和数组区别，切片传参和数组传参，分别会不会发生变化
//答案： 切片由于是数组的引用，所以切片会变，数组不会变，
// 如果传的切片参数在函数内部进行了扩容，那么在函数内部的参数在扩容后修改切片就不会影响函数外部的原始切片，但是在扩容之前还是可以影响到外部的
func sliceModify(slice []int) { //传参俩者的地址并不是一样的，只是在内部的底层指向的还是外部的引用，直至扩容才会改变引用，但是外部在扩容之前不受影响
	slice[0] = 88 //切片传参，没修改之前，他还是原来的引用切片，修改了他会变,
	// append扩容之后由于因为地址发生了变化，所以信息也就没变了，只是这个函数里的切片变量变了，
	// 因为是在这里=赋值新的变量的，外部的依然没有扩容，外部的只有在内部没有扩容之前才会发生变化受影响
	fmt.Printf("address of slice %p add of Arr %p \n", &slice) //0xc0000964a0
	fmt.Printf("睡眠疫苗 \n")
	time.Sleep(1 * time.Second) //这里就算加上睡眠，那么上方和下方的地址也是一样的，
	slice = append(slice, 6)    //在切片扩容之前可以影响到外部，在扩容之后不会影响外部
	// 其实就算不扩容方法里面的地址和外面的也不是一个，就相当于内部的参数是外部参数的一个引用，
	// 所以本质俩个都不是一个地址， 只是内部的地址底层引用的还是外部的地址
	slice[0] = 99
	fmt.Println("ssss:", slice)
	fmt.Printf("address of slice %p add of Arr %p \n", &slice) //0xc0000964a0
}

// func main() { //切片底层本质也是数组本身，所以扩容改变的指向数组的实际引用，并不是简单的地址改变
// 	slice := []int{12, 34, 55, 66, 43}
// 	fmt.Println(slice)
// 	fmt.Printf("address of slice %p add of Arr %p \n", &slice) //0xc000096440
// 	sliceModify(slice)
// 	fmt.Println(slice)
// 	slice = append(slice, 6)                                   //这里扩容和不扩容那么main中的都一样，因为slice扩容不是说slice地址会改变，而是他底层指向的数组的地址会改变，
// 	fmt.Printf("address of slice %p add of Arr %p \n", &slice) //0xc000096440
// }

// 请记住以下两条规则：

// 如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
// 如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。

//可以利用select条件执行完才走 default的特性嵌套执行select这样可以间接的实现优先级
func main() {
	a := make(chan bool, 100)
	b := make(chan bool, 100)
	c := make(chan bool, 100)
	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case <-a:
			print("< a")

		default:

			select {

			case <-b:
				print("< b")

			case <-c:
				print("< c")

			default:
				print("< default")
			}
		}
	}
}
