CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       is_admin BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE tasks (
                       id UUID PRIMARY KEY,
                       user_id UUID NOT NULL,
                       title TEXT NOT NULL,
                       description TEXT,
                       completed BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT now(),
                       CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
