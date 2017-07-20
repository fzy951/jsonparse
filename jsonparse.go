package jsonparse

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JsonData map[string]interface{}

func setJsonData(m map[string]interface{}, val interface{}, keys []string) error {
	length := len(keys)
	if length == 0 {
		return fmt.Errorf("The keys is empty!")
	}
	key := keys[0]
	if length == 1 {
		m[key] = val
	} else {
		m1, ok := m[key].(map[string]interface{})
		if !ok {
			m1 = make(map[string]interface{})
		}
		setJsonData(m1, val, keys[1:])
		m[key] = m1
	}
	return nil
}

type jsonData struct {
	jd   *JsonData
	keys []string
}

func (jsd *jsonData) Set(value interface{}) error {
	return jsd.jd.Set(value, jsd.keys...)
}

func (jsd *jsonData) Get() (interface{}, error) {
	return jsd.jd.Get(jsd.keys...)
}

func (jd *JsonData) Key(keys ...string) *jsonData {
	return &jsonData{jd: jd, keys: keys}
}

func (jd *JsonData) Set(value interface{}, keys ...string) error {
	if len(keys) == 0 {
		m, ok := value.(map[string]interface{})
		if ok {
			*jd = m
			return nil
		} else {
			return fmt.Errorf("The keys is empty and the value is not a json map!")
		}
	}
	val := *jd
	err := setJsonData(val, value, keys)
	*jd = val
	return err
}

func (jd *JsonData) Get(keys ...string) (interface{}, error) {
	val := *jd
	length := len(keys)
	if length == 0 {
		return val, nil
	}
	for i, key := range keys {
		ret, ok := val[key]
		if !ok {
			return nil, fmt.Errorf("There's no key <%s> exist in the json data", strings.Join(keys[:i+1], "."))
		}
		if i == length-1 {
			return ret, nil
		}
		val, ok = ret.(map[string]interface{})
		if ok {
			continue
		} else {
			return nil, fmt.Errorf("The key <%s> is not a json map,value: %v", strings.Join(keys[:i+1], "."), ret)
		}
	}
	return nil, fmt.Errorf("Unexcept operation...")
}

func (jd *JsonData) Marshal() ([]byte, error) {
	return json.Marshal(*jd)
}

func (jd *JsonData) Unmarshal(data []byte) error {
	return json.Unmarshal(data, jd)
}
