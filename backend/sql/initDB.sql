
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    parent_category_id INT REFERENCES categories(id) NULL,
    name VARCHAR(30) NOT NULL
);

CREATE TABLE sessions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    start TIMESTAMPTZ,
    "end" TIMESTAMPTZ,
    category_id INT REFERENCES categories(id),
    focus SMALLINT,
	delta INTERVAL GENERATED ALWAYS AS ("end" - start ) STORED,
    category_line INT[] GENERATED ALWAYS AS (category_line(category_id)) STORED,
	CHECK (focus >= 1 AND focus <= 5),
    CHECK ("end" > start - '10 minutes'::interval AND "end" < NOW())
);


