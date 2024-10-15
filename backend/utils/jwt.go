package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtSecret []byte
var encryptionKey []byte

// Again...
func init() {
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	jwtSecret = []byte(viper.GetString("JWT_SECRET"))
    // encryptionKey = make([]byte, 32)
    // if _, err := rand.Read(encryptionKey); err != nil {
    //     panic(err)
    // }
	encryptionKey = []byte("12345678901234561234567890123456")
}

func GenerateToken(userId, name, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"name":   name,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // 2 hours timeout
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	encryptedToken, err := encryptToken(signedToken)
	if err != nil {
		return "", err
	}

	return encryptedToken, nil
}

func GetUserIDFromToken(encryptedToken string) (string, error) {
	decryptedToken, err := decryptToken(encryptedToken)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(decryptedToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	}
	return "", err
}

func ValidateToken(encryptedToken string) (jwt.MapClaims, error) {
	decryptedToken, err := decryptToken(encryptedToken)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(decryptedToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func encryptToken(token string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypting token
	encryptedToken := gcm.Seal(nonce, nonce, []byte(token), nil)
	return base64.URLEncoding.EncodeToString(encryptedToken), nil
}

func decryptToken(encryptedToken string) (string, error) {
	encryptedBytes, err := base64.URLEncoding.DecodeString(encryptedToken)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return "", errors.New("malformed encrypted token")
	}

	// Extracting nonce and ciphertext
	nonce, ciphertext := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]
	decryptedToken, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedToken), nil
}

var InvalidTokens = make(map[string]struct{})

func InvalidateToken(token string) {
	InvalidTokens[token] = struct{}{}
}
