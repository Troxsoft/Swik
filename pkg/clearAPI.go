package pkg

func (db *DB) APIClear() {
	db.data.Keys.KeyValueData = make(map[string]any)

}
