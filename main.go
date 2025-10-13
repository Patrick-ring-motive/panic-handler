package main

import (
  "log"
  "fmt"
  "github.com/Patrick-ring-motive/panic-handler/panics"
)


func main() {

  panics.PanicHandler{
    Try: func(){
      log.Println("Try\n")
      panic(fmt.Errorf("Thrown Error"))
    },
    Catch: func(p panics.Panic){
      log.Printf("Catch\n%+v\n",p.Message)
    },
    Finally: func(){
      log.Println("Finally\n")
    },
  }.Handle()
  
}
