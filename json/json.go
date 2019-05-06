package json

import (
	"encoding/json"
	"fmt"
)

func IndentJson(objs interface{}) {

	data, err := json.MarshalIndent(objs, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
