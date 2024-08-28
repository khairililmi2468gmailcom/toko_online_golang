package controllers

import (
    "net/http"
    "fmt"
)

func Home(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w,"Selamat datang di page toko")
}