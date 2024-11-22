package forum

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
    visitors = make(map[string]*visitor)
    mu       sync.Mutex
    rateLimit	= 1 // requests per second
    burstLimit	= 5 // maximum burst size
)

type visitor struct {
    limiter  *rate.Limiter
    lastSeen time.Time
}

// GetVisitor retourne le rate limiter pour une adresse IP donnée.
// Si l'adresse IP n'existe pas encore dans le map, un nouveau rate limiter est créé.
func GetVisitor(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()

	v, exists := visitors[ip]
    if !exists {
        limiter := rate.NewLimiter(rate.Limit(rateLimit), burstLimit)
        visitors[ip] = &visitor{limiter, time.Now()}
        return limiter
    }

    v.lastSeen = time.Now()
    return v.limiter
}

// CleanupVisitors supprime les visiteurs qui n'ont pas été vus depuis plus de 3 minutes.
// Cette fonction est exécutée périodiquement pour nettoyer le map des visiteurs.
func CleanupVisitors() {
    for {
        time.Sleep(time.Minute)
        mu.Lock()
        for ip, v := range visitors {
            if time.Since(v.lastSeen) > 3*time.Minute {
                delete(visitors, ip)
            }
        }
        mu.Unlock()
    }
}

// LimitMiddleware est un middleware HTTP qui limite le nombre de requêtes par seconde pour chaque adresse IP.
// Si la limite est dépassée, une réponse HTTP 429 Too Many Requests est renvoyée.
func LimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        limiter := GetVisitor(ip)
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            fmt.Printf("Too Many Requests from %s\n", ip)
            return
        }
        fmt.Printf("Request allowed from %s\n", ip)
        next.ServeHTTP(w, r)
    })
}
