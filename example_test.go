package dyno_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/icza/dyno"
)

func Example() {
	person := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Bob",
			"last":  "Archer",
		},
		"age": 22,
		"fruits": []interface{}{
			"apple", "banana",
		},
	}

	// pp prints the person
	pp := func(err error) {
		json.NewEncoder(os.Stdout).Encode(person) // Output JSON
		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}

	// Print initial person and its first name:
	pp(nil)
	v, err := dyno.Get(person, "name", "first")
	fmt.Printf("First name: %v, error: %v\n", v, err)

	// Change first name:
	pp(dyno.Set(person, "Alice", "name", "first"))

	// Change complete name from map to a single string:
	pp(dyno.Set(person, "Alice Archer", "name"))

	// Print and increment age:
	age, err := dyno.GetInt(person, "age")
	fmt.Printf("Age: %v, error: %v\n", age, err)
	pp(dyno.Set(person, age+1, "age"))

	// Change a fruits slice element:
	pp(dyno.Set(person, "lemon", "fruits", 1))

	// Add a new fruit:
	pp(dyno.Append(person, "melon", "fruits"))

	// Output:
	// {"age":22,"fruits":["apple","banana"],"name":{"first":"Bob","last":"Archer"}}
	// First name: Bob, error: <nil>
	// {"age":22,"fruits":["apple","banana"],"name":{"first":"Alice","last":"Archer"}}
	// {"age":22,"fruits":["apple","banana"],"name":"Alice Archer"}
	// Age: 22, error: <nil>
	// {"age":23,"fruits":["apple","banana"],"name":"Alice Archer"}
	// {"age":23,"fruits":["apple","lemon"],"name":"Alice Archer"}
	// {"age":23,"fruits":["apple","lemon","melon"],"name":"Alice Archer"}
}

func ExampleGet() {
	m := map[string]interface{}{
		"a": 1,
		"b": map[interface{}]interface{}{
			3: []interface{}{1, "two", 3.3},
		},
	}

	printValue := func(v interface{}, err error) {
		fmt.Printf("Value: %-5v, Error: %v\n", v, err)
	}

	printValue(dyno.Get(m, "a"))
	printValue(dyno.Get(m, "b", 3, 1))
	printValue(dyno.Get(m, "x"))

	sl, _ := dyno.Get(m, "b", 3) // This is: []interface{}{1, "two", 3.3}
	printValue(dyno.Get(sl, 4))

	// Output:
	// Value: 1    , Error: <nil>
	// Value: two  , Error: <nil>
	// Value: <nil>, Error: missing key: x (path element idx: 0)
	// Value: <nil>, Error: index out of range: 4 (path element idx: 0)
}

func ExampleSet() {
	m := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"3": []interface{}{1, "two", 3.3},
		},
	}

	printMap := func(err error) {
		json.NewEncoder(os.Stdout).Encode(m) // Output JSON
		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}

	printMap(dyno.Set(m, 2, "a"))
	printMap(dyno.Set(m, "owt", "b", "3", 1))
	printMap(dyno.Set(m, 1, "x"))

	sl, _ := dyno.Get(m, "b", "3") // This is: []interface{}{1, "owt", 3.3}
	printMap(dyno.Set(sl, 1, 4))

	// Output:
	// {"a":2,"b":{"3":[1,"two",3.3]}}
	// {"a":2,"b":{"3":[1,"owt",3.3]}}
	// {"a":2,"b":{"3":[1,"owt",3.3]},"x":1}
	// {"a":2,"b":{"3":[1,"owt",3.3]},"x":1}
	// ERROR: index out of range: 4 (path element idx: 0)
}

func ExampleAppend() {
	m := map[string]interface{}{
		"a": []interface{}{
			"3", 2, []interface{}{1, "two", 3.3},
		},
	}

	printMap := func(err error) {
		fmt.Println(m)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}

	printMap(dyno.Append(m, 4, "a"))
	printMap(dyno.Append(m, 9, "a", 2))
	printMap(dyno.Append(m, 1, "x"))

	// Output:
	// map[a:[3 2 [1 two 3.3] 4]]
	// map[a:[3 2 [1 two 3.3 9] 4]]
	// map[a:[3 2 [1 two 3.3 9] 4]]
	// ERROR: missing key: x (path element idx: 0)
}
