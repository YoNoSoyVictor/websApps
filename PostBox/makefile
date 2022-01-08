# This file is will quickly setup a docker postgres container with all the necessary tables, users, etc. 
# meant for local developement.

# Requires a postgres:alpine image, if you don't have one run "docker pull postgres:alpine"

# To run it first execute: runPostgres, then createUser with "password" as password, and finally createDB


runPostgres:
	docker run \
	--name DockerPostgres \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=safestpasswordeverlol \
	-p 5432:5432 \
	-d postgres:alpine

createUser:
	docker exec -it DockerPostgres createuser -P web

dropUser:
	dropuser web

createDB:
	docker exec -it DockerPostgres createdb ThoughtBox ; \
	docker exec -it DockerPostgres psql -d ThoughtBox \
	\
	\
	-c "CREATE TABLE posts (\
	id SERIAL PRIMARY KEY, \
	title VARCHAR (100) NOT NULL, \
	message VARCHAR NOT NULL, \
	created_on TIMESTAMP);" \
	\
	-c "CREATE TABLE users (\
	id SERIAL PRIMARY KEY, \
	name VARCHAR NOT NULL, \
	email VARCHAR NOT NULL, \
	password VARCHAR NOT NULL, \
	created_on TIMESTAMP, \
	active BOOLEAN);" \
	\
	\
	-c "alter table posts alter column created_on set default NOW();" \
	-c "alter table users alter column created_on set default NOW();" \
	-c "alter table users alter column active set default TRUE;" \
	\
	-c "GRANT SELECT, INSERT, UPDATE on posts TO web;" \
	-c "GRANT ALL ON SEQUENCE posts_id_seq TO web;" \
	-c "GRANT SELECT, INSERT, UPDATE on users TO web;" \
	-c "GRANT ALL ON SEQUENCE users_id_seq TO web;" 

dropDB:
	docker exec -it DockerPostgres dropdb ThoughtBox

.PHONY: runPostgres createUser dropUser createDB dropDB
.SILENT: runPostgres createUser dropUser createDB dropDB