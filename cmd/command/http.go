package command

import (
	"fmt"
	"livecode/config"
	"livecode/internal/adapter/repository"
	"livecode/internal/entrypoint"
	"livecode/internal/entrypoint/handler"
	"livecode/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func httpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Run the HTTP server",
		RunE: func(_ *cobra.Command, args []string) error {
			initializeConfigs()
			return runHttp(&cfg)
		},
	}
}
func runHttp(cfg *config.Config) error {
	db, err := initializeDatabase(cfg)
	if err != nil {
		return fmt.Errorf("fail to init database: %w", err)
	}

	e := echo.New()

	//init repo
	userRepo := repository.NewUserRepository(db)

	//init service
	userService := usecase.NewUserService(userRepo)

	//init handler
	userHandler := handler.NewUserHandler(userService)

	//register routes
	entrypoint.RegisterRoute(e,
		userHandler,
	)

	if err = e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return nil
}

func initializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return db, nil
}
