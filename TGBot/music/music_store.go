package music

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	conn *sqlx.DB
}

func NewPosthresDB(cfg string) (*Store, error) {

	db, err := sqlx.Open("postgres", cfg)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	//migration
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		return nil, err
	}
	fmt.Println("migration is done")

	return &Store{conn: db}, err
}

func (s *Store) CreateMusic(m *Music) error {
	/*fmt.Printf(m.Music_name)
	fmt.Printf(m.Author)
	fmt.Printf(m.Music_text)*/
	_, err := s.conn.ExecContext(context.Background(), `
	INSERT INTO music(music_name,author,music_text)
	VALUES($1,$2,$3);
	`, m.Music_name, m.Author, m.Music_text)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteMusic(m *Music) error {
	_, err := s.conn.ExecContext(context.Background(), `
	DELETE FROM music 
	WHERE music_name = $1;
	`, m.Music_name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) AddMusicList(m *Music, user_id int64) error {
	_, err := s.conn.ExecContext(context.Background(), `
	INSERT INTO mymusiclist(user_id,music_id)
	VALUES($1,$2);
	`, user_id, m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteMusicList(m *Music, user_id int64) error {
	_, err := s.conn.ExecContext(context.Background(), `
	Delete FROM mymusiclist
	WHERE user_id = $1 AND music_id = $2);
	`, user_id, m.Id)
	if err != nil {
		return err
	}
	return nil
}

// ????
func (s *Store) CheckMusicList(u *Users) ([]Music, error) {
	var musiclist []Music
	err := s.conn.SelectContext(context.Background(), &musiclist, `
	Select music_id
	FROM mymusiclist
	WHERE tg_id = $1;
	`, u.Tg_id)
	if err != nil {
		return nil, fmt.Errorf("query fail:%w", err)
	}
	return musiclist, nil
}

func (s *Store) CheckMusicText(m *Music) (string, error) {

	id, err := s.ChecMusicIdByName(m)
	if err != nil {
		return "", fmt.Errorf("query fail:%w", err)
	}

	var str string

	err1 := s.conn.SelectContext(context.Background(), &str, `
	Select music_text
	FROM music
	WHERE id = $1;
	`, id)
	if err1 != nil {
		return "", fmt.Errorf("query fail:%w", err1)
	}
	return str, nil
}

func (s *Store) ChecMusicIdByName(m *Music) (int, error) {
	err := s.conn.SelectContext(context.Background(), &m.Id, `
	Select id
	FROM music
	WHERE music_name = $1;
	`, m.Music_name)
	if err != nil {
		return -1, fmt.Errorf("query fail:%w", err)
	}
	return m.Id, nil
}

func (s *Store) AddUserTG(u *Users) error {
	_, err := s.conn.ExecContext(context.Background(), `
	INSERT INTO users(user_name,tg_id)
	VALUES($1,$2);
	`, u.Name, u.Tg_id)
	if err != nil {
		return err
	}
	return nil
}

// //??????????????
func (s *Store) ChecUserName(u *Users) error {
	err := s.conn.SelectContext(context.Background(), &u.Tg_id, `
	Select id
	FROM users
	WHERE tg_id = $1;
	`, u.Tg_id)
	if err != nil {
		return err
	}
	return nil
}

// ???????????
func (s *Store) CheckNameMusic(m *Music) bool {
	err := s.conn.SelectContext(context.Background(), &m.Music_name, `
	Select id
	FROM music
	WHERE music_name = $1;
	`, m.Music_name)
	log.Printf(err.Error())
	if err != nil {
		return true
	}
	return false
}
