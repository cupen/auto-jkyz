package config

type Config struct {
	Account *Account
	Chrome  *Chrome
}

type Account struct {
	IDType   string
	Username string
	Password string
}

type Chrome struct {
	Path string
}

type Sevice struct {
	Verify struct {
		ClientID string `toml:"client_id"`
		Secret   string
	}
}

func (c *Chrome) GetPath() string {
	if c == nil || c.Path == "" {
		return "C:/Program Files/Google/Chrome/Application/chrome.exe"
	}
	return c.Path
}
