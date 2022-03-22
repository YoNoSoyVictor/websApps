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
	docker exec -it DockerPostgres createdb PostBox ; \
	docker exec -it DockerPostgres psql -d PostBox \
	\
	\
	-c "CREATE TABLE posts (\
	id SERIAL PRIMARY KEY, \
	title VARCHAR (255) NOT NULL, \
	message VARCHAR NOT NULL, \
	created_on TIMESTAMP DEFAULT NOW());" \
	\
	-c "CREATE TABLE users (\
	id SERIAL PRIMARY KEY, \
	name VARCHAR NOT NULL, \
	email VARCHAR UNIQUE NOT NULL, \
	password CHAR(60) NOT NULL, \
	created_on TIMESTAMP DEFAULT NOW(), \
	active BOOLEAN DEFAULT TRUE);" \
	\
	\
	-c "GRANT SELECT, INSERT, UPDATE on posts TO web;" \
	-c "GRANT ALL ON SEQUENCE posts_id_seq TO web;" \
	-c "GRANT SELECT, INSERT, UPDATE on users TO web;" \
	-c "GRANT ALL ON SEQUENCE users_id_seq TO web;" 

dropDB:
	docker exec -it DockerPostgres dropdb PostBox

.PHONY: runPostgres createUser dropUser createDB dropDB
.SILENT: runPostgres createUser dropUser createDB dropDB