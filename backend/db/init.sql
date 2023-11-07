CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    isAdmin BOOLEAN DEFAULT FALSE
);

CREATE TABLE Documents (
    id SERIAL PRIMARY KEY ,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    owner_id INTEGER NOT NULL,
    content JSON,
    link TEXT,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);

CREATE TABLE Logs (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    document_id INTEGER NOT NULL,
    prev_content JSON,
    new_content JSON,
    edited_by INTEGER NOT NULL,
    FOREIGN KEY (document_id) REFERENCES Documents(id)
);
