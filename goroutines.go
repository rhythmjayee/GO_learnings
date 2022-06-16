/*
Inside of the go runtime, we have got a scheduler
thats going to map these go routines onto these
operating system threads for periods of time and the
scheduler will then take turns with every CPU thread thats
available and assign the different go routines
a certain amount of processing time on those threads

*/
func main1()  {
  //green thread -> goroutine
  go sayHello()

//race condition -> msg value change
  var msg = "Hello"
  //form closure
  //has access to msg
  go func ()  {
    //msg will hold latest val
    //if changed
    fmt.Println(msg)
  }()
  masg = "New Hello"



//to avoid change in val
  var msg = "Hello"
  go func (msg string)  {
    //msg will hold latest val
    //if changed
    fmt.Println(msg)
  }(msg)
  msg = "New Hello"
}
func sayHello()  {
  fmt.Println("Hello")
}


//race condition example
var wg = sync.WaitGroup{}
var counter = 0
func main2()  {
  sync_funs()
}
func sync_funs()  {
    for i := 0; i<10; i++{
      wg.Add(2)
      go sayHello()
      go increment()
    }
    wg.Wait()
}



func increment()  {
  counter++
  wg.Done()
}
func sayHello()  {
  fmt.Println("Hello",counter)
  wg.Done()
}





//go routines with mutex
// but still both functions are not in sync
//just race condition is avoided
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex
func main3()  {
  sync_funs()
}
func sync_funs()  {
    for i := 0; i<10; i++{
      wg.Add(2)
      go sayHello()
      go increment()
    }
    wg.Wait()
}



func increment()  {
  m.Lock()
  counter++
  m.Unlock()
  wg.Done()
}
func sayHello()  {
  m.RLock()
  fmt.Println("Hello",counter)
  m.RUnlock()
  wg.Done()
}




//Sync Go routines
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex
func main4()  {
  sync_funs()
}
func sync_funs()
//set system threads
  runtime.GOMAXPROCS(100)
    for i := 0; i<10; i++{
      wg.Add(2)
      m.RLock()
      go sayHello()
      m.Lock()
      go increment()
    }
    wg.Wait()
}



func increment()  {
  counter++
  m.Unlock()
  wg.Done()
}
func sayHello()  {
  fmt.Println("Hello",counter)
  m.RUnlock()
  wg.Done()
}
