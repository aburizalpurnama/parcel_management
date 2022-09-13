# PARCEL MANAGEMENT SYSTEM
### Aplikasi management pengambilan paket di apartement 

## 1. Create Database Schema
- Menggunakan dbdiagram.io (sintaks menggunakan DBML)
- trus export ke bahasa database yng diinginkan

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

        migrate create -ext sql -dir db/migration -seq create_user_table


## NOTE

### ~> Jika terdapat session masih aktif saat drop database.

1. Masuk ke terminal postgres container
    
	    docker exec -it postgres14 psql -U postgres

2. Jalankan Query untuk kill semua session

	    SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '<db-name>' AND pid <> pg_backend_pid();

3. Kemudian Keluar dari terminal postgres container
	
        \q

### ~> Make file Error :
    
    Error : makefile:28: *** target pattern contains no '%'.  Stop.
    atau
    Error : Makefile: *** missing separator. Stop

1. Cek pada tiap baris makefile, pastikan tidak ada whitespace (spasi), hanya boleh ada tab.