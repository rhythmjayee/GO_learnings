//if statements
func main()  {
  m := make(map[int][int]);
  m[0] = 0;
  m[1] = 1;

  //first part is intializing pop,ok
  //2nd part checks boolean
  if pop, ok := m[0]; ok{
    fmt.Println(pop)//0
  }
  //or
  if 0 == 1 {
    fmt.Println("In if")
  }else if 1 == 2{
    fmt.Println("In else if")
  }else{
    fmt.Println("In else ")
  }

  //switch
  var_name := 3
  switch var_name {
    case 1,5,10 : fmt.Println("Case 1")
    case 2 : fmt.Println("Case 2")
    case 3 : fmt.Println("Case 3")
    default : fmt.Println("Default")

  }
  //other way
  switch {
    case var_name == 1 :
          fmt.Println("Case 1")
          fallthrough //also checks other cases
    case var_name == 2 : fmt.Println("Case 2")
    case var_name == 3 : fmt.Println("Case 3")
    default : fmt.Println("Default")

  }

  var i interface{} = 1.0
    switch i.(type){
        case float64 :  fmt.Println("Case float")
        default : fmt.Println("unknown type")
    } //float64


    for j := 0; j<=10; j++ {
        fmt.Println("In loop",j)
    }

    for j,k := 0,0; j<=10; j,k = j+1,k+1 {
        fmt.Println("In loop",j,k)
    }

    j := 0
    for ; j<5; {
        fmt.Println("In while type loop",j)
        j++
    }
    //Infinite loop
    // for{
    //
    // }

//label
  Loop:
    for j := 0; j<=10; j++ {
        fmt.Println("In loop",j)
        InLoop:
        for k := 0; k<=10; k++ {
            fmt.Println("In loop",k)
            if k == 5{
               break InLoop
            }else if j == 5{
              break Loop
            }
        }
    }


    //loop over collections
    s := []int{1,2,3}
    for k,v := range s{
      fmt.Println(k,v)
    }

    str := "String"
    for k,v := range s{
      fmt.Println(k,string(v))
    }

    //when need single variable
    ss := []int{1,2,3}
    for _,v := range ss{
      fmt.Println(v)
    }

}
