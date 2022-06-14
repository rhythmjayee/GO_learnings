/*
types :
  primitive
    bool
    int
    uint (unsigned int)
    float
    complex64
      real() , imag()
    string -> immutable
      []byte(string_name) -> array of byte of ascii
    rune  -> r:= 'a'  -> type int32, value = 97


  constants
  const
    cant change value
    can be shadowed
  iota -> enumerated constants
    const( a = iota -> assign 0 value if its first
           b = iota -> assign 1 value
    )
    if they are in const block
    else all const have 0 as value

    _ -> write only value


    arrays
      var names [3]int
      fmt.Printf("arr %v",names)//arr [0 0 0]

      arr_name := [3]int{0,1,2}
      arr_name := [...]int{0,1,2,3}
      fmt.Printf("arr %v",arr) //arr [0 1 2 3]

      var names [3]string
      len(names) //length of array

    //multi-d arrays
      var matrix [3][3]int
      matrix[0] = [3]int{1,0,0}
      matrix[1] = [3]int{0,1,0}
      matrix[2] = [3]int{0,0,1}

      ref_var := matrix //deep copy of matrix
      ref_var := &matrix// same matrix


    Slices
      slice_name := []int{0,1,2,3,4,5,6,7,8,9,10}
      cap(slice_name)
      len(slice_name)

      ref_slice := slice_name //same slice

      //slicing operations ->also work on arrays

      ref_slice := slice_name[:]//whole slice
      ref_slice := slice_name[3:]//4th ele to end
      ref_slice := slice_name[:6]//first 6 ele
      ref_slice := slice_name[3:6]//[3,6)

      append(slice_name,elements by commas)
      //return new slice if cap < than number of elements
        slice_name = append(slice_name,1,2,3)

        //using spread operator
        slice_name = append(slice_name,int[]{1,2,3}...)


    make(type_of_slice,len,cap)
      a := make([]int,3,100);


    Slice Examples :
    //slicing op can change the original slice
      a := []int{1,2,3,4,5}
      b := append(a[:2], a[3:]...)



**/
