package jsonparse

import (
	"fmt"
	"testing"
)

func Test_Parse(t *testing.T) {
	jd := JsonData{}
	fmt.Println(jd.Set("hello"))
	fmt.Println(jd.Set(map[string]interface{}{"h": "2"}))
	fmt.Println(jd)

	jd.Key("key").Set("hello")
	fmt.Println(jd.Key("key").Get())

	jd.Set("hello", "world")
	jd.Set("asd", "1", "2", "3")
	jd.Set("asd", "1", "3")
	fmt.Println(jd.Get("world"))
	fmt.Println(jd.Get("1", "2", "3", "4"))
	fmt.Println(jd.Get("hhh"))
	fmt.Println(jd)

	jd.Set("helo", "wo")
	jd.Set("hello world", "Peersafe_data")

	data, err := jd.Marshal()
	fmt.Println(string(data), "\n***", err)

	jd1 := JsonData{}
	err = jd1.Unmarshal(data)
	fmt.Println(err, "\n***", jd1)
}
