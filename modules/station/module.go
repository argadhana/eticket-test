package station

import (
	"eticket-test/internal/pkg/bus"
	"eticket-test/internal/pkg/logger"
	"eticket-test/internal/pkg/middleware"
	"eticket-test/modules/station/domain/entity"
	"eticket-test/modules/station/domain/repository"
	"eticket-test/modules/station/domain/service"
	"eticket-test/modules/station/handler"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Module struct {
	db             *gorm.DB
	log            *logger.Logger
	event          *bus.EventBus
	stationService *service.StationService
	stationHandler *handler.StationHandler
	stationRepo    repository.StationRepository
	middleware     *middleware.MiddlewareProvider
}

func (m *Module) Name() string {
	return "station"
}

func (m *Module) Initialize(db *gorm.DB, log *logger.Logger, event *bus.EventBus) error {
	m.db = db
	m.log = log
	m.event = event

	m.stationRepo = repository.NewStationRepositoryImpl(m.db)
	m.stationService = service.NewStationService(m.stationRepo)
	m.stationHandler = handler.NewStationHandler(m.log, m.stationService)

	// init middleware provider
	m.middleware = middleware.NewMiddlewareProvider()

	return nil
}

func (m *Module) RegisterRoutes(e *echo.Echo, basePath string) {
	// pakai group + middleware JWTAuth
	g := e.Group(basePath+"/stations", m.middleware.GetJWTAuth())
	m.stationHandler.RegisterRoutes(g) // biar handler bisa register di group yg sudah protected
}

func (m *Module) Migrations() error {
	return m.db.AutoMigrate(&entity.Station{})
}

func (m *Module) Logger() *logger.Logger {
	return m.log
}

func NewModule() *Module {
	return &Module{}
}
