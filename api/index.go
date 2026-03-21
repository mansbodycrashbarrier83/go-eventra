package handler

import (
  "context"
  "encoding/json"
  "net/http"
  "sync"

  "eventra/bootstrap"
)

var (
  initOnce sync.Once
  initErr  error
  router   http.Handler
)

func Handler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/health" {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "eventra-auth"})
    return
  }

  initOnce.Do(func() {
    router, initErr = bootstrap.NewHTTPHandler(context.Background())
  })

  if initErr != nil {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "error": "backend initialization failed",
    })
    return
  }

  router.ServeHTTP(w, r)
}
