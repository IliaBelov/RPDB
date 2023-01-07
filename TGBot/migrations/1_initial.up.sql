BEGIN;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  tg_id INT,
  name VARCHAR(255)
);

CREATE TABLE music   /*Музыка*/
(
    id SERIAL PRIMARY KEY,
    music_name VARCHAR(255) not null,
    author VARCHAR(255),
    music_text VARCHAR(2000)
);

CREATE TABLE mymusiclist    /*Список моей музыки*/
(
    user_id INT references users (id),
    music_id INT references music (id)
);

COMMIT;