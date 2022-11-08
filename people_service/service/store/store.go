package store

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		return fmt.Errorf("migrate: %v", err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %v", err)
	}
	fmt.Println("migration is done")
	return &Store{
		conn: conn,
	}

}
func (s *Store) ListPeople() ([]People, error) {
	row, err := s.conn.Query(context.Background(), "Select name, id From people")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}
	defer row.Close()
	var peoplesList []People
	for row.Next() {
		var people People

		if err = row.Scan(&people.Name, &people.ID); err != nil {
			fmt.Fprintf(os.Stderr, "Scann failed: %v", err)
			return nil, err
		}
		peoplesList = append(peoplesList, people)

	}
	return peoplesList, err
}

func (s *Store) GetPeopleByID(id string) (People, error) {
	var name string
	err := s.conn.QueryRow(context.Background(), "Select name From people where id=%1 ", id).Scan(&name)
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not found")
	}
	return People{Name: name, ID: i}, err

}
