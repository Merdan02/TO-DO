CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       role TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('admin', 'user')),
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now()
);


CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL,
                       title TEXT NOT NULL,
                       description TEXT,
                       done BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now(),
                       CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
