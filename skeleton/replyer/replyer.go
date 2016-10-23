package replyer

import "net/http"

// Replyer write result to res
type Replyer func(res http.ResponseWriter)
