package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent/dialect"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen/role"
	"github.com/kimsehyoung/dongle/app/auth/server/validator"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/logger"
)

type Account struct {
	email          string
	hashedPassword string
	name           string
	phoneNumber    string
}

var rootAccount *Account

func rollback(tx *authdbgen.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

// Parse root account Info
func parse() error {
	email := flag.String("email", "", "text@example.com")
	password := flag.String("password", "", "password")
	name := flag.String("name", "", "name")
	phoneNumber := flag.String("number", "", "010-1234-1234")
	flag.Parse()

	err := validator.IsValidAccount(&validator.Account{
		Email:       *email,
		Password:    *password,
		Name:        *name,
		PhoneNumber: *phoneNumber,
	})
	if err != nil {
		return fmt.Errorf("Failed to validate Account: %v ", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Failed to generate hashed password: %v", err)
	}
	rootAccount = &Account{
		email:          *email,
		hashedPassword: string(hashedPassword),
		name:           *name,
		phoneNumber:    *phoneNumber,
	}
	return nil
}

// Connect to db with auth schema
func open() (*authdbgen.Client, error) {
	// Create DB client
	db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s?search_path=%s",
		dialect.Postgres,
		env.GetEnvString("DONGLE_DB_USER", "postgres"),
		env.GetEnvString("DONGLE_DB_PASSWORD", "1234"),
		env.GetEnvString("DONGLE_DB_HOST", "dongle-postgres"),
		env.GetEnvString("DONGLE_DB_PORT", "5432"),
		env.GetEnvString("DONGLE_DB_NAME", "dongle"),
		env.GetEnvString("DONGLE_DB_SCHEMA", "auth"),
	))
	if err != nil {
		return nil, fmt.Errorf("Failed to create dongle DB client: %v", err)
	}
	// Create Schema
	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA %[1]s", env.GetEnvString("DONGLE_DB_SCHEMA", "auth")))
	if err != nil {
		return nil, fmt.Errorf("Failed to create schema: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := authdbgen.NewClient(authdbgen.Driver(drv))
	return client, nil
}

// Create 'root', 'admin', 'user' roles
func createRoles(ctx context.Context, tx *authdbgen.Tx) error {
	_, err := tx.Role.Create().SetType("root").Save(ctx)
	if err != nil {
		return fmt.Errorf("Create role %v", err)
	}

	_, err = tx.Role.Create().SetType("admin").Save(ctx)
	if err != nil {
		return fmt.Errorf("Create role %v", err)
	}

	_, err = tx.Role.Create().SetType("user").Save(ctx)
	if err != nil {
		return fmt.Errorf("Create role %v", err)
	}
	return nil
}

// Create 'root' account
func createRoot(ctx context.Context, tx *authdbgen.Tx) error {
	role, err := tx.Role.
		Query().
		Where(role.Type("root")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("Check root account %v", err)
	}

	_, err = tx.Account.
		Create().
		SetRole(role).
		SetEmail(rootAccount.email).
		SetHashedPassword(rootAccount.hashedPassword).
		SetName(rootAccount.name).
		SetPhoneNumber(rootAccount.phoneNumber).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("Create root account %v", err)
	}
	return nil
}

func main() {
	// Parse root account Info
	if err := parse(); err != nil {
		logger.Fatalf("[parse fail] %v", err)
	}

	// Connect to db with auth schema
	client, err := open()
	if err != nil {
		logger.Fatalf("[open fail] %v", err)
	}
	defer client.Close()

	// Create tables
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logger.Fatalf("Failed to create tables: %v", err)
	}

	// Create transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		logger.Fatalf("Failed to create transaction: %v", err)
	}

	// Create roles
	if err := createRoles(ctx, tx); err != nil {
		logger.Fatalf("Failed to create roles: %v", rollback(tx, err))
	}

	// Create root account
	if err := createRoot(ctx, tx); err != nil {
		logger.Fatalf("Failed to create root: %v", rollback(tx, err))
	}

	// Commit the transaction.
	err = tx.Commit()
	if err != nil {
		logger.Fatalf("Failed to commit transaction: %v", err)
	}
	logger.Infof("The Dongle Database is initialized successfully(root: %s)", rootAccount.email)
}
