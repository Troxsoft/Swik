package pkg

import (
	"errors"
	"os"

	"github.com/bytedance/sonic"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

type KeyValue struct {
	KeyValueData map[string]any
}
type DBdata struct {
	Keys KeyValue
}
type DB struct {
	data      DBdata
	filename  string
	jsRuntime *goja.Runtime
	registry  *require.Registry
}

func (s *DB) JS() *goja.Runtime {
	return s.jsRuntime
}

var (
	DatabaseAlreadyExists = errors.New("The database already exists")
	KeyAlrearyExists      = errors.New("The key/value already exists")
	KeyNotExists          = errors.New("The key/value not exists")
)

func (s *DB) ExistsKey(key string) bool {

	_, ok := s.data.Keys.KeyValueData[key]

	return ok
}

func (s *DB) Data() DBdata {
	return s.data
}

func CreateDB(filename string) error {
	runt := goja.New()
	runt.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	regis := new(require.Registry)
	regis.Enable(runt)
	console.Enable(runt)

	db := &DB{
		filename: filename,
		data: DBdata{
			Keys: KeyValue{
				KeyValueData: make(map[string]any),
			},
		},
		jsRuntime: runt,
		registry:  regis,
	}
	b, err := sonic.Marshal(db.data)
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
	b, err := sonic.Marshal(s.data)
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
	err = sonic.Unmarshal(d, &dbContent)

	if err != nil {
		return nil, err
	}
	runt := goja.New()
	runt.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	regis := new(require.Registry)

	regis.Enable(runt)
	console.Enable(runt)

	db := &DB{
		filename:  filename,
		data:      dbContent,
		jsRuntime: runt,
		registry:  regis,
	}
	return db, nil
}
