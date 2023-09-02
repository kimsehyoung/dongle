package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent/dialect"
	"github.com/kimsehyoung/dongle/app/auth/ent/authgen"
	"github.com/kimsehyoung/dongle/app/auth/ent/authgen/role"
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

func rollback(tx *authgen.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func initDatabase(ctx context.Context, tx *authgen.Tx) error {
	// Create 'root', 'admin', 'user' roles
	_, err := tx.Role.Create().SetType("root").Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.Role.Create().SetType("admin").Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.Role.Create().SetType("user").Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	// Create 'root' account
	role, err := tx.Role.
		Query().
		Where(role.Type("root")).
		Only(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	account, err := tx.Account.
		Create().
		SetRole(role).
		SetEmail(rootAccount.email).
		SetHashedPassword(rootAccount.hashedPassword).
		SetName(rootAccount.name).
		SetPhoneNumber(rootAccount.phoneNumber).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	logger.Infof("The root account %s has been created", account.Email)

	// Commit the transaction.
	return tx.Commit()
}

func main() {
	// Parse root account Info
	email := flag.String("email", "", "text@example.com")
	password := flag.String("password", "", "password")
	name := flag.String("name", "", "name")
	phoneNumber := flag.String("numer", "", "010-1234-1234")
	flag.Parse()

	err := validator.IsValidAccount(&validator.Account{
		Email:       *email,
		Password:    *password,
		Name:        *name,
		PhoneNumber: *phoneNumber,
	})
	if err != nil {
		logger.Errorf("Failed to validate Account: %v ", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("Failed to generate hashed password: %v", err)
	}
	rootAccount = &Account{
		email:          *email,
		hashedPassword: string(hashedPassword),
		name:           *name,
		phoneNumber:    *phoneNumber,
	}

	// Create DB client
	db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dialect.Postgres,
		env.GetEnvString("AUTH_DB_USER", "postgres"),
		env.GetEnvString("AUTH_DB_PASSWORD", "1234"),
		env.GetEnvString("AUTH_DB_HOST", "auth-postgres"),
		env.GetEnvString("AUTH_DB_PORT", "5432"),
		env.GetEnvString("AUTH_DB_NAME", "auth"),
	))
	if err != nil {
		logger.Errorf("Failed to create auth DB client: %v", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := authgen.NewClient(authgen.Driver(drv))
	defer client.Close()

	// Create transaction
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		logger.Fatalf("Failed creating transaction: %v", err)
	}

	// Create schema
	if err := client.Schema.Create(ctx); err != nil {
		logger.Fatalf("Failed creating schema: %v", err)
	}

	err = initDatabase(ctx, tx)
	if err != nil {
		logger.Fatalf("Failed to initialize auth DB: %v", err)
	}
	logger.Info("The Auth Database is initialized successfully")
}
