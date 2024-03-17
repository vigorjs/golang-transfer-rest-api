*Dibuat sebagai test rekrutment Backend di PT Mitra Kasih Perkasa (MKP)
Nama   : Virgo Fajar Pamungkas

-REST API Bioskop Golang + Jwt Auth
Ini adalah sistem pemesanan tiket bioskop sederhana yang dibangun menggunakan bahasa pemrograman Golang, menggunakan framework Gin untuk routing HTTP dan Gorm sebagai ORM nya.

-Instalasi dan Penggunaan
1. go mod tidy.
2. cp .env.example .env
3. Konfigurasikan file .env
4. Run dengan perintah go run main.go atau menggunakan compiler daemon dengan perintah compiledaemon --command="./bioskop_golang"

-Postman Collection
1. Auth : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-9ea84973-ef2a-4549-95d8-d281da70a586?action=share&creator=29862535
2. Jadwal : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-7974ad35-d6dd-46ed-a3bb-7bf4fc6f737b?action=share&creator=29862535
3. Film : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-67d70f1b-c81b-4712-976c-3091c2e57fe4?action=share&creator=29862535
4. Bioskop : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-461520b7-bac6-4538-a458-08456932925e?action=share&creator=29862535

   
-Teknologi yang Digunakan
Golang
Gin
Gorm
bcrypt
godotenv
JWT (JSON Web Token)
Postgresql
daemonCompiler
