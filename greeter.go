package gomoduletemp
import(
	"fmt"
	"time"
) 

func Hello(user string) {
	fmt.Printf("Hello, %s!", user)
}

func Greet(user string) {
	var timenow = time.Now().UTC().Hour() + 7 //GMT+7
	phase := "Midnight"
	if timenow < 6 || timenow > 22{
		phase = "Night"
	} else if timenow >= 6 && timenow < 12 {
		phase = "Morning"
	} else if timenow >= 12 && timenow < 17 {
		phase = "Afternoon"
	} else {
		phase = "Evening"
	}
		fmt.Printf("Good %s, %s!", phase, user)
	
}