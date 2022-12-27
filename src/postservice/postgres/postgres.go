package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(connString string) (*gorm.DB, error) {
	// Do not use main context since some business logic closing down might still
	//  need to commit to database. Be sure to defer pool.Close in main.
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

// QueryExample is a example of how you can use sqlc to create your database layer
//func (db *DB) QueryExample() (int, error) {
//	i, err := db.conn.SampleQuery(context.Background())
//	if err != nil {
//		return int(i), errors.Wrap(err, "Failed to query SampleQuery")
//	}
//	return int(i), nil
//}
