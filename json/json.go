package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Encoding basic datatypes

	// Encoding bool
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // -> true

	// Encoding int
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // -> 1

	// Encoding float
	floatB, _ := json.Marshal(1.32)
	fmt.Println(string(floatB)) // -> 1.32

	// Encoding string
	stringB, _ := json.Marshal("golang")
	fmt.Println(string(stringB)) // -> "golang"

	// Encoding slice
	slcD := []string{"peach", "punch"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // -> ["peach","punch"]

	// Encoding map
	mapD := map[string]string{"key1": "value1", "key2": "value2"}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // -> {"key1":"value1","key2":"value2"}

	// Encoding custom datatype
	// Only exported fields are marshalled
	// Field names used as key
	res1D := &response1{1, []string{"apple", "peach", "orange"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // -> {"Page":1,"Fruits":["apple","peach","orange"]}

	// Encoding custom datatype
	// Only exported fields are marshalled
	// customize key names near struct fields using -> {field} {field_type} `json:{key_name}`
	res2D := &response2{1, []string{"apple", "peach", "orange"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // -> {"page":1,"fruits":["apple","peach","orange"]}

	// Decoding data

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Variable to hold decoded JSON data
	// key -> string
	// value -> interface (arbitrary datatype)
	var dat map[string]interface{}

	// Check for err before proceeding
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat) // -> map[num:6.13 strs:[a b]]

	// Cast value in map to associate type before using
	num := dat["num"].(float64)
	fmt.Println(num) // -> 6.13

	// Accessing nested data requires series of casts
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1) // -> a

	// Decode JSON to custom type
	// Automatically take care of type safety
	str := `{"page": 1, "fruits":["peach", "orange"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           // -> {1 [peach orange]}
	fmt.Println(res.Fruits[0]) // -> peach

	// stream JSON encoding to stdout, stderr or http response
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "orange": 10}
	// marshall json and write to stdout
	enc.Encode(d) // -> {"apple":5,"orange":10}
}
