package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define a string flag with a default value and a description
    name := flag.String("name", "World", "Your name")
	fmt.Println(*name)
    
    // Define an integer flag with a default value and a description
    age := flag.Int("age", 30, "Your age")

    // // Parse the flags
    // flag.Parse()

    // Access the flag values
    fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
