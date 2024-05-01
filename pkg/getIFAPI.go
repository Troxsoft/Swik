package pkg

func (db *DB) APIGetIF(f func(APIKeyValue) bool) *APIKeyValue {
	for k, v := range db.data.Keys.KeyValueData {
		vol := APIKeyValue{
			Key: APIKey{
				Key: k,
			},
			Value: APIValue{
				Value: v,
			},
		}
		if f(vol) {
			return &vol
		}
	}

	return nil
}
