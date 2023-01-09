package music

import (
	"fmt"
	"log"
)

func (s *Store) AddMusic(m *Music) error {
	fmt.Printf(m.Music_name)
	err := s.createMusic(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DelMusic(m *Music) error {
	err := s.deleteMusic(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) AddList(m *Music, u *Users) string {

	us, err := s.checUserName(u)
	if err != nil {
		return "Пользователь не найден"
	}
	u.Id = us[0].Id
	ms, err := s.checkIDMusicByName(m)
	if err != nil {
		return "Ошибка проверки музыки"
	}
	if len(ms) == 0 {
		return "Нет такой музыки"
	}
	m.Id = ms[0].Id
	if err := s.addMusicList(m, u); err != nil {
		return "Ошибка добавления музыки"
	}
	return "Музыка добавлена"
}
func (s *Store) AddUser(us *Users) error {
	log.Println(us.Name)
	mass, err := s.checUserName(us)
	if err != nil {
		return err
	}
	if len(mass) == 0 {
		err = s.addUserDB(us)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) DelList(m *Music, u *Users) string {

	us, err := s.checUserName(u)
	if err != nil {
		return "Пользователь не найден"
	}
	u.Id = us[0].Id
	ms, err := s.checkIDMusicByName(m)
	if err != nil {
		return "Ошибка проверки музыки"
	}
	if len(ms) == 0 {
		return "Нет такой музыки"
	}
	m.Id = ms[0].Id
	msl, err := s.checkMusicInList(u, m)
	if err != nil {
		return "Ошибка проверки музыки в playlist"
	}
	if len(msl) == 0 {
		return "Нет такой музыки"
	}
	err = s.deleteMusicList(m, u)
	if err != nil {
		return err.Error()
	}
	return "Удаление успешно"
}

func (s *Store) CheckList(u *Users) string {
	var list string
	knm, err := s.checUserName(u)
	if err != nil {
		return "Ошибочка"
	}
	log.Println(knm[0].Id)
	uid, err := s.checkMusicList(&knm[0])
	if err != nil {
		return "Playlist пуст, милорд"
	}
	for _, v := range uid {
		mass, err := s.checkNameMusicById(v.Music_id)
		if err != nil {
			return "Неизвестная ошибка"
		}
		list += fmt.Sprintf("%s - %s\n", mass[0].Music_name, mass[0].Author)
	}

	return list
}
func (s *Store) CheckText(m *Music) string {
	msc, err := s.checkMusicText(m)
	if err != nil {
		return "Неизвестная ошибка"
	}
	if len(msc) == 0 {
		return "Нет такой музыки"
	}
	return msc[0].Music_text
}

func (s *Store) SearchMusic(m *Music) string {
	msc, err := s.searchMusic(m)
	if err != nil {
		return "Неизвестная ошибка"
	}
	if len(msc) == 0 {
		return "Нет такой музыки"
	}
	return fmt.Sprintf("Название: %s\nАвтор: %s\nТекст: %s", msc[0].Music_name, msc[0].Author, msc[0].Music_text)

}

func (s *Store) SearchAuthor(m *Music) string {
	var list string
	msc, err := s.searchAuthor(m)
	if err != nil {
		return "Неизвестная ошибка"
	}
	if len(msc) == 0 {
		return "Нет музыки такого автора"
	}
	for _, v := range msc {

		list += fmt.Sprintf("Название: %s  Автор: %s\nТекст: %s\n", v.Music_name, v.Author, v.Music_text)
	}
	return list

}
