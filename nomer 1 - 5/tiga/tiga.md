# soal

```go
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80
```


# Jawaban

ada kesalahan pada LISTEN 80

```go
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
EXPOSE 80
```

Perintah "FROM golang" menentukan bahwa Docker image akan dibangun dari image golang sebagai base image.
Perintah "ADD . /go/src/github.com/telkomdev/indihome/backend" menambahkan file dari lokal ke direktori backend dalam image.
Perintah "WORKDIR /go/src/github.com/telkomdev/indihome" mengubah direktori kerja saat membangun Docker image.
Perintah "RUN go get github.com/tools/godep" menginstall package godep.
Perintah "RUN godep restore" menjalankan perintah godep untuk memulihkan dependensi aplikasi.
Perintah "RUN go install github.com/telkomdev/indihome" menginstall aplikasi.
Perintah "ENTRYPOINT /go/bin/indihome" menentukan bahwa perintah default saat container dijalankan adalah menjalankan aplikasi indihome.
Perintah "LISTEN 80" tidak terdefinisi dalam Dockerfile, dan harus diganti dengan EXPOSE.
Dalam konteks ini, Docker image yang dibangun akan menjalankan aplikasi Go pada port 80.