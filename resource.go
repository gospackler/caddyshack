package caddyshack

type Resource struct {
    Host string     `json:"host"`
    Port string     `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Timeout string  `json:"timeout"`
}
