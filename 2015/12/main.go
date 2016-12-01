package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func main() {
	var data interface{}
	if err := json.NewDecoder(os.Stdin).Decode(&data); err != nil {
		log.Fatalf("Could not decode input: %s", err)
	}

	r := count(data, true)
	log.Printf("Result with reds: %f", r)
	r = count(data, false)
	log.Printf("Result without reds: %f", r)
}

func count(data interface{}, reds bool) float64 {
	var r float64
	switch reflect.ValueOf(data).Kind() {
	case reflect.Map:
		var o float64
		for _, v := range data.(map[string]interface{}) {
			o += count(v, reds)
			if v == "red" && !reds {
				return 0
			}
		}
		r += o
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for _, v := range data.([]interface{}) {
			r += count(v, reds)
		}
	case reflect.Float64:
		return data.(float64)
	}
	return r
}
