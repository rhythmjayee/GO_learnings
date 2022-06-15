func main()  {
    /*
    start
    end
    middle
    */
    //defer ->delay the execution
    fmt.Println("start")
    defer fmt.Println("middle") //run later after all the things has been executed
    fmt.Println("end")

    /*
    Last in First out -> of all defer statements

    end
    middle
    start
    */
    defer fmt.Println("start")
    defer fmt.Println("middle")
    defer fmt.Println("end")

//defer stores value of variables till when is called
    a := 5
    defer fmt.Println(a) // 5
    a = 10


    //panic() is used to throw error
    //happens after defer statements
    //It stops the execution of function in which panic happended
    //but flow of calling function of function in which panic occured continues
    panic("Error....")



    //recover() -> handles panic
    //useful inside defer functions
    fmt.Println("start")//start
    paincker()//Error....
    fmt.Println("end")//end

}
func paincker()  {
  defer func(){
      err := recover()
      if err != nil{
          fmt.Println(err)
          // panic(err)
          //panic below function in class stack also
          //then end wont be printed
      }
  }()
  panic("Error....")
}
