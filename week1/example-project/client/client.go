package main

import (
	"client/helper"
	"fmt"
	"strconv"
	"util"
)

/*
*
To run this main function client must be in "package main", see line 1
*/
func main() {

	strNum := "5"
	num, err := strconv.Atoi(strNum)
	//use the check function from 'client extended
	check(err)

	//Call the `Add` function from the helper package
	res := helper.Add(num, num)

	//Try uncommenting the line below to see that the `add` function is not exported
	// res := helper.add(num, num)
	fmt.Printf("We got: %d \n", res)

	// Call the `Mult` function from the util module
	res2 := util.Mult(num, num)
	fmt.Printf("And then: %d \n", res2)

	x := util.Mult(3, 4)
	fmt.Println(x)
}
