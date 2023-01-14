package model

type Todos struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=5,max=50"`
	Description string `json:"description" validate:"required,min=5"`
	Status      string `json:"status"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

/*
Struct Todos merupakan sebuah struktur data yang digunakan untuk menyimpan data TodoList, yang terdiri dari empat field yaitu:

Id: field untuk menyimpan ID dari TodoList
Title: field untuk menyimpan judul dari TodoList
Description: field untuk menyimpan deskripsi dari TodoList
Status: field untuk menyimpan status dari TodoList

Struct Response merupakan sebuah struktur data yang digunakan untuk menyimpan respon dari sebuah request, yang terdiri dari tiga field yaitu:

Status: field untuk menyimpan status dari respon tersebut
Message: field untuk menyimpan pesan dari respon tersebut
Data: field untuk menyimpan data dari respon tersebut, yang merupakan slice dari struct Todos
*/
