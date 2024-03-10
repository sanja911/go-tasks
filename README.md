# Fitur Utama:

- [x] Pendaftaran dan Otentikasi Pengguna:
Pengguna dapat mendaftar dan membuat akun.
Otentikasi menggunakan JWT (JSON Web Token).

- [ ] Manajemen Tugas:
Pengguna dapat membuat tugas baru dengan judul, deskripsi, dan tenggat waktu.
Tugas dapat ditandai sebagai selesai atau belum selesai.
Pengguna dapat mengedit atau menghapus tugas.

- [ ] Kategori Tugas:
Tugas dapat dikategorikan ke dalam berbagai kategori (Pekerjaan, Pendidikan, Pribadi, dll.).
Pengguna dapat membuat kategori baru dan mengelola tugas berdasarkan kategori.

- [ ] Pengingat dan Pemberitahuan:
Pengguna dapat mengatur pengingat untuk tugas tertentu.
Sistem pemberitahuan untuk memberi tahu pengguna tentang tugas yang akan segera berakhir.

- [ ] Antarmuka Pengguna Web:
Dashboard yang intuitif untuk melihat tugas yang sedang berlangsung.
Halaman manajemen tugas untuk membuat, mengedit, dan menghapus tugas.
Statistik sederhana untuk melihat progres dan produktivitas.

# Pengujian Unit dan Integrasi:
Menyertakan unit testing dan integrasi testing untuk memastikan keandalan aplikasi.

# Teknologi:

- Menggunakan Golang sebagai bahasa pemrograman utama.
- Database menggunakan PostgreSQL atau MongoDB.
- Framework web seperti Gin atau Echo.
- JWT untuk otentikasi.
- HTML/CSS dan JavaScript untuk antarmuka pengguna.

# Alur Proses GoTask:

**1. Pengguna Mendaftar dan Masuk:**
  Pengguna baru mendaftar dengan memberikan nama pengguna, alamat email, dan kata sandi.
Setelah mendaftar, pengguna dapat masuk menggunakan kredensial mereka.

**2. Dashboard Pengguna:**
  Setelah masuk, pengguna diarahkan ke dashboard yang menampilkan daftar tugas yang sedang berlangsung.
Pengguna dapat melihat tugas-tugas yang belum selesai, tenggat waktu, dan kategori.

**3. Manajemen Tugas:**
  Pengguna dapat membuat tugas/tasks baru dengan menentukan judul, deskripsi, dan tenggat waktu.
Tugas-tugas yang ada ditampilkan dalam daftar, dengan opsi untuk menandai sebagai selesai atau menghapusnya.

**4. Kategori Tugas:**
  Pengguna dapat membuat kategori baru untuk mengelompokkan tugas sesuai kebutuhan (pekerjaan, pendidikan, pribadi, dll.).
Tugas-tugas dapat dikelompokkan berdasarkan kategori untuk organisasi yang lebih baik.

**5. Pengaturan dan Preferensi Pengguna:**
  Pengguna dapat mengatur preferensi seperti pengaturan notifikasi atau tampilan dashboard.
Kemungkinan pengaturan profil seperti mengubah kata sandi atau informasi akun.

**6. Pengingat dan Pemberitahuan:**
  Sistem pengingat mengirimkan pemberitahuan kepada pengguna untuk tugas yang akan segera berakhir.
Pengguna menerima pemberitahuan tentang tugas yang telah diselesaikan atau telah diperbarui.
