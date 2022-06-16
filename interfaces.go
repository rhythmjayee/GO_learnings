func main()  {
  main1()
  main2()
}
func main1()  {
  //variable with type interface
  //Using interface functions with Type struct
  var w Writer = ConsoleWriter{}
  w.Write([]byte("Hello World"))
}
//Interfce with name Writer
//Describe behaviour
type Writer interface{
  //paramters   //return vals
  Write([]byte)(int, error)
}
//structure
type ConsoleWriter struct{}

//func associated with ConsoleWriter ->method  struct using interface
func (cw ConsoleWriter) Write(data []byte)(int, error)  {
  n, err := fmt.Println(string(data))
  return n, err
}

//any type can have methods associated
func main2()  {
  myInt := IntCounter(0)
  //if any method has pointer reciever then & is important
  var inc Incremeter = &myInt
  fmt.Println(inc.Incremeter())
}
type Incremeter interface {
  Incremeter() int
}
//type variable
type IntCounter int

func(ic *IntCounter) Incremeter() int{
  *ic++
  return int(*ic)
}
//Interfaces Examples
//io.Writer
//io.Reader
//interface{}


//composing Interfaces -> add interfaces inside interface
type Writer interface{
  //paramters   //return vals
  Write([]byte)(int, error)
}
type Closer interface{
  //paramters   //return vals
  Close() error
}

type WriterCloser interface{
  Writer
  Closer
}
