create table queries (
id int NOT NULL GENERATED ALWAYS AS IDENTITY,
queries varchar,
CONSTRAINT requests_pkey PRIMARY KEY (id)
)
-- создать таблицу