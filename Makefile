.PHONY: adminer migrate

DB_FILE=data/telegram.sqlite
DB_ENCRYPTED_FILE=data/telegram.sqlite.gpg
MIGRATIONS_DIR=${PWD}/migrations
DB_DIR=${PWD}/data

encrypt-database:
	gpg --quiet --batch --yes --symmetric --cipher-algo AES256 --passphrase=${LARGE_SECRET_PASSPHRASE} ${DB_FILE}
	rm data/telegram.sqlite

decrypt-database:
	gpg --quiet --batch --yes --decrypt --passphrase=${LARGE_SECRET_PASSPHRASE} --output ${DB_FILE} ${DB_ENCRYPTED_FILE}

migrate:
	docker run --rm -it \
		-v ${MIGRATIONS_DIR}:/migrations \
		-v ${DB_DIR}:/data \
		ghcr.io/mbtamuli/didactic-train/migrate \
		-source file:///migrations \
		-database sqlite3://${DB_FILE} \
		up

migrate-down:
	docker run --rm -it \
		-v ${MIGRATIONS_DIR}:/migrations \
		-v ${DB_DIR}:/data \
		ghcr.io/mbtamuli/didactic-train/migrate \
		-source file:///migrations \
		-database sqlite3://${DB_FILE} \
		down

cli:
	go build -o telegram-users-cli cmd/telegram-users-cli/main.go