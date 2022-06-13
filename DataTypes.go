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
**/
