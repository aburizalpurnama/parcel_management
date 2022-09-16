# PARCEL MANAGEMENT SYSTEM
### Aplikasi management pengambilan paket di apartement 

## 1. Create Database Schema
- Menggunakan dbdiagram.io (sintaks menggunakan DBML)
- trus export ke bahasa database yng diinginkan

    ### Implement UUID in postgress

        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

        CREATE TABLE "roles" (
        "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        "name" varchar NOT NULL,
        );

        CREATE TABLE "users" (
        "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        "name" varchar NOT NULL,
        "role_id" uuid,
        );

## 2. Create Makefile

- ### Create & Drop Database
- ### Docker Compose Up & Down
- ### Golang Migrate Up & Down

## 3. Create Migrate
- ### Install golang migrate (kalo belum)
- ### Buat migrate

        migrate create -ext <file extention> -dir <migration file destination directory> -seq <schema>

    example:

        migrate create -ext sql -dir db/migration -seq init_schema

        atau 

        migrate create -ext sql -dir db/migration -seq create_user_table

## 4. Implement PSQL
- ### Instal PSQL (using Docker)
        docker pull kjconroy/sqlc

- ### Buat SQL Query di folder query untuk setiap model dengan nama nama_model.sql
        -- name: CreateUnit :one
        INSERT INTO units (
            no, email, phone
        ) VALUES (
            $1, $2, $3
        )
        RETURNING *;

- ### Generate PSQL (using Docker)
        docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate


## 5. CRUD Testing
- ### Create Random Function
    Buat otomatis dapetin data yng tidak mungkin sama, sehingga tidak error saat insert field dengan unique constraint.


- ### Buat Main Test File (main_test.go)
        const (
        dbDriver = "postgres"
        dbSource = "postgresql://<username>:<password>@localhost:5432/<db-name>?sslmode=disable"
        )

        var testQueries *Queries

        func TestMain(m *testing.M) {
        conn, err := sql.Open(dbDriver, dbSource)
        if err != nil {
                log.Fatal("cannot connect to db:", err)
        }

        testQueries = New(conn)

        os.Exit(m.Run())
        }

    *Kalo ada Error, Lihat di Note Section

- ### Buat Unit Test File Untuk Tiap Model atau Entity(model_test.go)
    
    - Buat function untuk Creation Operation
        (Create Account, Create User, Create Unit, etc)

## NOTE

### ~> Jika terdapat session masih aktif saat drop database.

1. Masuk ke terminal postgres container
    
	    docker exec -it postgres14 psql -U postgres

2. Jalankan Query untuk kill semua session

	    SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'parcel_management' AND pid <> pg_backend_pid();

3. Kemudian Keluar dari terminal postgres container
	
        \q

### ~> Make file Error :
    
    Error : makefile:28: *** target pattern contains no '%'.  Stop.

    atau

    Error : Makefile: *** missing separator. Stop

1. Cek pada tiap baris makefile, pastikan tidak ada whitespace (spasi), hanya boleh ada tab.

### ~> cannot connect to db:sql: unknown driver "postgres" (forgotten import?)

1. Download Dependency sql driver

        go get github.com/lib/pq

2. Import dan ignore it

        import (
            _ "github.com/lib/pq"
        )

### ~> Can't Generate Random while running rapid test in vscode
It because the vscode chache the value, so it don't ge the other random value anymore.
 Solution of this issue is to set the golang test flag with -count=1. so it will not catch the value.

 1. Go to File > Preference > Setting > search "go.testFlag" > open setting.json file, and add following code.

        "go.testFlags": [
        "-count=1"
        ]