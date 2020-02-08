package utils

import (
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	Server   string
	Database string
}

var server = Server{}

var db *mgo.Database

func (m *Server) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func CreateQR(id string) string {
	qrCode, _ := qr.Encode(id, qr.H, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 250, 250)
	base := "/home/jigar/playground/DocBabu/backend"
	file, _ := os.Create(filepath.Join(base, "assets", id+".png"))
	// fmt.Println(filepath.Join("..", "assets", id+".png"))
	defer file.Close()

	png.Encode(file, qrCode)
	return file.Name()
}
