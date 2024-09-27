package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type MySQlStorage struct {
	db *sql.DB
}

// open + validate
func NewSqlStorage(cfg mysql.Config) *MySQlStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	// just validate without open the db itself
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &MySQlStorage{
		db: db,
	}
}

func (s *MySQlStorage) init() (*sql.DB, error) {
	// create the table here
	if err := s.createProjectTable(); err != nil {
		return nil, err
	}

	if err := s.createUserTable(); err != nil {
		return nil, err
	}

	if err := s.createTastTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQlStorage) createTastTable() error {
	_, err := s.db.Exec(
		`
		CREATE TABLE IF NOT EXIST tasks (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO' , 'INPROGRESS','IN_TESTING' ,'DONE' ) NOT NULL DEFAULT 'TODO',
			projectId Int UNSIGNED NOT NULL,
			assignedToID Int UNSIGNED NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FORGIN KEY (assignedToID) REFERENCE users(id),
			FORGIN KEY (projectId) REFERENCE projects(id),
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`)
	if err != nil {
		return err
	}

	return err
}

func (s *MySQlStorage) createProjectTable() error {
	_, err := s.db.Exec(
		`
		CREATE TABLE IF NOT EXIST projects (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`)
	if err != nil {
		return err
	}

	return err
}

func (s *MySQlStorage) createUserTable() error {
	_, err := s.db.Exec(
		`
		CREATE TABLE IF NOT EXIST users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`)
	if err != nil {
		return err
	}

	return err
}
