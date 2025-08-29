package auth

import (
	"eticket-test/internal/pkg/bus"
	"eticket-test/internal/pkg/config"
	"eticket-test/internal/pkg/database"
	"eticket-test/internal/pkg/jwt"
	"eticket-test/internal/pkg/logger"
	"eticket-test/modules/auth/domain/entity"
	"eticket-test/modules/auth/domain/repository"
	"eticket-test/modules/auth/domain/service"
	"eticket-test/modules/auth/handler"
	"github.com/labstack/echo"

	"gorm.io/gorm"
)

// Module implements the application Module interface for the auth module
type Module struct {
	db          *gorm.DB
	logger      *logger.Logger
	event       *bus.EventBus
	authService *service.AuthService
	authHandler *handler.AuthHandler
	authRepo    repository.AuthUserRepository
}

// Name returns the name of the module
func (m *Module) Name() string {
	return "auth"
}

// Initialize initializes the module
func (m *Module) Initialize(db *gorm.DB, log *logger.Logger, event *bus.EventBus) error {
	m.db = db
	m.logger = log
	m.event = event
	signatureKey := config.GetString("jwt.signature_key")
	dayExpired := config.GetInt("jwt.day_expired")
	jwtImpl := jwt.NewJWTImpl(signatureKey, dayExpired)
	m.logger.Info("Initializing auth module")

	m.authRepo = repository.NewUserRepositoryImpl(m.db)

	// Services
	m.authService = service.NewAuthService(m.authRepo, jwtImpl)

	// Handlers
	m.authHandler = handler.NewAuthHandler(m.logger, m.authService)

	m.logger.Info("Auth module initialized successfully")
	return nil
}

// RegisterRoutes registers the module's routes
func (m *Module) RegisterRoutes(e *echo.Echo, basePath string) {
	m.logger.Info("Registering auth routes at %s", basePath)
	m.authHandler.RegisterRoutes(e, basePath)
	m.logger.Debug("Auth routes registered successfully")
}

// Migrations returns the module's migrations (auth tidak perlu tabel sendiri)
func (m *Module) Migrations() error {
	m.db.AutoMigrate(&entity.User{})
	database.Seed(m.db)
	return nil
}

// Logger returns the module's logger
func (m *Module) Logger() *logger.Logger {
	return m.logger
}

// NewModule creates a new auth module
func NewModule() *Module {
	return &Module{}
}
