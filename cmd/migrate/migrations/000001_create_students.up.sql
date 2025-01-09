CREATE EXTENSION IF NOT EXISTS citext;


CREATE TABLE IF NOT EXISTS students (id BIGSERIAL PRIMARY KEY,
                                                          first_name VARCHAR(255) NOT NULL,
                                                                                  last_name VARCHAR(255) NOT NULL,
                                                                                                         email CITEXT UNIQUE NOT NULL,
                                                                                                                             age INT NOT NULL,
                                                                                                                                     sex VARCHAR(10) NOT NULL,
                                                                                                                                                     created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW())