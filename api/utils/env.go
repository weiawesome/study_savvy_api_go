package utils

import "os"

func EnvMySqlDb() string {
	var dbName string
	if dbName = os.Getenv("MYSQL_DB"); len(dbName) == 0 {
		dbName = "DefaultDb"
	}
	return dbName
}
func EnvMySqlAddress() string {
	var ip string
	var port string
	if ip = os.Getenv("MYSQL_IP"); len(ip) == 0 {
		ip = "localhost"
	}
	if port = os.Getenv("MYSQL_PORT"); len(port) == 0 {
		port = "3306"
	}
	return ip + ":" + port
}
func EnvMySqlPassword() string {
	var password string
	if password = os.Getenv("MYSQL_PASSWORD"); len(password) == 0 {
		password = "DefaultPassword"
	}
	return password
}
func EnvMySqlUser() string {
	var user string
	if user = os.Getenv("MYSQL_USER"); len(user) == 0 {
		user = "DefaultUser"
	}
	return user
}
func EnvJwtSecret() string {
	var jwtSecret string
	if jwtSecret = os.Getenv("JWT_SECRET"); len(jwtSecret) == 0 {
		jwtSecret = "DefaultJwtSecret"
	}
	return jwtSecret
}
func EnvJwtIssuer() string {
	var issuer string
	if issuer = os.Getenv("JWT_ISSUER"); len(issuer) == 0 {
		issuer = "DefaultIssuer"
	}
	return issuer
}
func EnvJwtExpireDays() string {
	var jwtExpireDays string
	if jwtExpireDays = os.Getenv("JWT_EXPIRE_DAYS"); len(jwtExpireDays) == 0 {
		jwtExpireDays = "14"
	}
	return jwtExpireDays
}
func EnvRedisPassword() string {
	var password string
	if password = os.Getenv("REDIS_PASSWORD"); len(password) == 0 {
		password = ""
	}
	return password
}
func EnvRedisAddress() string {
	var ip string
	var port string
	if ip = os.Getenv("REDIS_IP"); len(ip) == 0 {
		ip = "localhost"
	}
	if port = os.Getenv("REDIS_PORT"); len(port) == 0 {
		port = "6379"
	}
	return ip + ":" + port
}
func EnvRedisDb() string {
	var db string
	if db = os.Getenv("REDIS_DB"); len(db) == 0 {
		db = "0"
	}
	return db
}
func EnvPageSize() string {
	var pageSize string
	if pageSize = os.Getenv("PAGE_SIZE"); len(pageSize) == 0 {
		pageSize = "10"
	}
	return pageSize
}
