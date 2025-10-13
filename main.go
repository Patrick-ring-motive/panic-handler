package main

import (
  "log"
  "fmt"
  "runtime/debug"
)

func init(){
  debug.SetPanicOnFault(true)
}

type Panic struct{
  Recover any
  Message string
  Stack string
}

func (p Panic)Error() string{
  return p.Message
}

type PanicHandler struct{
  Try func()
  Catch func(Panic)
  Finally func()
}

func (p PanicHandler)Handle(){
  if p.Try == nil{
    panic("Try should not be nil")
  }
  debug.SetPanicOnFault(true)
  if p.Finally != nil{
    defer p.Finally()
  }
  if p.Catch != nil{
    defer func(){
      if r := recover(); r != nil{
        p.Catch(Panic{
          Recover: r,
          Message: fmt.Sprintf("Panic recovered: %T %+v",r,r),
          Stack: string(debug.Stack()),
        })
      }
    }()
  }
  p.Try()
}

func main() {

  PanicHandler{
    Try: func(){
      log.Println("Try\n")
      panic(fmt.Errorf("Thrown Error"))
    },
    Catch: func(p Panic){
      log.Printf("Catch\n%+v\n",p.Message)
    },
    Finally: func(){
      log.Println("Finally\n")
    },
  }.Handle()
  
}
