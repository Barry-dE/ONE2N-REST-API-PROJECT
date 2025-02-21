CREATE EXTENSION IF NOT EXISTS citext;


CREATE TABLE IF NOT EXISTS students (id BIGSERIAL PRIMARY KEY,
                                                          firstName VARCHAR(255) NOT NULL,
                                                                                 lastName VARCHAR(255) NOT NULL,
                                                                                                       email CITEXT UNIQUE NOT NULL,
                                                                                                                           age INT NOT NULL,
                                                                                                                                   sex VARCHAR(10) NOT NULL,
                                                                                                                                                   created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
                                                                                                                                                                                                           updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW())