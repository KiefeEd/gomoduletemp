package main
import(
 timehelper "github.com/KiefeEd/gomoduletemp"
)

main() {
  go Greet("Kiefe", 7)
  go DisplayTimeLocal()
}
