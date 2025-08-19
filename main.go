package main

import (
	"fmt"

	"github.com/lucaslimafernandes/clt-pj-calc/internal/readf"
)

func main() {

	r := readf.ReadToml()

	fmt.Println(r)

}
