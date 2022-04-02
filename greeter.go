package gomoduletemp
import(
	"fmt"
	"time"
) 
type Timezone struct {
	use_local bool
	fromutc int8
}
func Hello(user string) {
	fmt.Printf("Hello, %s!", user.name)
}

func Greet(user string, fromutc int) {
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
func GreetLocal(user string) {
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

func (t Timezone) TimeOfTheDay()(string) {
	if t.use_local == true {
		var timenow = time.Now().Local().Hour() 
	} else {
		var timenow = time.Now().UTC().Hour() + t.fromutc
	}

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
	}