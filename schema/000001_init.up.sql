CREATE TABLE users (
  id UUID PRIMARY KEY not null, 
  firstname varchar(255) not null, 
  lastname varchar(255) not null, 
  email varchar(255) not null, 
  age int, 
  created timestamp with time zone
)
