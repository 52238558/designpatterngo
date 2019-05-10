package creational

import "fmt"

type IProductA interface {
	OperateA()
}

type IProductB interface {
	OperateB()
}

type ProductA1 struct {
}

type ProductA2 struct {
}

type ProductB1 struct {
}

type ProductB2 struct {
}

type IFactory interface{
	CreateA() IProductA
	CreateB() IProductB
}

type Factory1 struct {
}

type Factory2 struct {
}

func (r *ProductA1)OperateA(){
	fmt.Println("ProductA1 OperateA")
}

func (r *ProductB1)OperateB(){
	fmt.Println("ProductB1 OperateB")
}

func (r *ProductA2)OperateA(){
	fmt.Println("ProductA2 OperateA")
}

func (r *ProductB2)OperateB(){
	fmt.Println("ProductB2 OperateB")
}

func (r *Factory1)CreateA() IProductA{
	return &ProductA1{}
}

func (r *Factory1)CreateB() IProductB{
	return &ProductB1{}
}

func (r *Factory2)CreateA() IProductA{
	return &ProductA2{}
}

func (r *Factory2)CreateB() IProductB{
	return &ProductB2{}
}