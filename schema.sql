DROP TABLE IF EXISTS cat;

CREATE TABLE "cat" (
   "id" int NOT NULL PRIMARY KEY,
   "name" VARCHAR(30) NOT NULL,
   "age" INT
);

INSERT INTO cat (id, name, age) VALUES (1,'cat1',1);
INSERT INTO cat (id, name, age) VALUES (2,'cat3',2);