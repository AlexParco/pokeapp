DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS pokefavs CASCADE;


CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(32) NOT NULL CHECK (username <> ''),
  password VARCHAR(250)  NOT NULL CHECK (octet_length(password) <> 0), 
  role VARCHAR(10) NOT NULL DEFAULT 'ROLE_USER'
);

CREATE TABLE comments (
  comment_id SERIAL PRIMARY KEY ,
  body TEXT NOT NULL,
  user_id INT NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
  pokemon_id INT NOT NULL
);

CREATE TABLE pokefavs(
  pokefav_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
  pokemon_id INT NOT NULL
);
