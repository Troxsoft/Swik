package pkg

func (db *DB) APIGet(key string) *APIKeyValue {
	if !db.ExistsKey(key) {
		return nil
	}
	return &APIKeyValue{
		Key: APIKey{
			Key: key,
		},
		Value: APIValue{
			Value: db.data.Keys.KeyValueData[key],
		},
	}
}
