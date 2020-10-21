package main

import (
	"fmt"
	"math/big"
)

func main()  {
	 var a,b int64
	 a = 1;
	 b = 3;
	 fmt.Println("a,b值是：",a,b)

	 a1:=big.NewInt(a)//大整数
	 fmt.Println( "a的大整数是",a1)

	 var c big.Int
	 c1:=big.NewInt(199)
	 c2:=big.NewInt(100)
	 c3:=big.NewInt(5)
	 //rs:=(c1,c2)
	 bijiao:=c1.Cmp(c2)
	 fmt.Println(bijiao)
	 rs:=c.Add(c1,c2)
	 fmt.Println(rs)

	 sub:=c.Sub(c2,c1)
	 fmt.Println(sub)

	 n:=c.Lsh(c3,2)
	 fmt.Println("n的值是：",n)

}
