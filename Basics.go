/*
Why Go ?
Python -> easy to use but slow
Java -> With large system becomes complex
C/C++ -> Complex + Slow compile time

Go ->
Strong & Statically typed
Simple
Fast compile time
garbage collected
Built in concurrency - Better at handle concurrency
Compile to standalone binaries
*/

//variables
var glb_var int = 0;
func main() {
    //declare it is variable, variable name, variable type
    var n1 int
    n1 = 3
    var n2 float32 = 4
    //default type int
    n3 := 5
    //float64 type
    n4 := 5.0
    fmt.Println(n1,n2,n3)
    //Printf -> format string
    //"%v %T" -> value, type
    fmt.Printf("%v %T",n4,n4)

    //wrapp variables
    var(
        name string = "Rhythm"
        age int = 21
    )
    fmt.Println(name,age)

    //shadowing of variable -> internal var take precedence
    fmt.Println(glb_var) //0
    var glb_var int = 99;
    fmt.Println(glb_var)//99

    //Scope of variables
    // lower case first letter -> scoped with package
    //Upper case first letter -> exported at the front of package, globally visible
    //lock scope -> {} variable declared inside


    //type conversions
    /*
    int()
    float32()
    int -> string == string() -> return ascii
    */
   var intt int = 97
   fmt.Printf("%v %T\n",intt,intt)
   var floatt float32
   floatt = float32(intt)
   fmt.Printf("%v %T\n",floatt,floatt)

   var str string = string(intt)
   fmt.Println(str)//a->ascii

   //strconv package
   str = strconv.Itoa(intt)
   fmt.Printf("%v %T\n",str,str)


}
