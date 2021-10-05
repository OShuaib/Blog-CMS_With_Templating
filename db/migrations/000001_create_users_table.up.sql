CREATE TABLE IF NOT EXISTS users (
    id VARCHAR (50) NOT NULL PRIMARY KEY,
    email VARCHAR (50) NOT NULL UNIQUE ,
    firstname VARCHAR (50) NOT NULL ,
    lastname VARCHAR (50) NOT NULL ,
    password VARCHAR (100) NOT NULL ,
    created_at INT,
    updated_at INT
    );
