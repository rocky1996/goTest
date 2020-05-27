package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}
type IPhone struct {

}

func (NokiaPhone NokiaPhone) call()  {
	fmt.Println("I am Nokia, I can call you!")
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone, I can call you!")
}

func main()  {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}