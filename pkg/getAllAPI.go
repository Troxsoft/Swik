package pkg

func (db *DB) APIGetALL() []*APIKeyValue {
	j := []*APIKeyValue{}

	for k, v := range db.data.Keys.KeyValueData {
		j = append(j, &APIKeyValue{
			Key: APIKey{
				Key: k,
			},
			Value: APIValue{
				Value: v,
			},
		})
	}

	return j
}
