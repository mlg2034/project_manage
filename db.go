package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySqlStore struct {
	db *sql.DB
}

func NewSqlStorage(cfg mysql.Config) *MySqlStore {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to MySQL")
	return &MySqlStore{db: db}
}

func (s *MySqlStore) Init() (*sql.DB, error) {

	//initialize tha tables
	if err := s.createProjectTable(); err != nil {
		return nil, err
	}
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}
	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil

}

func (s *MySqlStore) createProjectTable() error {

	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS project (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`)
	if err != nil {
		return err
	}
	return nil
}

func (s *MySqlStore) createTasksTable() error {

	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    name VARCHAR(255) NOT NULL,
    status ENUM('TODO','IN_PROGRESS','IN_TESTING','DONE')NOT NULL DEFAULT 'TODO',
    projectId INT UNSIGNED NOT NULL,
    assignedToId INT UNSIGNED NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (assignedToId) REFERENCES users(id),
    FOREIGN KEY (projectId) REFERENCES project(id),
    
    
)
ENGINE=InnoDB DEFAULT CHARSET=utf8;
`)
	if err != nil {
		return err
	}
	return nil

}

func (s *MySqlStore) createUsersTable() error {

	_, err := s.db.Exec(`
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    email VARCHAR(255) NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    UNIQUE KEY (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`)
	return err
}
