package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SUPER_SECRET_KEY_COMPANY_XYZ_2026")

type contextKey string
const UserIDKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Ambil header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ErrorResponse(w, http.StatusUnauthorized,"butuh token autentikasi")
			return
		}

		// 2. Format harus: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ErrorResponse(w, http.StatusUnauthorized,"format token harus Bearer <token>")
			return
		}

		tokenString := parts[1]

		// 3. Parse dan Validasi Token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing-nya adalah HMAC (HS256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak terduga: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ErrorResponse(w, http.StatusUnauthorized,"token tidak valid atau sudah kedaluwarsa")
			return
		}

		// 4. Ambil claims (data) dari dalam token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ErrorResponse(w, http.StatusUnauthorized,"gagal membaca data token")
			return
		}

		// Ambil user_id dari claims
		userID, ok := claims["user_id"].(float64) // JWT otomatis membaca angka sebagai float64
		if !ok {
			ErrorResponse(w, http.StatusUnauthorized,"data user tidak ditemukan di token")
			return
		}

		// 5. Masukkan user_id ke dalam Context agar bisa dipakai di Handler/Service jika butuh
		ctx := context.WithValue(r.Context(), UserIDKey, int(userID))
		
		// 6. Lolos verifikasi! Lanjutkan request ke handler berikutnya
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}