package configs

export hostKey 	   = key("hostKey")
export usernameKey = key("usernameKey")
export passwordKey = key("passwordKey")
export databaseKey = key("databaseKey")



type Config struct {
	Debug       bool
	Port        int
	User        string
	Users       []string
	Rate        float32
	ColorCodes  map[string]int
}



