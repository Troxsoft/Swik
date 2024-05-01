package pkg

func (db *DB) APIRemove(key string) bool {
	if !db.ExistsKey(key) {
		return false
	}
	delete(db.data.Keys.KeyValueData, key)
	return true
}
