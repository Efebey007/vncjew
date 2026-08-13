package main

import "time"
import "os"

var adminPassword = getEnv("ADMIN_PASS", "admin123")
var clientPassword = getEnv("CLIENT_PASS", "client123")

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok && value != "" {
        return value
    }
    return fallback
}

var CFGAdminAccount = map[string]string{
    "admin": adminPassword,
}

var CFGClientAccount = map[string]string{
    "client": clientPassword,
    "admin": CFGAdminAccount["admin"],
}
var CFGMaxVNCConns = 100
var CFGMaxConcurrentOCR = 1
var CFGDb = "database.sqlite3"
var CFGClientPing = 5 * time.Second
var CFGClientTimeout = 60 * time.Second
var CFGIPInfoToken = "fd737c5e5030e3"
var CFGPasswords = []string{"123456", "password", "admin", "user", "default", "", "123456789", "111111", "password", "qwerty", "abc123", "12345678", "password1", "1234567", "123123", "pu", "god", "secret"}
var CFGVNCTimeout = "15"
var CFGVNCScreenshotBin = "./vncscreenshot"
var CFGTesseractBin = "tesseract"
