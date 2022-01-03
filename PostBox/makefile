# This file is will quickly setup a docker postgres container, meant for local developement.

# Use postgres alpine
runPostgres:
	docker run --name DockerPostgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=safestpasswordeverlol -p 5432:5432 -d postgres:alpine

createUser:
	docker exec -it DockerPostgres createuser -P web

dropUser:
	dropuser web

createDB:
	docker exec -it DockerPostgres createdb ThoughtBox ; docker exec -it DockerPostgres psql -d ThoughtBox -c "CREATE TABLE posts (id SERIAL PRIMARY KEY, title VARCHAR (100) NOT NULL, message VARCHAR (1000) NOT NULL, created_on TIMESTAMP);" -c "alter table posts alter column created_on set default NOW();" -c "GRANT SELECT, INSERT, UPDATE on posts TO web;" -c "GRANT ALL ON SEQUENCE posts_id_seq TO web;"

dropDB:
	docker exec -it DockerPostgres dropdb ThoughtBox
