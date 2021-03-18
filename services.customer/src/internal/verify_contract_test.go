package customer

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"github.com/suadev/microservices/services.customer/src/entity"
)

var port, _ = utils.GetFreePort()

func Test_Verify_Pacts(t *testing.T) {

	pact := dsl.Pact{
		Provider:                 "customer-service",
		LogDir:                   "../../../_logs",
		PactDir:                  "../../../_pacts",
		DisableToolValidityCheck: true,
		LogLevel:                 "INFO",
	}

	go startProvider()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../../../_pacts", dir)

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:" + strconv.Itoa(port),
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/order-service-customer-service.json", pactDir))},
		// BrokerURL:                  "http://localhost:9292",
		BrokerUsername:  "admin",
		BrokerPassword:  "admin",
		ProviderVersion: "1.0.0",
		// PublishVerificationResults: true,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func getBasketItems(context *gin.Context) {
	basketItems := [1]entity.BasketItem{{}}
	context.JSON(http.StatusOK, basketItems)
}

func startProvider() {
	router := gin.Default()
	router.GET("/customers/:id/basketItems", getBasketItems)
	router.Run(fmt.Sprintf(":%d", port))
}
