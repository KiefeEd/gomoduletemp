package gomoduletemp
import(
	"fmt"
	"time"
) 
type User struct {
	name string
	age int8
}
func (user User) Hello() {
	fmt.Printf("Hello, %s!", user.name)
}

func (user User) Greet(fromutc int) {
	if fromutc < -11 || fromutc > 13 {
		fmt.Println("Invalid Timezone from UTC provided. Example input: 7")
	} else {
	var timenow = time.Now().UTC().Hour() + fromutc //GMT+7
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
		fmt.Printf("Good %s, %s!", phase, user.name)
	}	
}
func (user User) GreetLocal() {
	var timenow = time.Now().Local().Hour() 
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
		fmt.Printf("Good %s, %s!", phase, user.name)
	}	
}
func DisplayTime(fromutc int) {
	if fromutc < -11 || fromutc > 13 {
		fmt.Println("Invalid Timezone from UTC provided. Example input: 7")
	} else {
	hournow := time.Now().UTC().Hour() + fromutc
	minutenow := time.Now().UTC().Minute()
	fmt.Printf("%d:%d", hournow, minutenow)
	}
}
func DisplayTimeLocal() {
	hournow := time.Now().Local().Hour()
	minutenow := time.Now().Local().Minute()
	fmt.Printf("%d:%d", hournow, minutenow)
}