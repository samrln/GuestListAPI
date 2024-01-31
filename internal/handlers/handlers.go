package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"

    "guestListAPI/internal/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Guests)
}

func GetGuest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range models.Guests {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Guest{})
}

func GetConfirmed(w http.ResponseWriter, r *http.Request) {
    var confirmed []models.Guest
    for _, item := range models.Guests {
        if item.Confirmed {
            confirmed = append(confirmed, item)
        }
    }
    json.NewEncoder(w).Encode(confirmed)
}

func GetNotConfirmed(w http.ResponseWriter, r *http.Request) {
    var notConfirmed []models.Guest
    for _, item := range models.Guests {
        if !item.Confirmed {
            notConfirmed = append(notConfirmed, item)
        }
    }
    json.NewEncoder(w).Encode(notConfirmed)
}

func AddGuest(w http.ResponseWriter, r *http.Request) {
    var guest models.Guest
    _ = json.NewDecoder(r.Body).Decode(&guest)
    models.Guests = append(models.Guests, guest)
    json.NewEncoder(w).Encode(guest)
}

func UpdateConfirmed(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range models.Guests {
        if item.Name == params["name"] {
            models.Guests[index].Confirmed = true
            json.NewEncoder(w).Encode(models.Guests[index])
            return
        }
    }
    json.NewEncoder(w).Encode(models.Guest{})
}
