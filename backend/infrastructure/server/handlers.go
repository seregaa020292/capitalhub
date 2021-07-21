package server

import (
	"github.com/labstack/echo/v4"
	"net/http"

	authHttp "github.com/seregaa020292/capitalhub/internal/auth/delivery/http"
	authRepository "github.com/seregaa020292/capitalhub/internal/auth/repository"
	authUseCase "github.com/seregaa020292/capitalhub/internal/auth/usecase"

	portfolioHttp "github.com/seregaa020292/capitalhub/internal/portfolio/delivery/http"
	portfolioRepository "github.com/seregaa020292/capitalhub/internal/portfolio/repository"
	portfolioUseCase "github.com/seregaa020292/capitalhub/internal/portfolio/usecase"

	assetHttp "github.com/seregaa020292/capitalhub/internal/asset/delivery/http"
	assetRepository "github.com/seregaa020292/capitalhub/internal/asset/repository"
	assetUseCase "github.com/seregaa020292/capitalhub/internal/asset/usecase"

	marketHttp "github.com/seregaa020292/capitalhub/internal/market/delivery/http"
	marketSocket "github.com/seregaa020292/capitalhub/internal/market/delivery/socket"
	marketRepository "github.com/seregaa020292/capitalhub/internal/market/repository"
	marketService "github.com/seregaa020292/capitalhub/internal/market/service"
	marketUseCase "github.com/seregaa020292/capitalhub/internal/market/usecase"

	providerRepository "github.com/seregaa020292/capitalhub/internal/provider/repository"
	providerUseCase "github.com/seregaa020292/capitalhub/internal/provider/usecase"

	instrumentRepository "github.com/seregaa020292/capitalhub/internal/instrument/repository"
	instrumentUseCase "github.com/seregaa020292/capitalhub/internal/instrument/usecase"

	currencyRepository "github.com/seregaa020292/capitalhub/internal/currency/repository"
	currencyUseCase "github.com/seregaa020292/capitalhub/internal/currency/usecase"

	registerRepository "github.com/seregaa020292/capitalhub/internal/register/repository"
	registerUseCase "github.com/seregaa020292/capitalhub/internal/register/usecase"

	sessionRepository "github.com/seregaa020292/capitalhub/infrastructure/session/repository"
	sessionUseCase "github.com/seregaa020292/capitalhub/infrastructure/session/usecase"

	"github.com/seregaa020292/capitalhub/infrastructure/service"
	"github.com/seregaa020292/capitalhub/pkg/mailer"
	"github.com/seregaa020292/capitalhub/pkg/quote"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Map Server Handlers
func (server *Server) MapHandlers(echoInstance *echo.Echo) error {
	// Init clients
	mailerClient := mailer.NewClient(server.cfg.Mailer)
	quotesClient := quote.NewClient(server.webSocketClient, server.cfg, server.logger)

	// Init repositories
	authRepo := authRepository.NewAuthRepository(server.db)
	portfolioRepo := portfolioRepository.NewPortfolioRepository(server.db)
	authAWSRepo := authRepository.NewAuthAWSRepository(server.awsClient)
	authRedisRepo := authRepository.NewAuthRedisRepository(server.redisClient)
	marketRepo := marketRepository.NewMarketRepository(server.db)
	marketRedisRepo := marketRepository.NewMarketRedisRepository(server.redisClient)
	assetRepo := assetRepository.NewAssetRepository(server.db)
	sessionRepo := sessionRepository.NewSessionRepository(server.redisClient, server.cfg)
	providerRepo := providerRepository.NewProviderRepository(server.db)
	instrumentRepo := instrumentRepository.NewInstrumentRepository(server.db)
	currencyRepo := currencyRepository.NewCurrencyRepository(server.db)
	registerRepo := registerRepository.NewRegisterRepository(server.db)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(server.cfg, authRepo, authRedisRepo, authAWSRepo, server.logger)
	portfolioUC := portfolioUseCase.NewPortfolioUseCase(server.cfg, portfolioRepo, server.logger)
	marketUC := marketUseCase.NewMarketUseCase(server.cfg, marketRepo, marketRedisRepo, server.logger)
	assetUC := assetUseCase.NewAssetUseCase(server.cfg, assetRepo, server.logger)
	sessUC := sessionUseCase.NewSessionUseCase(sessionRepo, server.cfg)
	providerUC := providerUseCase.NewProviderUseCase(server.cfg, providerRepo, server.logger)
	instrumentUC := instrumentUseCase.NewInstrumentUseCase(server.cfg, instrumentRepo, server.logger)
	currencyUC := currencyUseCase.NewCurrencyUseCase(server.cfg, currencyRepo, server.logger)
	registerUC := registerUseCase.NewRegisterUseCase(server.cfg, registerRepo, server.logger)

	// Init services
	emailService := service.NewEmailService(server.cfg, mailerClient, server.logger)
	marketTCSParseService := marketService.NewTCSParseService(server.cfg, marketUC, providerUC, instrumentUC, currencyUC, registerUC)

	// Init handlers
	authHttpHandlers := authHttp.NewAuthHandlers(server.cfg, authUC, portfolioUC, sessUC, emailService, server.logger)
	marketHttpHandlers := marketHttp.NewMarketHandlers(server.cfg, marketUC, marketTCSParseService, server.logger)
	marketSocketHandlers := marketSocket.NewMarketHandlers(server.cfg, marketUC, quotesClient, server.logger)
	assetHttpHandlers := assetHttp.NewAssetHandlers(server.cfg, assetUC, portfolioUC, server.logger)
	portfolioHttpHandlers := portfolioHttp.NewPortfolioHandlers(server.cfg, assetUC, portfolioUC, server.logger)

	// Init middlewares
	middleware := InitMiddleware(echoInstance, server, sessUC, authUC)

	// Init http routes
	v1 := echoInstance.Group("/v1")

	authHttpGroup := v1.Group("/auth")
	marketHttpGroup := v1.Group("/market")
	assetHttpGroup := v1.Group("/asset")
	portfolioHttpGroup := v1.Group("/portfolio")

	authHttp.MapAuthRoutes(authHttpGroup, authHttpHandlers, middleware)
	marketHttp.MapMarketRoutes(marketHttpGroup, marketHttpHandlers, middleware)
	assetHttp.MapAssetRoutes(assetHttpGroup, assetHttpHandlers, middleware)
	portfolioHttp.MapPortfolioRoutes(portfolioHttpGroup, portfolioHttpHandlers, middleware)

	// Init websocket routes
	ws := echoInstance.Group("/ws")

	marketSocketGroup := ws.Group("/market")

	marketSocket.MapMarketRoutes(marketSocketGroup, marketSocketHandlers, middleware)

	echoInstance.GET("/health", func(echoCtx echo.Context) error {
		if server.cfg.Server.Mode != "Development" {
			server.logger.Infof("Health check RequestID: %s", utils.GetRequestID(echoCtx))
		}
		return echoCtx.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
