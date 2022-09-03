# Endpoint Design

> File ini akan dijadikan acuan dalam membuat API.

## Auth
- [x] POST `/register`

    :green_circle: **Membuat akun baru.**

    Expected result :
    - [x] Data user baru terekam di database
    - [x] Username bersifat unik
    - [x] Format email harus benar
    - [x] Email bersifat unik
    - [x] Panjang password 6 karakter
    - [x] Data: username, email, password, user_photo(tidak wajib)

- [x] POST `/login`

    :green_circle: **Masuk ke akun.**

    Expected result :
    - [x] Status login berubah menjadi `true`
    - [x] Error jika username tidak ditemukan
    - [x] Error jika username dan password tidak sama
    - [x] Error jika user login dengan username yang sudah status loginnya bernilai `true`
    - [x] Data: username, password

- [x] POST `/logout`

    :green_circle: **Keluar dari akun.**

    Expected result :
    - [x] Status login berubah menjadi `false`
    - [x] Error jika logout menggunakan username yang belum login
    - [x] Data: token

## Quote
- [x] GET `/quote`
    
    :green_circle: **Menampilkan random quote.**

    Expected result :
    - [x] Bisa menampilkan quote secara acak
    - [x] Bersifat publik

- [ ] GET `/quote/:id`
    
    :green_circle: **Menampilkan quote berdasarkan id.**

    Expected result :
    - [ ] Bisa menampilkan quote berdasarkan id
    - [ ] Hanya login user yang bisa mengakses endpoint

- [ ] GET `/quote/list`
    
    :green_circle: **Menampilkan list quote.**

    Expected result :
    - [ ] Bisa menampilkan list quote
    - [ ] Hanya login user yang bisa mengakses endpoint
    - [ ] Bisa filter berdasarkan tag

- [ ] GET `/quote/list/:username`
    
    :green_circle: **Menampilkan list quote berdasarkan username.**

    Expected result :
    - [ ] Bisa menampilkan list quote berdasarkan username
    - [ ] Hanya login user yang bisa mengakses endpoint

- [x] POST `/quote`

    :green_circle: **Membuat quote.**
    
    Expected result :
    - [x] Data quote terekam di database
    - [x] Data quote_count bertambah 1
    - [x] Hanya login user yang bisa membuat quote
    - [x] Panjang quote maksimal 100 character
    - [x] Tidak bisa hanya mengandung whitespace, tanda baca, atau angka
    - [x] Data : author, quote, tags

- [ ] PUT `/quote/:id`

    :green_circle: **Edit quote berdasarkan id.**
    
    Expected result :
    - [ ] Hanya user pembuat quote yang bisa mengedit quote terkait
    - [ ] Panjang quote maksimal 100 character
    - [ ] Tidak bisa hanya mengandung whitespace, tanda baca, atau angka
    - [ ] Hanya bisa edit author, quote, dan tags

- [ ] DELETE `/quote/:id`

    :green_circle: **Menghapus quote berdasarkan id.**
    
    Expected result :
    - [ ] Hanya user pembuat quote dan `super_user` yang bisa menghapus quote terkait

## Tag
- [ ] GET `/tags/list`

     :green_circle: **Menampilkan list tag.**

     Expected result :
     - [ ] Bisa menampilkan list tag
     - [ ] Bisa diurutkan berdasarkan `tag_count` (DSC)

# User
- [ ] GET `/profile/:id`

    :green_circle: **Menampilkan profile berdasarkan id.**

    Expected result :
    - [ ] Bisa menampilkan data user berdasarkan id
    - [ ] Hanya login user yang bisa mengakses endpoint

- [ ] GET `profile/list`

    :green_circle: **Menampilkan list profile.**

    Expected result :
    - [ ] Bisa menampilkan list data user
    - [ ] Hanya login user yang bisa mengakses endpoint

- [ ] PATCH `profile/:id`

    :green_circle: **Edit profile berdasarkan id.**

    Expected result :
    - [ ] Bisa edit data user berdasarkan id
    - [ ] Hanya login user yang bisa mengakses endpoint
    - [ ] Hanya user terkait yang bisa mengedit profile
    - [ ] Hanya bisa edit nama, username, user_photo
    - [ ] Edit username dengan validasi unik

- [ ] PATCH `profile/:id/password`

    :green_circle: **Ganti password.**

    Expected result :
    - [ ] Bisa edit password user berdasarkan id
    - [ ] Hanya login user yang bisa mengakses endpoint
    - [ ] Hanya user terkait yang bisa mengedit profile
    - [ ] Panjang password minimal 6 karakter
    - [ ] Data : password baru

## Rating
- [ ] GET `/leaderboard`

    :green_circle: **Menampilkan top 10 user yang bisa diurutkan berdasarkan `username`, `quote_count`, `user_rating`.**

    Expected result :
    - [ ] Menampilkan list top user
    - [ ] Bisa diurutkan berdasarkan `username` (ACS dan DSC)
    - [ ] Bisa diurutkan berdasarkan `quote_count` (ACS dan DSC)
    - [ ] Bisa diurutkan berdasarkan `user_rating` (ACS dan DSC)

- [ ] POST `/upvote`

    :green_circle: **Upvote quote.**

    Expected result :
    - [ ] Bisa memberikan upvote pada quote
    - [ ] Nilai upvote bertambah 1
    - [ ] Nilai user_rating bertambah 1
    - [ ] Hanya login user yang bisa memberikan upvote
    - [ ] Data : quote_id

- [ ] POST `/downvote`

    :green_circle: **Downvote quote.**

    Expected result :
    - [ ] Bisa memberikan downvote pada quote
    - [ ] Nilai downvote bertambah 1
    - [ ] Nilai user_rating bertambah -1
    - [ ] Hanya login user yang bisa memberikan downvote
    - [ ] Data : quote_id

## Model
| Auth | Users | Quotes | Ratings | Tags |
| ---- | ---- | ---- | ---- | ---- |
| auth_id | user_id | quote_id | rating_id | tag_id |
| user_id | username | user_id | user_id | tag_name |
| username | email | quote | quote_count | tag_count |
| token | password | author | user_rating |
| islogin | token | tags | | |
| | user_photo | upvote | | |
| | created_at | downvote | | |
| | updated_at | created_at | | |
| | |updated_at | | |