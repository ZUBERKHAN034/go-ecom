package db

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql" // Import the gorm mysql driver with an alias
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Connect() error {
	// Read the CA certificate
	caCert := []byte(config.Env.DBCACert)

	// Create a certificate pool and add the CA certificate
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate")
	}

	// Register a custom TLS config
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	mysql.RegisterTLSConfig("custom", tlsConfig)

	// Configure the MySQL DSN
	dsnConfig := &mysql.Config{
		User:      config.Env.DBUser,
		Passwd:    config.Env.DBPassword,
		Addr:      config.Env.DBAddress,
		DBName:    config.Env.DBName,
		ParseTime: true,
		TLSConfig: "custom",
		Net:       "tcp",
	}
	dsn := dsnConfig.FormatDSN()

	// Open a connection to the database using the gorm mysql driver
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Ping the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set the dbInstance
	dbInstance = db

	log.Println("Connected to the database successfully!")
	return nil
}

func GetDB() *gorm.DB {
	return dbInstance
}
