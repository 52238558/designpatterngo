package creational

import "fmt"

type Product struct {
	partA string
	partB string
	partC string
}
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
}
//这个地方简化了builder和director的关联，合并成一个
type MacBookBuilderDirector struct {
	macBook *Product
}


func (r *MacBookBuilderDirector)Build() *Product{
	r.macBook = &Product{}
	r.BuildPartA()
	r.BuildPartB()
	r.BuildPartC()
	return r.macBook
}

func (r *MacBookBuilderDirector)BuildPartA(){
	r.macBook.partA = "macBook.partA.welldone"
}

func (r *MacBookBuilderDirector)BuildPartB(){
	r.macBook.partB = "macBook.partB.welldone"
}

func (r *MacBookBuilderDirector)BuildPartC(){
	r.macBook.partC = "macBook.partC.welldone"
}

func (r *Product)Print(){
	fmt.Println(r.partA)
	fmt.Println(r.partB)
	fmt.Println(r.partC)
}
