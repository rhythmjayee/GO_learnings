func main()  {
  //Pointers
  //Cant do pointer arithmetic -> but go gives unsafe package to that
    var a int = 1
    var b *int = &a//pointer b, holds a's address
    fmt.Println(a,*b)//dereferncing b
    a = 2
    fmt.Println(a,*b)//2 2
    b = 3
    fmt.Println(a,*b)//3 3
    fmt.Printf("%p",b)//prints address



    var ms *myStruct
     fmt.Println(ms)//default value of pointer -> nil
     ms = &myStruct{foo : 1}
     fmt.Println(ms)//&{1}

     //using new
     ms = new(myStruct)//-> intilize with default
     fmt.Println(ms)//&{0}
     (*ms).foo = 42
     //or ms.foo = 42.. same for getting value
     fmt.Println((*ms).foo)




}
   type myStruct struct{
       foo int
   }
