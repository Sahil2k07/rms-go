# Go Assessment SynergyLabs

### Recruitement Management System

## Tech Used

- `Golang`: Utilized Golang's standard HTTP library for routing and building a RESTful API.

- `Mysql`: Employed MySQL as the database solution, using raw SQL queries for efficient and flexible data handling.

- `Validator`: Integrated for robust input validation, ensuring data integrity and security.

- `Goose`: Managed database migrations with Goose, facilitating seamless schema version control and updates.

- `Docker`: Containerized the application with Docker, simplifying setup for the team. Testing can be done easily with `docker-compose up`.

## Set-Up Project Locally

### Docker

1. First clone the Project locally:

   ```bash
   git clone https://github.com/Sahil2k07/rms-go.git
   ```

2. Move to the Project directory:

   ```bash
   cd rms-go
   ```

3. Set these environment variables in the `.env` file.

   ```dotenv
    API_URL=
    API_KEY=
   ```

4. Run the command to start your Containerized Application

   ```bash
   docker-compose up
   ```

   or

   ```bash
   docker-compose up -d
   ```

5. If you have Docker Compose Plugin, Use this command instead

   ```bash
   docker compose up
   ```

   or

   ```bash
   docker compose up -d
   ```

6. You will be able to access this application in `localhost:3000` of your machine.

### Golang

1. Clone the repository to the local:

   ```bash
   git clone https://github.com/Sahil2k07/rms-go.git
   ```

2. Move to the project directory:

   ```bash
   cd rms-go
   ```

3. Set up all the required env variable by making a `.env` file. A `.env.example` file has been given for reference.

   ```dotenv
   PORT=localhost:3000

   JWT_SECRET=GoAppSecret

   ALLOWED_ORIGINS=*

   MYSQL_URL=root:<password>@tcp(127.0.0.1:3306)/<dbName>

   API_KEY=
   API_SECRET=
   ```

4. Run the command to download all the dependencies to your local machine:

   ```bash
   go mod vendor
   ```

5. Now you will need to apply all the migrations in your database. First install goose locally:

   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

6. Move to the migrations directory:

   ```bash
   cd src/database/migrations
   ```

7. Run the command to apply the migrations. Make sure to modify the command to have your's used `MySQL` database url in the `.env`. Make sure to have a database created before-hand:

   ```bash
   goose mysql "root:<password>@tcp(127.0.0.1:3306)/<dbName>" up
   ```

8. After applying the migrations traverse back to the root directory:

   ```bash
   cd ../../../
   ```

9. Build the Binary to start the Project:

   ```bash
   go build
   ```

10. Start the Project:

    ```bash
    ./rms-go
    ```
