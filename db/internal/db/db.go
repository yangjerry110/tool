/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:55:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:28:33
 * @Description: The db package provides an interface for managing database clients and transactions.
 * It defines the DbInterface and a factory function to create database clients.
 */
package db

import (
	"context"

	"gorm.io/gorm"
)

// DbInterface defines the methods required to manage database clients and transactions.
// It provides functionality for creating clients, retrieving clients, and managing transactions.
type DbInterface interface {
	// CreateAllClient initializes all database clients.
	CreateAllClient() error

	// CreateClient initializes a database client for the specified database name.
	CreateClient(dbName string) error

	// GetClient retrieves a database client for the specified database name.
	GetClient(ctx context.Context, dbName string) (*gorm.DB, error)

	// TransactionBegin starts a new transaction for the specified database name.
	TransactionBegin(ctx context.Context, dbName string) (context.Context, error)

	// TransactionCommit commits the transaction for the specified database name.
	TransactionCommit(ctx context.Context, dbName string) error

	// TransactionRollback rolls back the transaction for the specified database name.
	TransactionRollback(ctx context.Context, dbName string) error
}

// CreateDb is a factory function that returns an implementation of the DbInterface.
// It allows for flexible creation of database clients based on the provided DbInterface implementation.
//
// Parameters:
//   - DbInterface: The implementation of the DbInterface to be used for creating database clients.
//
// Returns:
//   - DbInterface: The provided DbInterface implementation.
func CreateDb(DbInterface DbInterface) DbInterface {
	return DbInterface
}
