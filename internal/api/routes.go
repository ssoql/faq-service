package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/internal/infrastructure"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/internal/useCases"

	"github.com/ssoql/faq-service/internal/api/handlers"
)

func RegisterRoutes(router *gin.Engine, dbClient *db.ClientDB) {

	faqReadRepository := infrastructure.NewFaqReadRepository(dbClient)
	faqWriteRepository := infrastructure.NewFaqWriteRepository(dbClient)

	faqGetUseCase := useCases.NewGetFaqUseCase(faqReadRepository)
	saveFaqUseCase := useCases.NewCreateFaqUseCase(faqWriteRepository)
	updateFaqUseCase := useCases.NewUpdateFaqUseCase(faqWriteRepository)
	deleteFaqUseCase := useCases.NewDeleteFaqUseCase(faqWriteRepository)

	faqCreateHandler := handlers.NewFaqCreateHandler(saveFaqUseCase)
	faqGetHandler := handlers.NewFaqGetHandler(faqGetUseCase)
	faqUpdateHandler := handlers.NewFaqUpdateHandler(updateFaqUseCase)
	faqDeleteHandler := handlers.NewFaqDeleteHandler(deleteFaqUseCase)

	router.POST("/faq", faqCreateHandler.Handle)
	router.GET("/faq/:faq_id", faqGetHandler.Handle)
	router.PATCH("/faq/:faq_id", faqUpdateHandler.Handle)
	router.DELETE("/faq/:faq_id", faqDeleteHandler.Handle)
	router.POST("/faqs", faqCreateHandler.Handle)
}
