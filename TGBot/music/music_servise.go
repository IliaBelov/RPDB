package music

import (
	"fmt"
)

func (s *Store) AddMusic(m *Music) error {
	fmt.Printf(m.Music_name)
	/*if s.CheckNameMusic(m) {

	}*/
	err := s.CreateMusic(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DelMusic(m *Music) error {
	err := s.DeleteMusic(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) AddList(m *Music, user_id int64) error {
	err := s.AddMusicList(m, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DelList(m *Music, user_id int64) error {
	err := s.DeleteMusicList(m, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CheckList(u *Users) ([]Music, error) {
	var list []Music
	list, err := s.CheckMusicList(u)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (s *Store) AddUser(u *Users) error {
	err := s.AddUserTG(u)
	if err != nil {
		return nil
	}
	return err
}
