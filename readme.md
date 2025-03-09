# Bulk transaction service

<img src="https://img.shields.io/badge/-blue?style=for-the-badge&logo=go&logoColor=white&logoSize=auto" alt="GO" height="28px"> <img src="https://img.shields.io/badge/postgresql-4169e1?style=for-the-badge&logo=postgresql&logoColor=white&logoSize=auto" alt="postgresql" height="28px"> <img src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff&style=for-the-badge&logoSize=auto" alt="Docker Badge">

Bulk transaction service written in ***Go*** using **PostgreSQL**.\
The goal of the project is to demonstrate the implementation of such a functionality in ***Go***.

The app uses Echo web framework, Goose migration tool, psx PostgreSQL driver, and godotenv library to setup env variables from .env file.

### Prerequisites:
1. postgreSQL instance.
The repo provides docker compose file to set up Postgres with docker.
You can also use standalone instance of Postgres.

2. go 1.23.5 installed.

3. .env file filled with proper values. Use .template.env file as example.

### How to run (using docker)

1. clone the repo and navigate to the corresponding directory.
2. run command:
```bash
docker compose up
```
>Note: you might need to run it with sudo command

3. then execute:
```bash
make run
```