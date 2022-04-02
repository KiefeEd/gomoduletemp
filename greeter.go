package gomoduletemp
import(
	"fmt"
	"time"
) 
type Timezone struct {
	use_local bool
	fromutc int
}
func Hello(user string) {
	fmt.Printf("Hello, %s!", user)
}

func Greet(user string, fromutc int) {
    tz := Timezone{
    use_local:false,
    fromutc:fromutc,
  }
		fmt.Printf("Good %s, %s!", tz.TimeOfTheDay(), user)
	
}
func GreetLocal(user string) {
	totd := Timezone{
		use_local: true,
	}
 
		fmt.Printf("Good %s, %s!",totd.TimeOfTheDay(), user)
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
  timenow := 0
	if t.use_local == true {
		timenow = time.Now().Local().Hour() 
  
	} else {
		timenow = time.Now().UTC().Hour() + t.fromutc
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
  return phase;
	}
	




