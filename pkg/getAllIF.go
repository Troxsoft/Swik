package pkg

func (db *DB) APIGetALLIF(f func(APIKeyValue) bool) []*APIKeyValue {
	j := []*APIKeyValue{}

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

			j = append(j, &vol)
		}
	}

	return j
}
