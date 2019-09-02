package main

// import "fmt"
// meerdere packages invoegen
import (
	"fmt"
	"math"

	"github.com/marijnst/go_crash_course/03_packages/strutils"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	// eigen package gebruiken
	fmt.Println(strutils.Reverse("olleh"))
}
