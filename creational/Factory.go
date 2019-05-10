package creational

import "fmt"

type IProduct interface{
	Operate()
}

type ProductA struct {
}

type ProductB struct{
}

func (r *ProductA)Operate(){
	fmt.Println("A Operate")
}

func (r *ProductB)Operate(){
	fmt.Println("B Operate")
}

func Creat(ptype string) IProduct{
	switch ptype{
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		fmt.Println("not a valid type")
	}
	return nil
}

