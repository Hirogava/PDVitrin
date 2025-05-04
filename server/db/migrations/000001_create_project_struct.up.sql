CREATE TABLE projects (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description text NOT NULL,
    purpose text NOT NULL,
    relevance text NOT NULL,
    result text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE thematic (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description text
);

CREATE TABLE project_thematics (
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    thematic_id INTEGER REFERENCES thematics(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, thematic_id)
);

CREATE TABLE specializations (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description text
);

CREATE TABLE project_specializations (
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    specialization_id INTEGER REFERENCES specializations(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, specialization_id)
);