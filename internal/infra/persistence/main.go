/* Package persistence implement entity repositories with gorm */
package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open open a connection to MySQL database and return a DB object with the
// connection
func Open(dsn string) (*DB, error) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &DB{conn}, err
}

// DB wraps gorm.DB so other package don't need to know anything
// about it (gorm.DB)
type DB struct {
	conn *gorm.DB
}

// Close disconnect from database
func (d *DB) Close() error {
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
