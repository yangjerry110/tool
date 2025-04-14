/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-27 16:03:27
 * @Description: The gormdb package provides functionality for managing GORM database clients.
 * It includes methods for creating, retrieving, and managing database connections and transactions.
 */
package gormdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/yangjerry110/tool/toolerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormDbClient is a struct that manages GORM database clients.
// It includes synchronization mechanisms to ensure thread safety.
type GormDbClient struct {
	SyncOnce  sync.Once  // Ensures that initialization happens only once.
	SyncMutex sync.Mutex // Ensures thread-safe access to shared resources.
}

// transactionContextKey is a type used to store transaction context keys.
type transactionContextKey string

// gormDbClients is a thread-safe map that stores GORM database clients.
var gormDbClients = sync.Map{}

// transactionKey is the key used to store transaction clients in the context.
var transactionKey transactionContextKey = "transaction"

// CreateAllClient initializes all GORM database clients defined in GormDbConfs.
// It ensures that initialization happens only once using sync.Once.
//
// Returns:
//   - error: An error if any issue occurs during client creation.
func (g *GormDbClient) CreateAllClient() error {
	// If there are no configurations, return immediately.
	if len(GormDbConfs) == 0 {
		return nil
	}

	// Use sync.Once to ensure initialization happens only once.
	g.SyncOnce.Do(func() {
		// Iterate over all configurations and create clients.
		for dbName := range GormDbConfs {
			if err := g.CreateClient(dbName); err != nil {
				panic(err)
			}
		}
	})
	return nil
}

// CreateClient initializes a GORM database client for the specified database name.
// It ensures thread-safe access using sync.Mutex.
//
// Parameters:
//   - dbName: The name of the database to initialize.
//
// Returns:
//   - error: An error if any issue occurs during client creation.
func (g *GormDbClient) CreateClient(dbName string) error {
	// Lock to ensure thread-safe access.
	g.SyncMutex.Lock()
	// Unlock when the function exits.
	defer g.SyncMutex.Unlock()

	// Retrieve the configuration for the specified database name.
	gormDbConf, isExist := GormDbConfs[dbName]
	if !isExist {
		return toolerrors.WithFields("dbName", dbName).New("db Err : gormDB CreateClient conf is not exist")
	}

	// Configure the GORM client.
	config := &gorm.Config{}
	config.SkipDefaultTransaction = gormDbConf.SkipDefaultTransaction
	config.Logger = logger.Default.LogMode(gormDbConf.LoggerLevel)

	// Initialize the GORM client using the MySQL driver.
	db, err := gorm.Open(mysql.Open(gormDbConf.Dsn), config)
	if err != nil {
		return err
	}

	// Store the client in the thread-safe map.
	gormDbClients.Store(dbName, db)
	return nil
}

// GetClient retrieves a GORM database client for the specified database name.
// It checks for an existing transaction client in the context before returning the default client.
//
// Parameters:
//   - ctx: The context containing the transaction client (if any).
//   - dbName: The name of the database to retrieve the client for.
//
// Returns:
//   - *gorm.DB: The GORM database client.
//   - error: An error if the client does not exist.
func (g *GormDbClient) GetClient(ctx context.Context, dbName string) (*gorm.DB, error) {
	// Check for an existing transaction client in the context.
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if isExisttransactionDbClient && transactionDbClient != nil {
		return transactionDbClient, nil
	}

	// Retrieve the default client from the thread-safe map.
	gormDbClient, isExist := gormDbClients.Load(dbName)
	if !isExist {
		return nil, toolerrors.WithFields("dbName", dbName).New("db Err : gormDB GetClient client is not exist")
	}

	// Return the default client.
	return gormDbClient.(*gorm.DB), nil
}

// TransactionBegin starts a new transaction for the specified database name.
// It stores the transaction client in the context for later use.
//
// Parameters:
//   - ctx: The context to store the transaction client.
//   - dbName: The name of the database to start the transaction for.
//
// Returns:
//   - context.Context: The updated context containing the transaction client.
//   - error: An error if the transaction cannot be started.
func (g *GormDbClient) TransactionBegin(ctx context.Context, dbName string) (context.Context, error) {
	// Retrieve the default client for the specified database name.
	dbClient, dbClientIsExist := gormDbClients.Load(dbName)
	if !dbClientIsExist {
		return ctx, toolerrors.WithFields("dbName", dbName).New("db Err : gormDB TransactionBegin client is not exist")
	}

	// Start a new transaction.
	transactionDbClient := dbClient.(*gorm.DB).Begin()

	// Check for errors during transaction initialization.
	if transactionDbClient.Error != nil {
		return ctx, transactionDbClient.Error
	}

	// Store the transaction client in the context.
	return context.WithValue(ctx, g.transactionKey(dbName), transactionDbClient), nil
}

// TransactionCommit commits the transaction for the specified database name.
//
// Parameters:
//   - ctx: The context containing the transaction client.
//   - dbName: The name of the database to commit the transaction for.
//
// Returns:
//   - error: An error if the transaction cannot be committed.
func (g *GormDbClient) TransactionCommit(ctx context.Context, dbName string) error {
	// Retrieve the transaction client from the context.
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if !isExisttransactionDbClient {
		return toolerrors.WithFields("dbName", dbName).New("db Err : gormDB TransactionCommit client is not exist")
	}

	// Commit the transaction.
	if err := transactionDbClient.Commit().Error; err != nil {
		return err
	}
	return nil
}

// TransactionRollback rolls back the transaction for the specified database name.
//
// Parameters:
//   - ctx: The context containing the transaction client.
//   - dbName: The name of the database to roll back the transaction for.
//
// Returns:
//   - error: An error if the transaction cannot be rolled back.
func (g *GormDbClient) TransactionRollback(ctx context.Context, dbName string) error {
	// Retrieve the transaction client from the context.
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if !isExisttransactionDbClient {
		return toolerrors.WithFields("dbName", dbName).New("db Err : gormDB TransactionRollback client is not exist")
	}

	// Roll back the transaction.
	if err := transactionDbClient.Rollback().Error; err != nil {
		return err
	}
	return nil
}

// transactionKey generates a unique key for storing transaction clients in the context.
//
// Parameters:
//   - dbName: The name of the database.
//
// Returns:
//   - transactionContextKey: The unique key for the transaction client.
func (g *GormDbClient) transactionKey(dbName string) transactionContextKey {
	return transactionContextKey(fmt.Sprintf("%s-%s", dbName, transactionKey))
}
