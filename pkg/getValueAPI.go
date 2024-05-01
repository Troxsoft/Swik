package pkg

func (db *DB) APIGetValue(key string) *APIValue {
	if !db.ExistsKey(key) {
		return nil
	}
	return &APIValue{
		Value: db.data.Keys.KeyValueData[key],
	}
}
