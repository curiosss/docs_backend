CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    password VARCHAR,
    role VARCHAR,
    fcm_token VARCHAR,
    created_at TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    parent_id INTEGER,
    color VARCHAR,
    FOREIGN KEY (parent_id) REFERENCES categories(id)
);

CREATE TABLE docs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    doc_name VARCHAR,
    doc_no VARCHAR,
    category_id INTEGER REFERENCES categories(id),
    end_date TIMESTAMP,
    notify_date TIMESTAMP,
    notif_sent BOOLEAN,
    status VARCHAR,
    permission INTEGER,
    created_at TIMESTAMP
);

CREATE TABLE doc_users (
    user_id INTEGER,
    doc_id INTEGER,
    permission INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (doc_id) REFERENCES docs(id)
);

CREATE TABLE actions (
    id SERIAL PRIMARY KEY,
    doc_id INTEGER REFERENCES docs(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    type VARCHAR,
    created_at TIMESTAMP
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR,
    size INTEGER,
    doc_id INTEGER REFERENCES docs(id)
);
