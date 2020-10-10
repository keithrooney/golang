package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// The zero value of a map is nil
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {

	fmt.Println(m)

	another := make(map[string]int)

	another["Answer"] = 42
	fmt.Println(another)

	another["Answer"] = 14
	fmt.Println(another)

	delete(another, "Answer")
	fmt.Println(another)

	v, ok := another["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
