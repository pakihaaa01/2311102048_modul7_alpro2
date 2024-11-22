# 2311102048_modul7_alpro2
laporan praktikum algoritma pemrograman 2 modul 7
1. deskripsi program No.1
Struct Titik dan Lingkaran
Struct Titik merepresentasikan sebuah titik pada bidang 2D dengan koordinat x dan y.
Struct Lingkaran merepresentasikan lingkaran, terdiri dari tengah (titik pusat lingkaran) dan jari (jari-jari lingkaran).
Fungsi kuadrat
Fungsi ini menghitung kuadrat dari bilangan n, digunakan untuk menghindari penghitungan akar kuadrat saat membandingkan jarak.
Fungsi beradaDiDalamLingkaran
Fungsi ini menentukan apakah sebuah titik berada di dalam atau di tepi lingkaran.
Perbandingan dilakukan menggunakan rumus jarak kuadrat dari titik ke pusat lingkaran terhadap kuadrat jari-jari.
Fungsi main
Membaca input untuk dua lingkaran (l1 dan l2) dan satu titik (t).
Memeriksa apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau tidak sama sekali.
Memberikan output sesuai kondisi.
2. deskripsi program No.2
Input Array:
Program meminta pengguna untuk memasukkan jumlah elemen array (n) dan nilai setiap elemen array satu per satu.
Array dibuat menggunakan fungsi make untuk menyesuaikan ukuran dinamis.
Menampilkan Isi Array:
Program menampilkan semua elemen array yang telah dimasukkan oleh pengguna.
Menampilkan Elemen Berdasarkan Indeks:
Indeks Ganjil: Menampilkan elemen array dengan indeks 1-based ganjil.
Indeks Genap: Menampilkan elemen array dengan indeks 1-based genap.
Menampilkan Elemen pada Indeks Kelipatan Tertentu:
Program meminta pengguna memasukkan nilai x, lalu menampilkan elemen array pada indeks yang merupakan kelipatan x.
Menghapus Elemen pada Indeks Tertentu:
Program meminta pengguna memasukkan indeks tertentu untuk dihapus dari array.
Jika indeks valid, elemen pada indeks tersebut dihapus, dan array diperbarui.
Menghitung Rata-Rata:
Program menghitung rata-rata dari elemen array menggunakan rumus:
Rata-rata = Total nilai elemen/Jumlah elemen 
Menghitung Standar Deviasi:
Program menghitung varians terlebih dahulu, kemudian mencari akar kuadratnya untuk mendapatkan standar deviasi menggunakan rumus:
Standar Deviasi = √∑(elemen−Rata-rata)^2/Jumlah elemen
Menghitung Frekuensi Elemen Tertentu:
Program meminta pengguna memasukkan nilai tertentu, lalu menghitung berapa kali nilai tersebut muncul dalam array.
3. deskripsi program No.3
Input Nama Klub:
Pengguna diminta untuk memasukkan nama kedua klub yang akan bertanding (klubA dan klubB).
Input Skor Pertandingan:
Program meminta pengguna memasukkan skor dari kedua klub dalam format: skorA skorB.
Jika skorA < 0 atau skorB < 0, program akan berhenti meminta input dan keluar dari loop.
Menentukan Pemenang:
Untuk setiap pertandingan, pemenang ditentukan berdasarkan perbandingan skor:
Jika skorA > skorB, maka klubA dimenangkan.
Jika skorA < skorB, maka klubB dimenangkan.
Jika skorA == skorB, hasilnya adalah Draw (imbang).
Pemenang dari setiap pertandingan dicatat dalam slice pemenang.
Menampilkan Hasil Pertandingan:
Setelah input skor selesai (salah satu skor negatif), program menampilkan hasil semua pertandingan, termasuk pemenang setiap pertandingan.
Akhir Program:
Program menampilkan pesan bahwa pertandingan selesai.
4. deskripsi program No.4
Input Kata:
Program meminta pengguna untuk memasukkan sebuah kata.
Kata disimpan sebagai array tabel dengan kapasitas maksimum NMAX = 127.
Panjang kata disimpan dalam variabel n.
Membalik Kata:
Program membalik urutan karakter dalam array menggunakan fungsi balikanArray.
Hasil kata yang sudah dibalik ditampilkan di konsol.
Memeriksa Palindrom:
Program memeriksa apakah kata tersebut adalah palindrom.
Palindrom adalah kata yang sama jika dibaca dari depan maupun belakang, misalnya "radar" atau "level".
Hasil pengecekan ditampilkan sebagai true jika palindrom, dan false jika tidak.
