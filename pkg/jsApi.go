package pkg

type APIKeyValue struct {
	Key   APIKey   `json:"key"`
	Value APIValue `json:"value"`
}
type APIKey struct {
	Key string `json:"key"`
}
type APIValue struct {
	Value any `json:"value"`
}

func (db *DB) ImplementAPI() {
	//db.JS().RunString(ApiJS)
	db.JS().Set("db", map[string]any{

		"get":          db.APIGet,
		"getAll":       db.APIGetALL,
		"getKey":       db.APIGetKey,
		"getValue":     db.APIGetValue,
		"getAllKeys":   db.APIGetALLKeys,
		"getAllValues": db.APIGetALLValues,
		"set":          db.APISet,
		"clear":        db.APIClear,
		"remove":       db.APIRemove,
		"getIf":        db.APIGetIF,
		"getAllIf":     db.APIGetALLIF,
		"removeIf":     db.APIRemoveIF,
		"removeAllIf":  db.APIRemoveAllIF,
	})
}
