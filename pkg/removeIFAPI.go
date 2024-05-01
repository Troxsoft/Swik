package pkg

func (db *DB) APIRemoveIF(f func(APIKeyValue) bool) bool {
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
			delete(db.data.Keys.KeyValueData, k)
			return true
		}
	}

	return false
}
