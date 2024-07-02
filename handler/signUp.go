package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

type Env struct {
	Driver *sql.DB
}

func (d *Env) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	newUUID, err1 := exec.Command("uuidgen").Output()
	if err1 != nil {
		response := models.Response{Message: err1.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	encText, _ := Encrypt(user.PasswordHash, MySecret)

	_, err = d.Driver.Exec(`INSERT INTO 
	user(uuid, Name,Email,Address,UserType,PasswordHash,ProfileHeadline) 
	VALUES (?,?,?,?,?,?,?)`,
		newUUID,
		user.Name,
		user.Email,
		user.Address,
		user.IsAdmin,
		encText,
		user.ProfileHeadline)

	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}

	tokenString, tokenError := createToken(string(newUUID))
	if tokenError != nil {
		fmt.Println("error", tokenError)
	}

	fmt.Println("Token sent", tokenString)
	response := models.Response{Message: "Created new user", Status: "Success", Jwt: tokenString}
	SendResponse(w, response)
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecretKey string) (string, error) {
	fmt.Println(MySecretKey)
	block, err := aes.NewCipher([]byte(MySecretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}
