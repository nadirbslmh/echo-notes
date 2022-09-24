## echo-notes

REST API application to manage notes.

## How to use

1. Clone this repository.

2. Create a new database called `echo_notes`.

```sql
CREATE DATABASE echo_notes;
```

3. Create a new table called `notes` in `echo_notes` database.

```sql
USE echo_notes;

CREATE TABLE notes(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

4. Run the application with this command:

```sh
go run main.go
```
