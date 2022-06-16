func main()  {
  //entry point of go application -> main function
  a := 1
  fmt.Println("Val of a",a)
  passByVal(a)
  fmt.Println("Val of a after passByVal ",a)
  passByRef(&a)
  fmt.Println("Val of a after passByRef ",a)

  helper("Hello",2,3,4,5);//Hello [2 3 4 5]

  sum := helper2(1,2,3,4)
  fmt.Println(sum)

  //Anonymous function
  for i := 1; i<=5; i++ {
    func (i int)  {
        fmt.Println(i)
    }(i)
  }

//variable type function
var f func() = func (i int)  {
    fmt.Println(i)
}
f(1)
//or
 // f := func (i int)  {
 //     fmt.Println(i)
 // }

 mm := myStruct{}
 mm.name = "rhythm"
 mm.print();
}

//functions as methods
type myStruct struct{
  name string
}
//method
func (ms myStruct) print()  {
  //ms is copy
  fmt.Println(ms.name)
}
//or
func (ms *myStruct) print()  {
  //ms is same struct
  fmt.Println(ms.name)
}


func helper(name string, values ...int)  {//Variatic parameters
  //values will be slice of numbers
  fmt.Println(name,values)

}
//func with return type
func helper2(values ...int) int {//Variatic parameters
  //values will be slice of numbers
  sum := 0
  for _,val := range values{
    sum = sum + val
  }
  return sum
  //if pointer is returned from function then memory will asssigned in heap
}
func passByRef(a *int)  {
  *a = 2
}
func passByVal(a int)  {
  a = 2
}
