## Pengenalan Parallel Programming
- Saat ini kita hidup di era multicore, di mana jarang sekali kita pakai processor single core
- Maka dari itu terbitlah proses parallel di aplikasi 
- **Parallel Programming** : sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi lebih kecil-kecil, dan dijalankan bersamaan pada waktu yang bersamaan pula 

#### Contoh Parallel Programming
- Menjalankan aplikasi sekaligus (multi-tasking)
- Beberapa koki menyiapkan makanan di restoran 
- Antrian di bank, tiap teller melayani nasabah masing-masing 

## Process Vs Thread 
| **Process**      | **Thread** |
| ----------- | ----------- |
| Process adalah sebuah eksekusi program      | Thread adalah segmen dari process (bagian dari process)      |
| Process menggunakan memori besar   | Thread menggunakan memory kecil        |
| Process saling terisolasi dengan process lain   | Thread bisa saling berhubungan jika dalam process yang sama      |
| Process lama untuk dijalankan dan dihentikan   | Thread cepat untuk dijalankan dan dihentikan        |

- Analogi : Process (Google Chrome) | Thread (Tab Google Chrome)

## Parallel Vs Concurrency
| **Parallel**      | **Concurrency** |
| ----------- | ----------- |
| Menjalankan beberapa pekerjaan secara **bersamaan**      | Menjalankan beberapa pekerjaan secara **bergantian**      |
| Butuh banyak thread   | Butuh sedikit thread        |

### Contoh Concurrency
- Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, ngobrol lagi, dan berulang-ulang. Proses ini tidak bisa dilakukan secara bersamaan oleh seorang manusia (hanya bergantian dari satu hal ke hal yang lain)

## CPU-bound 
- Banyak algoritma dibuat yang hanya menggunakan CPU untuk menjalankannya. Algoritmaa jenis ini biasanya sangat tergantung kecepatan CPU.
- Contoh yang paling populer adalah ML (Machine Learning), oleh karena itu sekarang banyak ML yang pakai GPU karena simply core-nya lebih banyak dibanding CPU
- Jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implement Paraller Programming

## I/O-bound
- I/O-bound adalah kebalikan dari sebelumnya, di mana biasanya algoritma atau aplikasi sangat bergantung dengan kecepatan IO device yang dipakai
- Contohnya aplikasi seperti membaca file, database, dll.
- Kebanyakan aplikasi saat ini menggunakan I/O-bound
- Aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika pakai Concurrency Programming
- Bayangkan kita read dari database, dan Thread harus menunggu 1 detik untuk mendapat response, padahal waktu 1 detik jika menggunakan Concurrency Programming bisa digunakan untuk bergantian melakukan hal lain 


## Pengenalan Goroutine 
- Goroutine adalah sebuah **mini-thread** yang dikelola oleh Go Routine 
- Ukuran Goroutine sangat kecil, sekitar 2KB, jauh lebih kecil dibanding Thread yang bisa sampai 1MB,
- Namun tidak seperti thread yang berjalan paraller, Goroutine berjalan secara **concurrent**

### Cara kerja Goroutine 
- Goroutine dijalankan oleg Go Scheduler dalam Thread, di mana jumlah Threadnya sebanyak GOMAXPROCS (sejumlah core CPU)
- Goroutine berjalan di atas Thread, jadi tidak bisa dibilang Goroutine pengganti Thread 
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, sudah diatur oleh Go Scheduler

**Dalam Go-Scheduler, Terminologi umum :**
- **G** : Goroutine
- **M** : Thread (Machine)
- **P** : Processor 

### Cara Kerja Go-Scheduler (Tidak perlu manage secara manual, sudah otomatis)
- Thread akan mengambil process Goroutine yang ada di dalam queue(baik local maupun global)
- Thread akan mengambil process dari local queue-nya terlebih dahulu, setelah itu ke global queue, setelah itu ke local queue dari processor (core) lain
- **Apabila** ada proses Goroutine yang sangat panjang dan lama untuk dieksekusi maka Go-Scheduler akan secara bergantian (concurrent) mengembalikan kembali Goroutine tersebut ke dalam queue dan mengambil kembali untuk dieksekusi hingga selesai


### Pengenalan Channel 
- Channel adalah tempat komunikasi secara syncrhonous yang bisa dilakukan oleh Goroutine 
- Di channel terdapat pengirim dan penerima, biasanya mereka dari Goroutine yang berbeda
- Saat pengiriman data ke Channel, Goroutine akan ter-block, sampai ada yang menerima data tersebut 
- Maka dari itu, Channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternatif mekanisme async-await (non-blocking)

#### Mekanisme Channel 
- Membuat sebuah Channel (tempat mengirim data)
- Goroutine mengirim data ke dalam Channel, dan akan ditahan di channel sampai ada yang mengambil (pengambil adalah Goroutine yang lain)
- Goroutine mengambil data dari dalam Channel, apabila data belum ada maka akan ditunggu hingga data ada

#### Karakteristik Channel 
- Secara default Channel hanya bisa menampung satu data, jika ingin tambah data, harus menunggu data di Channel diambil dulu 
- Channel hanya bisa menerima satu jenis data 
- Channel bisa diambil dari lebih dari satu Goroutine 
- Channel harus di-close jika tidak dipakai, atau bisa menyebabkan **memory leak**




