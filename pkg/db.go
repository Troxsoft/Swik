package pkg

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/robertkrimen/otto"
)

type KeyValue struct {
	Keys   []string
	Values []string
}
type DBdata struct {
	Keys KeyValue
}
type DB struct {
	data      DBdata
	filename  string
	jsRuntime *otto.Otto
}

func (s *DB) JS() *otto.Otto {
	return s.jsRuntime
}

var (
	DatabaseAlreadyExists = errors.New("The database already exists")
	KeyAlrearyExists      = errors.New("The key/value already exists")
	KeyNotExists          = errors.New("The key/value not exists")
)

func (s *DB) ExistsKey(key string) bool {

	for _, v := range s.data.Keys.Keys {
		if v == key {
			return true
		}
	}
	return false
}
func (s *DB) GetIndexForKey(key string) (*int, error) {
	if !s.ExistsKey(key) {
		return nil, KeyNotExists
	}
	for i, v := range s.data.Keys.Keys {
		if v == key {
			return &i, nil
		}
	}
	return nil, KeyNotExists
}
func (s *DB) SetValueForKey(key, value string) error {
	if !s.ExistsKey(key) {
		return KeyNotExists
	}
	i, err := s.GetIndexForKey(key)
	if err != nil {
		return err
	}
	s.data.Keys.Values[*i] = value
	return nil
}
func (s *DB) AddKey(key string) error {
	if s.ExistsKey(key) {
		return KeyAlrearyExists
	}
	s.data.Keys.Keys = append(s.data.Keys.Keys, key)
	s.data.Keys.Values = append(s.data.Keys.Values, "")
	return nil
}
func (s *DB) Data() DBdata {
	return s.data
}
func (s *DB) RemoveKeyValue(key string) error {

	i, err := s.GetIndexForKey(key)
	if err != nil {
		return err
	}
	s.data.Keys.Keys = append(s.data.Keys.Keys[:*i], s.data.Keys.Keys[*i+1:]...)
	s.data.Keys.Values = append(s.data.Keys.Values[:*i], s.data.Keys.Values[*i+1:]...)
	return nil
}
func (s *DB) AddJSFunctions() {
	s.jsRuntime.Set("getKeys", func(call otto.FunctionCall) otto.Value {

		j, _ := json.Marshal(s.data.Keys.Keys)
		v, _ := s.jsRuntime.ToValue(string(j))
		return v
	})
	s.jsRuntime.Set("getAll", func(call otto.FunctionCall) otto.Value {
		keysValues := make(map[string]string)
		for i, v := range s.data.Keys.Keys {
			keysValues[v] = s.Data().Keys.Values[i]
		}
		j, _ := json.Marshal(keysValues)
		v, _ := s.jsRuntime.ToValue(string(j))
		return v

	})
	s.jsRuntime.Set("getValues", func(call otto.FunctionCall) otto.Value {

		j, _ := json.Marshal(s.data.Keys.Values)
		v, _ := s.jsRuntime.ToValue(string(j))
		return v
	})
	s.jsRuntime.Set("set", func(call otto.FunctionCall) otto.Value {
		key, err := call.Argument(0).ToString()
		if err != nil || call.Argument(0).IsUndefined() {
			e, _ := s.jsRuntime.ToValue(`{
"sucess":false
}`)
			return e
		}

		value, err := call.Argument(1).ToString()
		if err != nil || call.Argument(0).IsUndefined() {
			e, _ := s.jsRuntime.ToValue(`{
"sucess":false
}`)
			return e
		}
		if s.ExistsKey(key) {
			err := s.SetValueForKey(key, value)
			if err != nil {
				e, _ := s.jsRuntime.ToValue(`{
"sucess":false
}`)
				return e
			}
		} else {
			err := s.AddKey(key)
			if err != nil {
				e, _ := s.jsRuntime.ToValue(`{
"sucess":false
}`)
				return e
			}
			err = s.SetValueForKey(key, value)
			if err != nil {
				e, _ := s.jsRuntime.ToValue(`{
"sucess":false
}`)
				return e
			}
		}
		e, _ := s.jsRuntime.ToValue(`{
"sucess":true
}`)
		return e

	})

}
func CreateDB(filename string) error {

	db := &DB{
		filename: filename,
		data: DBdata{
			Keys: KeyValue{
				Keys:   []string{},
				Values: []string{},
			},
		},
		jsRuntime: otto.New(),
	}
	b, err := json.Marshal(db.data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, b, 0777)
	if err != nil {
		return err
	}
	return nil
}
func (s *DB) Save() error {
	b, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.filename, b, 0777)
	return nil
}
func NewDBFromFile(filename string) (*DB, error) {
	d, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dbContent DBdata
	err = json.Unmarshal(d, &dbContent)

	if err != nil {
		return nil, err
	}
	db := &DB{
		filename:  filename,
		data:      dbContent,
		jsRuntime: otto.New(),
	}
	return db, nil
}
