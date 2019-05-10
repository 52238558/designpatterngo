package creational

import "fmt"
//鉴于对go理解，利用了go值的性质实现的Prototype，不一定是正确的，有待考究
type Prototype struct {
	Property string
}

func (r *Prototype)Print(){
	fmt.Println(r.Property)
}

func (r Prototype)Clone() Prototype{
	return r
}
