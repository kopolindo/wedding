<img alt="gitleaks badge" src="https://img.shields.io/badge/protected%20by-gitleaks-blue">

Static webpage to keep track of guest invitations

# Setup
## Requirements
- [go](https://go.dev/doc/install)
- [docker](https://docs.docker.com/engine/install/)
- [docker compose](https://docs.docker.com/compose/install/)

## First run

1. Create the docker instance for MariaDB.  
In the root of the repository there is file `01_create_tables_template.sql`.  
Copy/move this file to `wedding/mariadb/data/init-db/` (if folder does not exist create it).

> Caveat! The compose file will create a mounted volume in `wedding/mariadb/data/` folder.
Every time you need to start from scratch delete the volumes or that folder content.

2. Create a CSV file named `guests.csv` holding the guests list in the root of the repository (`wedding/guests.csv`)

```.csv
Mario,Rossi
John,Doe
....
```

3. Create password files in the root of the repository
    - `wedding/password_db_root.txt`
    - `wedding/password_db.txt`
    - `wedding/cookie_passphrase.txt`

4. Create the wordlist, named `wedding/passphrase-generator-dictionary.txt`, for passphrase generator in the root of the repository

5. Run the DB
`docker compose -f "compose.yml" up -d --build`

# Roadmap

- [x] logging
- [x] more pictures of us
- [x] IBAN
- [ ] monitoring
- [ ] reverse proxy