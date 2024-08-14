CREATE TYPE quest_status AS ENUM ('active', 'completed', 'failed');


CREATE TABLE quests (
    id UUID PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status quest_status NOT NULL
);
