# GSLC-NetSysProg-2

# 1. Pengertian dari method-method yang ada di http
# GET
Method GET merupakan method yang digunakan untuk meminta data dari resource tertentu. Method GET memerintah server untuk mengirimkan resource target dan server akan mengirim resource dalam respons's body. Request GET dapat dicache, dan tidak boleh ada perubahan atau menghapus data diserver
# POST
Method POST  merupakan method yang digunakan untuk mengirim data ke resource tertentu. Method POST memberi tahu server bahwa kita mengirim data untuk dikaitkan dengan target resource. Method ini dapat mengubah atau memberikan efek samping pada server contohnya membuat resource baru pada server
# HEAD
Method HEAD merupakan method yang mirip seperti GET hanya saja method ini meminta header dan bukan body dari resource. Method HEAD berguna untuk mengambil detail seperti ukuran dari resource atau untuk mengecek apakah resource mengalami perubahan dalam rentan waktu tertentu
# PUT
Method PUT merupakan method yang digunakan untuk mengupdate/menambahkan resource yang ada sekarang dengan yang baru atau bahkan mengubah keseluruhannya dengan yang baru
# PATCH
Method PATCH merupakan method yang digunakna untuk memperbarui sebagian resource yang telah ada dengan resource yang baru dikirim, sehingga resource yang sudah ada tidak termodifikasi
# DELETE
Method DELET merupakan method yang digunakan untuk menghapus dan menghilangkan resource tertentu yang sudah ada dari server
# OPTIONS
Method OPTIONS merupakan method yang digunakan untuk memberikan informasi mengenai method apa saja yang bisa yang dapat digunakan. 
# CONNECT
Method CONNECT merupakan method yang digunakan untuk merequest dibukanya koneksi antara client dan server, dengan menjalankan HTTP tunneling, atau dengan membangun TCP session
# TRACE
Methode TRACE merupakan method yang digunakan untum mengambalikan request yang dikirim daripada mengeksekusinya. Method ini digunakna untuk debug atau testing, bisa juga untuk melihat apakah request kita mengalami mengalami modifikasi dari antara node 


# 3. Cara kerja HTTP
HTTP akan menjalan dan membentuk TCP connection ketika client mengirim HTTP request
- client akan merequest ke server, request ini termusuk method yang akan digunakan seperti GET, POST, PUT, DELETE 
- Server akan menerima request dan memprosesnya
- Server akan mengembalikan respons ke client, respond dapat berupa kode status, header, atau data yang direquest
- client menerima respon dari server kemudian memproses sesuai data yang di kembalikan oleh server
