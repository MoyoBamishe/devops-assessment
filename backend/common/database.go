package common

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres" // Added for Postgres support
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// GetDBPath returns the database path from environment or default.
func GetDBPath() string {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/gorm.db"
	}
	return dbPath
}

// GetTestDBPath returns the test database path from environment or default.
func GetTestDBPath() string {
	testDBPath := os.Getenv("TEST_DB_PATH")
	if testDBPath == "" {
		testDBPath = "./data/gorm_test.db"
	}
	return testDBPath
}

// ensureDir creates the directory for the database file if it doesn't exist
func ensureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir != "" && dir != "." {
		return os.MkdirAll(dir, 0750)
	}
	return nil
}

// Init opens a database and saves the reference to the `Database` struct.
func Init() *gorm.DB {
	var db *gorm.DB
	var err error
	dbURL := os.Getenv("DATABASE_URL")

	// DEBUG: Verify exactly what the app sees as the connection string
	fmt.Println("DEBUG: The DSN value is:", dbURL)

	if dbURL != "" {
		// Try connecting to Postgres if DATABASE_URL is provided
		fmt.Println("Connecting to Postgres via DATABASE_URL...")
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		// Fallback to original SQLite logic
		dbPath := GetDBPath()
		if err := ensureDir(dbPath); err != nil {
			fmt.Println("db err: (Init - create dir) ", err)
		}
		fmt.Println("Connecting to SQLite...")
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	}

	if err != nil {
		fmt.Println("db err: (Init) ", err)
		// If it errors here, return early to prevent nil pointer crashes later
		return nil 
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("db err: (Init - get sql.DB) ", err)
	} else {
		sqlDB.SetMaxIdleConns(10)
	}
	DB = db
	return DB
}

// TestDBInit creates a temporary database for running testing cases
func TestDBInit() *gorm.DB {
	testDBPath := GetTestDBPath()
	if err := ensureDir(testDBPath); err != nil {
		fmt.Println("db err: (TestDBInit - create dir) ", err)
	}

	test_db, err := gorm.Open(sqlite.Open(testDBPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}
	sqlDB, err := test_db.DB()
	if err != nil {
		fmt.Println("db err: (TestDBInit - get sql.DB) ", err)
	} else {
		sqlDB.SetMaxIdleConns(3)
	}
	DB = test_db
	return DB
}

// TestDBFree deletes the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	sqlDB, err := test_db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	testDBPath := GetTestDBPath()
	err = os.Remove(testDBPath)
	return err
}

// GetDB is used to get a connection pool.
func GetDB() *gorm.DB {
	return DB
}
