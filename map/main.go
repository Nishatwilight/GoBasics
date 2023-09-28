package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m)
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(len(m))
	// Another way to instantiate the map is
	makeMap := make(map[string]string, 6)
	fmt.Println("makeMap Length", len(makeMap))
	makeMap = map[string]string{"nisha": "Programmer", "Jeeva": "Designer"}
	fmt.Println("makeMap with string: ", makeMap)
	makeMap["Tamil"] = "Senior Developer"
	fmt.Println("One value is add in makeMap: ", makeMap) // O/p One value is add in makeMap:  map[Jeeva:Designer Tamil:Senior Developer nisha:Programmer]
	//Default map is not give the ordered value
	makeMap["Tamil"] = "Super Senior Developer"
	fmt.Println("Modified a value in makeMap and checking: ", makeMap) // Modified a value in makeMap and checking:  map[Jeeva:Designer Tamil:Super Senior Developer nisha:Programmer]
	fmt.Println("Extracting a value using the key: ", makeMap["nisha"])

	//We can even Delete the Map value using delete()
	delete(makeMap, "Tamil")
	fmt.Println("Checking if the delete works :", makeMap) //Checking if the delete works map[Jeeva:Designer nisha:Programmer]
	makeMap["Tamil"] = " Senior Developer"
	makeMap["Tamil"] = "Super" + makeMap["Tamil"]
	fmt.Println("makeMap string concadination : ", makeMap) // Senior Developer Super
	delete(makeMap, "Tamil")

	//Maps are iterable Datastructure, sowe can use for each loops
	for name, occu := range makeMap {
		fmt.Println(name, occu)
	}

	occu, ok := makeMap["Tamil"]
	if ok {
		fmt.Println("okay concept", occu)
	} else {
		fmt.Printf("Value Does not exist")
	}
}
