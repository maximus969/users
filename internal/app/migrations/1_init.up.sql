CREATE TABLE users (
  id UUID PRIMARY KEY not null, 
  firstname text not null, 
  lastname text not null, 
  email text not null, 
  age integer not null CHECK (age > 0),
  created timestamp with time zone DEFAULT now() NOT NULL
);