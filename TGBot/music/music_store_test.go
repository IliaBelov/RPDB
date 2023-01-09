package music

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreateTestDatabase() (testcontainers.Container, *sqlx.DB) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_PASSWORD": "pass",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		panic(err)
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		panic(err)
	}
	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		panic(err)
	}

	connString := fmt.Sprintf("postgres://postgres:pass@%v:%v/test?sslmode=disable", host, port.Port())
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Up()
	return dbContainer, db
}

func TestCreateMusic(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	err := connection.createMusic(&m)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteMusic(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Music_name: "Кукла колдуна"}
	err := connection.deleteMusic(&m)
	if err != nil {
		t.Error(err)
	}
}
func TestAddMusicList(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Music_name: "Кукла колдуна", Author: "КИШ", Music_text: "К настоящему колдуну"}
	u := Users{Name: "Terepor"}

	connection.addUserDB(&u)
	connection.AddMusic(&m)
	mid, err := connection.checkIDMusicByName(&m)
	uid, err := connection.checkUserByUsername(&u)

	err = connection.addMusicList(&mid[0], &uid)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteMusicList(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	u := Users{Name: "Terepor"}

	connection.addUserDB(&u)
	connection.AddMusic(&m)
	mid, err := connection.checkIDMusicByName(&m)
	uid, err := connection.checkUserByUsername(&u)

	err = connection.addMusicList(&mid[0], &uid)
	err = connection.deleteMusicList(&mid[0], &uid)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckMusicList(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	u := Users{Name: "Terepor"}
	connection.addUserDB(&u)
	uid, err := connection.checkUserByUsername(&u)

	_, err = connection.checkMusicList(&uid)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckMusicText(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	connection.AddMusic(&m)
	_, err := connection.checkMusicText(&m)
	if err != nil {
		t.Error(err)
	}
}

func TestChecUserName(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	u := Users{Name: "Terepor"}
	connection.addUserDB(&u)

	_, err := connection.checUserName(&u)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckIDMusicByName(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	connection.AddMusic(&m)

	_, err := connection.checkIDMusicByName(&m)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckUserByUsername(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	//m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	u := Users{Name: "Terepor"}
	connection.addUserDB(&u)

	_, err := connection.checkUserByUsername(&u)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckMusicInList(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	u := Users{Name: "Terepor"}

	connection.addUserDB(&u)
	connection.AddMusic(&m)

	_, err := connection.checkMusicInList(&u, &m)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckNameMusicById(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	connection.createMusic(&m)

	_, err := connection.checkNameMusicById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchMusic(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	connection.createMusic(&m)
	m1 := Music{Music_name: "Кукла колдуна"}
	_, err := connection.searchMusic(&m1)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchAuthor(t *testing.T) {
	container, conn := CreateTestDatabase()
	defer container.Terminate(context.Background())
	connection := Store{
		conn: conn,
	}
	m := Music{Author: "КИШ", Music_name: "Кукла колдуна", Music_text: "К настоящему колдуну"}
	connection.createMusic(&m)
	m1 := Music{Author: "КИШ"}
	_, err := connection.searchAuthor(&m1)
	if err != nil {
		t.Error(err)
	}
}
