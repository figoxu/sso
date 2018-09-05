package main

import (
	"testing"
	"log"
)

func TestNewSaltHelper(t *testing.T) {
	sh := NewSaltHelper("hello", "world")
	content := "figo'xu awesome"
	result := sh.Encode(content)
	log.Println(result)
	val := sh.Decode(result)
	log.Println(val)
	log.Println(val == content)
}

func TestBasicAuthDecode(t *testing.T) {
	s := "hello#world,how_are_you"
	content := BasicAuthEncode(s)
	log.Println(content)
	result := BasicAuthDecode(content)
	log.Println(result)
	log.Println(result == s)
}
