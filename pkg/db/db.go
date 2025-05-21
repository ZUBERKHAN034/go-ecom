package db

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql" // Import the gorm mysql driver with an alias
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// Read the CA certificate
	caCert := []byte(config.Env.DBCACert)

	// Create a certificate pool and add the CA certificate
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, fmt.Errorf("failed to append CA certificate")
	}

	// Register a custom TLS config
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	mysql.RegisterTLSConfig("custom", tlsConfig)

	// Configure the MySQL using string format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=custom&parseTime=True&loc=Local",
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBPublicHost,
		config.Env.DBName,
	)

	log.Println("Database connection opening & establishing...")
	// Open a connection to the database using the gorm mysql driver
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection opened & established successfully")
	return db, nil
}
