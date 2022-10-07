## echo-notes

REST API application to manage notes.

## How to use

1. Clone this repository.

2. Copy the `.env` file.

```sh
cp .env.example .env
```

3. Fill the values inside the `.env` file for the database configurations.

4. Create a new database called `echo_notes`.

```sql
CREATE DATABASE echo_notes;
```

5. Run the application with this command:

```sh
go run main.go
```

## Resources

- [Graceful Shutdown Implementation](https://medium.com/tokopedia-engineering/gracefully-shutdown-your-go-application-9e7d5c73b5ac).