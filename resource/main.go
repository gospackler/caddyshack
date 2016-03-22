package resource

import (
	"strconv"
	"time"
)

// Definition specifies the structure for a resource
type Definition struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DesDoc   string `json:"design"`
	Username string `json:"username"`
	Password string `json:"password"`
	Timeout  int    `json:"timeout"`
	Secure   bool   `json:"secure"`
}

// StrTimeout returns the timeout as a string
func (r *Definition) StrTimeout() string {
	return strconv.Itoa(r.Timeout)
}

// TimeoutDuration returns the port as an integer
func (r *Definition) TimeoutDuration() time.Duration {
	return time.Duration(r.Timeout) * time.Millisecond
}

// StrPort returns the port as a string
func (r *Definition) StrPort() string {
	return strconv.Itoa(r.Port)
}
