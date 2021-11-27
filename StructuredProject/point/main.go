package main

import "fmt"

func add(a int){
	a += 1
}

func add2(a *int)  {
	*a += 1
}

func main()  {
	//与 C语言一致  &取地址 *取指针的值
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(p)  //存放x的地址
	fmt.Println(&p) //存放p所在的地址     禁止套娃！！！
	fmt.Println(*p) // "1"
	*p = 2 // equivalent to x = 2  通过 *可以取出地址指向的值
	fmt.Println(x) // "2"

	//	x  value=1 地址 = 0xc0000b4008
	// p=&x  value=0xc0000b4008  p也有一个属于自己的地址 0xc0000ae018

	//指针作为参数调用函数 可以通过指针修改变量的值
	y := 0
	add(y)
	fmt.Println("y = ",y)
	add2(&y)
	fmt.Println("指针修改y = ",y)


	//new 函数创建变量  返回的是地址
	//每次调用new函数都是返回一个新的变量的地址
	f := new(int)
	g := new(int)
	fmt.Println(f)
	fmt.Println(g)
}
