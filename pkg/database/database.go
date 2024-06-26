package database

import (
	"fmt"

	"github.com/converged-computing/rainbow/pkg/utils"
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"log"
	"os"
)

type Database struct {
	filepath string
}

// Database types to serialize back into
type Cluster struct {
	Name   string
	Secret string
	Token  string
}

// cleanup removes the filepath
func (db *Database) cleanup() {
	// Delete a previous database that exists
	// Note that in the future we might not want to do this
	log.Printf("🧹️ cleaning up %s...", db.filepath)
	os.Remove(db.filepath)
}

// create the database
func (db *Database) create() error {

	log.Printf("✨️ creating %s...", db.filepath)

	// Create SQLite file (ensures that we can!)
	file, err := os.Create(db.filepath)
	if err != nil {
		return err
	}
	file.Close()
	log.Printf("   %s file created", db.filepath)

	// Open the created SQLite File (to test)
	conn, err := db.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	return err
}

// Connect to the database - the caller is responsible for closing
func (db *Database) connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", db.filepath)
	if err != nil {
		return nil, err
	}
	return conn, err
}

// countResults counts the results for a specific query
func countResults(conn *sql.DB, query string) (int64, error) {

	var count int

	err := conn.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

// GetCluster gets a cluster by name
func (db *Database) GetCluster(name string) (*Cluster, error) {

	// Connect!
	conn, err := db.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// First determine if it exists
	query := fmt.Sprintf("SELECT * from clusters WHERE name LIKE \"%s\" LIMIT 1", name)
	statement, err := conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Only allow one result, one cluster
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	// Unwrap result into cluster
	cluster := Cluster{}
	for rows.Next() {
		err := rows.Scan(&cluster.Name, &cluster.Token, &cluster.Secret)
		if err != nil {
			return nil, err
		}
	}
	// Debugging if needed
	// log.Printf("%s: %s\n", query, cluster.Name)
	return &cluster, nil
}

// ValidateClusterToken checks if a cluster token is valid
// The token is used for validating a submission request.
func (db *Database) ValidateClusterToken(name, token string) (*Cluster, error) {
	cluster, err := db.GetCluster(name)
	if err != nil {
		return nil, err
	}

	// Validate the name and token
	if cluster.Name == "" || cluster.Token != token {
		return nil, fmt.Errorf("request denied")
	}
	return cluster, nil
}

// ValidateClusterSecret checks if a cluster secret is valid
// The secret is used for validating a request for jobs.
func (db *Database) ValidateClusterSecret(name, secret string) (*Cluster, error) {
	cluster, err := db.GetCluster(name)
	if err != nil {
		return nil, err
	}

	// Validate the name and token
	if cluster.Name == "" || cluster.Secret != secret {
		return nil, fmt.Errorf("request denied")
	}
	return cluster, nil
}

// create the database
func (db *Database) createTables() error {

	conn, err := db.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// TODO none of these have logic for what to do on delete
	// Create the clusters table, where we store the name and secret
	// obviously the secret should not be stored in plain text - it's fine for now
	createClusterTableSQL := `
	CREATE TABLE clusters (
		name TEXT NOT NULL PRIMARY KEY,
		token TEXT,
		secret TEXT
	  );
	`
	createTimingTableSQL := `
	CREATE TABLE timinginfo (
		name TEXT NOT NULL PRIMARY KEY,
		time TEXT
	  );

	// A job has a unique id and is assigned a cluster
	createJobsTableSQL := `
	  CREATE TABLE jobs (
		  idJob integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		  cluster TEXT,
		  name TEXT,
		  jobspec string,
		  FOREIGN KEY(cluster) REFERENCES clusters(name)
		);`

	// Create single query for tables
	createSQL := createClusterTableSQL + "\n" + createJobsTableSQL
	log.Println("   🏓️ creating tables...")

	// Execute SQL query
	_, err = conn.Exec(createSQL)
	if err != nil {
		return err
	}
	log.Println("   🏓️ tables created")
	return nil
}

func InitDatabase(filepath string, cleanup bool) (*Database, error) {

	// Create a new database (todo, add cleanupc check)
	db := Database{filepath: filepath}

	if cleanup {
		db.cleanup()
	}

	// If we haven't created yet or cleaned up
	exists, err := utils.PathExists(db.filepath)
	if err != nil {
		return nil, err
	}
	if !exists {
		// Create the database
		err := db.create()
		if err != nil {
			return nil, err
		}
		err = db.createTables()
		if err != nil {
			return nil, err
		}
	}
	return &db, err
}
