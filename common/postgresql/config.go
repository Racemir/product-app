package postgresql

// postsgl'e hangi bilgilerle bağlanacağız
type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime string
}
