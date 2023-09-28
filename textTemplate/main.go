package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func main() {
	fmt.Println("Text Template")
	mySecret := secret{Name: "Jeeva M", Secret: "I am obsessed with GOLANG"}
	fmt.Println("mySecret", mySecret)
	templatePath := `C:/Users/TLSPC-106/Desktop/Gobasics/template.txt`
	t, err := template.New("template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Oops, something wrong!")
		fmt.Println(err) // Print the actual error
		return
	}
	err = t.Execute(os.Stdout, mySecret)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\ntemplate:", t)
	}
}
