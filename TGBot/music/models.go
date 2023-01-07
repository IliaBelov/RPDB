package music

type Music struct {
	Id         int
	Music_name string
	Author     string
	Music_text string
}

type MyMusicList struct {
	User_id  int
	Music_id int
}

type Users struct {
	Id    int
	Tg_id int
	Name  string
}
