package store

import (
	"context"

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

	// make migration

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	s.conn.Query(context.Background(), "Select name, id From people")
	return []People{}, nil
}

func (s *Store) GetPeopleByID(id string) (People, error) {
	s.conn.QueryRow(context.Background(), "Select name From people where id=%1 ", id).Scan(&People{Name: id})
	return People{}, nil
}
