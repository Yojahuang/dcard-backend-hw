package redirectHandler

import (
    "log"
    "net/http"
    "strconv"
    "time"
)

import (
    "github.com/gorilla/mux"
)

import (
    "gorm.io/gorm"
)

import (
    "dcard-backend-hw/model"
)

func hasExpired(url model.Url) bool {
    exp, err := time.Parse(time.RFC3339, url.ExpireAt)
    if err != nil {
        log.Fatal(err)
        return false
    }

    now := time.Now()
    if now.After(exp) {
        return true
    }

    return false
}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    tmp, _ := strconv.Atoi(vars["id"])
    Id := uint(tmp)

    var url model.Url
    result := db.Find(&url, Id)

    if result.RowsAffected == 0 || hasExpired(url) {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(http.StatusText(http.StatusNotFound)))
        if result.RowsAffected != 0 {
            db.Delete(&url)
        }
        return
    }

    w.Header().Add("Location", url.OriginUrl)
    w.WriteHeader(http.StatusSeeOther)
}
