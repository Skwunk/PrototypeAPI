# Go API Server for prototypeapi

Testing viability of OpenAPI2.0 for new URY API

### Overview
This server was generated by the [swagger-codegen]
(https://github.com/swagger-api/swagger-codegen) project.  

- API version: 1.0.0
- Build date: 2019-03-14T10:58:25.430Z

### Setup
To run the server, follow these simple steps:
First install golang and postgres. It is recommended that you use postman for testing api calls.

Change the config.toml.example to config.toml, change the values

Run this SQL snippet in postgres
```
CREATE USER web WITH
  LOGIN
  NOSUPERUSER
  INHERIT
  NOCREATEDB
  NOCREATEROLE
  NOREPLICATION;

CREATE DATABASE pqgotest
    WITH 
    OWNER = web
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United Kingdom.1252'
    LC_CTYPE = 'English_United Kingdom.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE SEQUENCE public.quotes_quote_id_seq;

ALTER SEQUENCE public.quotes_quote_id_seq
    OWNER TO web;

CREATE TABLE public.quotes
(
    quote_id bigint NOT NULL DEFAULT nextval('quotes_quote_id_seq'::regclass),
    text text COLLATE pg_catalog."default" NOT NULL,
    source integer NOT NULL,
    date timestamp with time zone NOT NULL DEFAULT now(),
    suspended boolean NOT NULL DEFAULT false,
    CONSTRAINT quote_id PRIMARY KEY (quote_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.quotes
    OWNER to web;
```

Run the server from the terminal
```
go run main.go
```

