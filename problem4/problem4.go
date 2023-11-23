package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	// your code here
	var nameEncode = ""
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			nameEncode += string((char-'a'+3)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			nameEncode += string((char-'A'+3)%26 + 'A')
		} else {
			nameEncode += string(char)
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	// your code here
	var nameDecode = ""
	for _, char := range s.nameEncode {
		if char >= 'a' && char <= 'z' {
			nameDecode += string((char-'a'+23)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			nameDecode += string((char-'A'+23)%26 + 'A')
		} else {
			nameDecode += string(char)
		}
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
