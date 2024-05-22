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

import (
	"fmt"
	"sync"
)

type Singleton struct {
    data string
    once sync.Once
}

func getInstance() *Singleton {
    var instance *Singleton
    instance = &Singleton{}
    instance.once = sync.Once{}
    return instance
}

func main() {
    singleton := getInstance()
    singleton.once.Do(func() {
        singleton.data = "Hello from Singleton!"
    })
    fmt.Printf("singleton: %s \n", singleton.data)

    anotherSingleton := getInstance()
    fmt.Printf("anotherSingleton: %s \n", anotherSingleton.data)
    fmt.Printf("singleton: %s \n", singleton.data)

}