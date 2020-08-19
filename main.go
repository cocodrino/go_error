package main

/*import (
	"fmt"
)


type IFoo interface {
	Call()
	WhatColor()
	DoubleCall()
}

// Foo es como una clase abstracta, solo implementa DoubleCall de la interface y otros métodos útiles
type Foo struct {
	base IFoo
	color string
}


func(f Foo) DoubleCall(){
	fmt.Println("double call")
	f.base.Call()
	f.base.Call()
}

func (f Foo) WhatColor(){
	fmt.Println("Color is ",f.color)
}

func (f *Foo) AddColor(color string){
	f.color = color
}

// baz compone a Foo agregandole los métodos para cumplir la interfaz IFoo
type Baz struct {
	Foo
}

func (Baz) Call() {
	fmt.Println("Baz Called")
}

func main() {
	baz := Baz{}
	baz.base = baz
	baz.DoubleCall() // prints "Baz Called"
	baz.AddColor("red")
	baz.WhatColor()
}

*/
import (
	"awesomeProject/src/DataProvider/HttpProviders"
	"awesomeProject/src/DataReceiver/DataReceivers"
)

func main()  {
	store := DataReceivers.NewTestStore()
	marketData := HttpProviders.NewMarketStack("d57d10f3d04b7b798ad322b7ffc8a139",[]string{"AAPL"},30)
	marketData.ConnectToReceiver(store)
	marketData.RetrieveData()

}
