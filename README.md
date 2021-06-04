# Belajar Docker untuk Pemula - Membuat TODO App

TODO app ini adalah contoh app untuk mendemokan proses membuat aplikasi dengan Docker, terdiri dari:
- Frontend dengan [Vue JS framework](https://docs.vuejs.id/v2/guide/)
- Backend dengan [Golang](https://dasarpemrogramangolang.novalagung.com/)
- Database dengan Postgres

Semua komponen dipackage dengan docker

## Menjalankan dengan Docker

Step 1: Jalankan Postgres
```bash
docker run -p 5432:5432 --name todo-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=rahasia -e POSTGRES_DB=belajar -v $(pwd)/postgres/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
```

Step 2: Export konfigurasi database sebagai environment variabel
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=rahasia
export DB_DATABASE=belajar
```

Step 3: Jalankan Backend Golang app di lokal komputer
```bash
go run backend/main.go
```

Step 4: Jalankan Frontend JS app di lokal komputer
```bash
cd frontend
yarn install
yarn serve
```

Step 5: Buka browser untuk mulai mengakses app TODO
```bash
http://localhost:8081
```

Step 6: Cek input TODO tersimpan di database
```bash
docker exec -it todo-postgres psql -U postgres -W belajar
SELECT * FROM todo;
```
