# cara menjalankan program

- untuk menjalankan api ini pertama clone terlebih dahulu.

- kemudian masuk ke directory nomer 6 dan jalankan go mod tidy untuk menginstall package  yang disediakan

- setelah package terinstall kemudian ganti enviroment dengan mysql anda

- jalankan program dengan go run main.go


# testing lewat postman
- url add
    ADD -> localhost:5000/api/add-produk
    Rubah method menjadi POST kemudian masukan raw json dengan format :
    ```go
    {
        "Kode": masukan uint,
        "name": masukan string,
        "qty" : masukan int,
    }
    ```

- url update
    Update -> localhost:5000/api/update-produk/{id}
    rubah method menjadi PATCH kemudian masukan raw json dengan format :
    ```go
    {
        "qty": masukan int
    }
    ```

- url delete 
    delete -> localhost:5000/api/delete-produk/{id}

    untuk menghapus database tinggal masukan saja url diatas dengan method DELETE dan ganti id yang ingin dihapus.
    contoh -> localhost:5000/api/delete-produk/1


- url filter by name
    filter name -> localhost:5000/api/filter/{name}
    ganti name dengan nama product yang ingin dicari dan rubah method menjadi GET
    contoh -> localhost:5000/api/filter/lalala

- url filter by qty
    filter qty -> localhost:5000/api/filter/{qty}
    ganti qty dengan qty product yang ingin dicari dan rubah method menjadi GET
    contoh -> localhost:5000/api/filter/1100