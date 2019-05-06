package bit

import (
	"encoding/json"
	"fmt"
)

const (
	IsA uint8 = 1
	IsB uint8 = (1 << 1)
)

type Result struct {
	MyFlag
	FlagResult
}

type FlagResult struct {
	IsA bool `json:"isA"`
	IsB bool `json:"isB"`
}

type MyFlag struct {
	Flags uint8 `json:"flags" comment:"标志位，从低到高，第1位：a，第2位：b"`
}

func ExampleDecode() {

	data := []byte(`{"isA": false, "isB": false}`)
	obj := Result{}
	json.Unmarshal(data, &obj)
	EncodeFlags(&obj.Flags, obj.FlagResult.IsA, obj.FlagResult.IsB)
	fmt.Println(obj.Flags)

	data = []byte(`{"isA": true, "isB": false}`)
	obj = Result{}
	json.Unmarshal(data, &obj)
	EncodeFlags(&obj.Flags, obj.FlagResult.IsA, obj.FlagResult.IsB)
	fmt.Println(obj.Flags)

	data = []byte(`{"isA": false, "isB": true}`)
	obj = Result{}
	json.Unmarshal(data, &obj)
	EncodeFlags(&obj.Flags, obj.FlagResult.IsA, obj.FlagResult.IsB)
	fmt.Println(obj.Flags)

	data = []byte(`{"isA": true, "isB": true}`)
	obj = Result{}
	json.Unmarshal(data, &obj)
	EncodeFlags(&obj.Flags, obj.FlagResult.IsA, obj.FlagResult.IsB)
	fmt.Println(obj.Flags)

	obj2 := Result{MyFlag{Flags: 0}, FlagResult{}}
	DecodeFlags(obj2.Flags, &obj2.FlagResult.IsA, &obj2.FlagResult.IsB)
	fmt.Println(obj2.IsA, obj2.IsB)

	obj2 = Result{MyFlag{Flags: 1}, FlagResult{}}
	DecodeFlags(obj2.Flags, &obj2.FlagResult.IsA, &obj2.FlagResult.IsB)
	fmt.Println(obj2.IsA, obj2.IsB)

	obj2 = Result{MyFlag{Flags: 2}, FlagResult{}}
	DecodeFlags(obj2.Flags, &obj2.FlagResult.IsA, &obj2.FlagResult.IsB)
	fmt.Println(obj2.IsA, obj2.IsB)

	obj2 = Result{MyFlag{Flags: 3}, FlagResult{}}
	DecodeFlags(obj2.Flags, &obj2.FlagResult.IsA, &obj2.FlagResult.IsB)
	fmt.Println(obj2.IsA, obj2.IsB)

	// Output:
}

func ExampleEncode() {

	flag := uint8(0)
	EncodeFlags(&flag, true, true, true, true, true, true, true, true)
	fmt.Println(flag)

	flag = uint8(0)
	EncodeFlags(&flag, true, true, true, true, true, true, true, true, true)
	fmt.Println(flag)

	flag = uint8(255)
	var b1, b2, b3, b4, b5, b6, b7, b8, b9 bool
	DecodeFlags(flag, &b1, &b2, &b3, &b4, &b5, &b6, &b7, &b8, &b9)
	fmt.Println(b1, b2, b3, b4, b5, b6, b7, b8, b9)
	DecodeFlags(flag, &b1, &b2, &b3, &b4, &b5, &b6, &b7, &b8)
	fmt.Println(b1, b2, b3, b4, b5, b6, b7, b8)

	flag = uint8(9)
	DecodeFlags(flag, &b1, &b2, &b3, &b4)
	fmt.Println(b1, b2, b3, b4)

	// Output:
}
