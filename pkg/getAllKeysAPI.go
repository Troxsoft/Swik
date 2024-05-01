package pkg

func (db *DB) APIGetALLKeys() []*APIKey {
	j := []*APIKey{}

	for k := range db.data.Keys.KeyValueData {
		j = append(j, &APIKey{
			Key: k,
		})
	}

	return j
}
