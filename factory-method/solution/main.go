// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Singleton struct {
// 	data string
// 	once sync.Once
// }

// func GetInstance() *Singleton {
// 	var instance *Singleton
// 	instance = &Singleton{}

// 	instance.once = sync.Once{}
// 	return instance
// }

// func main() {
// 	singleton := GetInstance()
// 	singleton.once.Do(func() {
// 		singleton.data = "Hello world \n"
// 	})

// 	anotherSingleton := GetInstance()

// 	fmt.Printf(singleton.data)
// 	fmt.Printf(anotherSingleton.data)
// }

package main

import "fmt"

type TransportService struct {
    transport Transport
}

type Transport interface {
    Deliver()
}

type Truck struct {}

type Ship struct {}

func (Truck) Deliver() {
    fmt.Printf("Deliver by Truck\n")
}

func (Ship) Deliver()  {
    fmt.Printf("Deliver by Ship\n")
}

func CreateDeliver(t string) Transport {
    if t == "TRUCK" {
        return Truck{}
    } else if t == "SHIP" {
        return Ship{}
    }

    return nil
}

func main() {
    t := TransportService{
        transport: CreateDeliver("SHIP"),
    }

    t.transport.Deliver()
}