package resource

// Definition specifies the structure for a resource
type Definition struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Timeout  string `json:"timeout"`
	Secure   bool   `json:"secure"`
}
