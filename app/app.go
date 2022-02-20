package app

import (
    "log"
    "net/http"
)

import (
    "github.com/gorilla/mux"
)

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

import (
    "dcard-backend-hw/model"
    "dcard-backend-hw/handler"
)

type requestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func requestFuncWrapper(db *gorm.DB, handler requestHandlerFunction) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        handler(db, w, r)
    }
}

func Init() {
    log.Println("Connecting to database.db...")
    db, err := gorm.Open(sqlite.Open("databases/database.db"), &gorm.Config{})
    log.Println("Done")

    log.Println("Migrating database.db...")
    db.AutoMigrate(&model.Url{})

    if err != nil {
        log.Fatal(err)
    }
    log.Println("Done")

    router := mux.NewRouter()

    router.HandleFunc("/api/v1/urls", requestFuncWrapper(db, handler.UploadURL)).Methods("POST")

    log.Println("API is running on port 4000!")
    log.Fatal(http.ListenAndServe(":4000", router))
}
