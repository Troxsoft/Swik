package pkg

func (db *DB) APIRemoveAllIF(f func(APIKeyValue) bool) int {
	num := 0
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
			num += 1
		}
	}
	return num

}
