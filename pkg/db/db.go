package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*sql.DB, error) {
	// Construct the path to ca.pem
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	caPath := filepath.Join(dir, "pkg/db", "ca.pem")

	// Read the CA certificate
	caCert, err := os.ReadFile(caPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a certificate pool and add the CA certificate
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate")
	}

	// Configure the MySQL driver with SSL/TLS
	cfg := mysql.Config{
		User:      config.Env.DBUser,
		Passwd:    config.Env.DBPassword,
		Addr:      config.Env.DBAddress,
		DBName:    config.Env.DBName,
		TLSConfig: "custom",
		Net:       "tcp",
	}

	// Register a custom TLS config
	err = mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: certPool,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Open a connection to the database
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database successfully!")
	return db, nil
}
