package main

import "fmt"

func main() {
	vp := readConfig()
	fmt.Println(vp.GetString("db.uril"))
}
