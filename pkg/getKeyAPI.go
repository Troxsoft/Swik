package pkg

func (db *DB) APIGetKey(key string) *APIKey {
	if !db.ExistsKey(key) {
		return nil
	}
	return &APIKey{
		Key: key,
	}
}
