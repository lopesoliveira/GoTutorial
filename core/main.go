package main

import (
	"GoTutorial/api"
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	var myString = "João Oliveira" //13 characters
	fmt.Println(myString)
	fmt.Println(len(myString))                    //14 - because ã from João occupies 2 bytes
	fmt.Println(utf8.RuneCountInString(myString)) //13 - The right number of characters the string has

	var p *int32 = new(int32)
	var i int32 = 2
	//*p = i
	//p = &i
	fmt.Printf("The value of p is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

	//maps()

	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSliceIndex(intSlice))
	fmt.Println(sumSlice(intSlice))

	var floatSlice []float64 = []float64{1.1, 2.2, 3.3}
	fmt.Println(sumSliceGeneric(floatSlice))
	fmt.Println(sumSliceGeneric(intSlice))

	var stringSlice = []string{"a", "b", "c"}
	fmt.Println(sumSliceGeneric(stringSlice))

	//var mySlice = []int{}  //GoLand prefers the 2nd option to declare an empty slice
	var mySlice []int
	fmt.Println(isSliceEmpty[int](mySlice))

	api.WriteSomething() // To execute a function from another package the function need to be capitalized. writeSomething won't work. WriteSomething will work
}

// Maps
func maps() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}

	for name, ageLoop := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, ageLoop)
	}
}

/**************************************************************************************************

Go How can I iterate over a map and preserve the order

In Go, maps are inherently unordered, meaning that the iteration order over map keys is not guaranteed to be consistent. If you need to iterate over a map in a specific order, you can achieve this by using a separate slice to store the keys in the desired order and then iterate over that slice.

Here’s how you can do it:

Create a map.
Extract the keys and sort them (if needed).
Iterate over the sorted keys and access the map values.

Example:

package main

import (
    "fmt"
    "sort"
)

func main() {
    // Create a map
    myMap := map[string]int{
        "apple":  5,
        "banana": 3,
        "cherry": 7,
    }

    // Extract the keys into a slice
    keys := make([]string, 0, len(myMap))
    for key := range myMap {
        keys = append(keys, key)
    }

    // Sort the keys (optional, if you need a specific order)
    sort.Strings(keys)

    // Iterate over the sorted keys and access the map values
    for _, key := range keys {
        fmt.Printf("%s: %d\n", key, myMap[key])
    }
}

Explanation:
Create a map: The map myMap is created with some key-value pairs.
Extract the keys: The keys of the map are extracted into a slice keys.
Sort the keys: The keys are sorted using sort.Strings(keys). This step is optional and depends on whether you need the keys in a specific order.
Iterate over the sorted keys: The sorted keys are iterated over, and the corresponding values from the map are accessed and printed.

Output
apple: 5
banana: 3
cherry: 7


By using this approach, you can control the order in which you iterate over the map.
If you need a different order (e.g., numerical order for integer keys), you can adjust the sorting step accordingly.

 *************************************************************************************************/
