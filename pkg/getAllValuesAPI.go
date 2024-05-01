package pkg

func (db *DB) APIGetALLValues() []*APIValue {
	j := []*APIValue{}

	for k := range db.data.Keys.KeyValueData {
		j = append(j, &APIValue{
			db.data.Keys.KeyValueData[k],
		})
	}

	return j
}
