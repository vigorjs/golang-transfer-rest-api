# Dibuat sebagai test rekrutment Backend di PT Mitra Kasih Perkasa (MKP)
Nama   : Virgo Fajar Pamungkas

# Review Teknis Setelah Interview
  setelah interview kemarin saya mencoba untuk memperbaiki sedikit review yang saya terima dan berikut adalah beberapa updatenya :
  
* handle ketersediaan tiketing kursi 
review : untuk handle ketersediaan tiket sebenarnya sudah ada kolom isAvailable(bool) dalam tabel kursi dari desain database yang lampirkan sebelumnya. 
update : Untuk contoh implementasinya sudah saya update agar saat user booking/hit API booking otomatis mengupdate isAvailable menjadi false/tidak tersedia.

* satu user dapat pesan multi tiket
review : untuk dapat melakukan ini juga sepengalaman saya dilihat dari desain database yang saya lampirkan sebelumnya seharusnya bisa untuk di implementasikan pemesanan multi tiket oleh satu user yang mana pada desain tersebut telah terdapat relasi many-to-many dari tabel Booking dan Kursi yang dimana dari relasi tersebut menghasilkan tabel booking_kursi yang bisa digunakan sebagai mapping.
update : untuk contoh implementasinya, saya membuat API untuk Booking dimana untuk menghandle pemesanan multi tiket diatur sesuai logicnya didalam fungsi API tersebut dengan tahap2 kurang lebih seperti berikut : 
get data kursi yang tersedia -> kemudian, set request data Booking sesuai modelnya dan set untuk request dari data kursi dalam bentuk array -> validasi-> kemudian query untuk data Booking dan data kursi secara terpisah ke masing2 database. 


# REST API Bioskop Golang + Jwt Auth
Ini adalah sistem pemesanan tiket bioskop sederhana yang dibangun menggunakan bahasa pemrograman Golang, menggunakan framework Gin untuk routing HTTP dan Gorm sebagai ORM nya.

# Instalasi dan Penggunaan
1. go mod tidy.
2. cp .env.example .env
3. Konfigurasikan file .env
4. Run dengan perintah go run main.go atau menggunakan compiler daemon dengan perintah compiledaemon --command="./bioskop_golang"

# Postman Collection (updated)
1. Auth : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-9ea84973-ef2a-4549-95d8-d281da70a586?action=share&creator=29862535
2. Jadwal : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-7974ad35-d6dd-46ed-a3bb-7bf4fc6f737b?action=share&creator=29862535
3. Booking : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-6abd53e6-057f-474b-b79c-960ef08cf2a6?action=share&creator=29862535
4. Film : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-67d70f1b-c81b-4712-976c-3091c2e57fe4?action=share&creator=29862535
5. Bioskop : https://www.postman.com/research-physicist-10707010/workspace/bioskop-golang/collection/29862535-461520b7-bac6-4538-a458-08456932925e?action=share&creator=29862535

   
# Teknologi yang Digunakan
Golang
Gin
Gorm
bcrypt
godotenv
JWT (JSON Web Token)
Postgresql
daemonCompiler
