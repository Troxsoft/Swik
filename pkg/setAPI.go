package pkg

func (db *DB) APISet(key string, value any) {
	db.data.Keys.KeyValueData[key] = value
}
