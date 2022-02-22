package apiHandler

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "log"
    "strconv"
    "gorm.io/gorm"
    "strings"
    "time"
)

import (
    "dcard-backend-hw/model"
)

func returnErr(w http.ResponseWriter, status int) {
    w.WriteHeader(status)
    w.Write([]byte(http.StatusText(status)))
}

func UploadURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
        returnErr(w, http.StatusUnsupportedMediaType)
        w.Write([]byte("\nYou should use Json and set your content-type into application/json"))
		return
	}

    reqBody, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Println(err)
        returnErr(w, http.StatusBadRequest)
        return
    }

    type Data struct {
        Url string `json:url`
        ExpireAt string `json:expireAt`
    }
    var data Data

    err = json.Unmarshal([]byte(reqBody), &data)
    if err != nil {
        log.Println(err)
        returnErr(w, http.StatusBadRequest)
        return
    }

    if (!strings.HasPrefix(data.Url, "http://") && !strings.HasPrefix(data.Url, "https://")) {
        returnErr(w, http.StatusBadRequest)
        return
    }

    exp, err := time.Parse(time.RFC3339, data.ExpireAt)
    if err != nil {
        log.Println(err)
        returnErr(w, http.StatusBadRequest)
        return
    }

    url := model.Url{OriginUrl: data.Url, ExpireAt: data.ExpireAt}
    db.Create(&url)

    response := make(map[string]string)
    response["id"] = strconv.Itoa(int(url.ID))
    response["shotUrl"] = "http://localhost/" + response["id"]
    jsonRes, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
        return
    }

    w.Header().Set("content-type", "application/json")
    w.Write(jsonRes)
}
