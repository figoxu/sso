package main

import (
	"testing"
	"fmt"
)

func TestUrlAppendParam(t *testing.T) {
	v := urlAppendParam("http://figoxu.me/welcome.jsp?auth=true&zoom=99&age=18&male=true", "token", "123456trewq4rfvCDE#")
	fmt.Print(v)
}
