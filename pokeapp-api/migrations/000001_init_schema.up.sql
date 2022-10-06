DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR(32) NOT NULL CHECK (username <> ''),
  email VARCHAR(64) UNIQUE NOT NULL CHECK (email <> ''),
  password VARCHAR(250)  NOT NULL CHECK (octet_length(password) <> 0), 
  role VARCHAR(10) NOT NULL DEFAULT 'user',
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
  comment_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
  pokemon_id iNT NOT NULL,
  message VARCHAR(1024) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
