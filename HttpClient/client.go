package main

import (
  "encoding/json"
  "net/http"
  "regexp"
)

// Структуры с гарантированным порядком полей для JSON
type Response struct {
  Status string      `json:"status"`
  Result interface{} `json:"result"`
}

// Структура для нормального ответа с данными
type ResultData struct {
  Greetings string `json:"greetings"`
  Name      string `json:"name"`
}

// Middleware для установки имени по умолчанию
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    if q.Get("name") == "" {
      q.Set("name", "stranger")
      r.URL.RawQuery = q.Encode()
    }
    next(w, r)
  }
}

// Middleware для проверки имени
func Sanitize(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    isValidName := regexp.MustCompile(`^[A-Za-z]+$`).MatchString(name)
    if !isValidName {
      // Возвращаем panic при некорректном имени
      panic("Invalid name")
    }
    next(w, r)
  }
}

// Обработчик Hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Query().Get("name")
  response := Response{
    Status: "ok",
  }

  // В нормальном случае, заполняем структуру с результатом
  response.Result = ResultData{
    Greetings: "hello",
    Name:      name,
  }

  // Отправляем ответ в формате JSON с нужным порядком полей
  w.Header().Set("Content-Type", "application/json")
  // Кодируем ответ в JSON
  json.NewEncoder(w).Encode(response)
}

// Middleware для обработки panic
func RPC(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // Обработчик panic
    defer func() {
      if err := recover(); err != nil {
        // Возвращаем ошибку в формате JSON при панике
        w.Header().Set("Content-Type", "application/json")
        // Ошибка должна быть с кодом 200 OK, как в тестах
        w.WriteHeader(http.StatusOK)

        // Ошибка - пустая структура для result
        errorResponse := Response{
          Status: "error",
          Result: struct{}{}, // Пустая структура
        }
        // Отправляем ошибку в правильном формате
        json.NewEncoder(w).Encode(errorResponse)
      }
    }()
    // Если нет паники, продолжаем обработку
    next(w, r)
  }
}

func main() {
  // Используем RPC, SetDefaultName и Sanitize для обработки запросов
  http.HandleFunc("/hello", RPC(SetDefaultName(Sanitize(HelloHandler))))
  http.ListenAndServe(":8080", nil)
}