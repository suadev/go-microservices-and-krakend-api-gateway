package httpclient

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/suadev/microservices/services.order/src/dto"
)

var pact dsl.Pact

func TestMain(m *testing.M) {

	setupMockServer()
	// Run all the tests
	var exitCode = m.Run()
	// Shutdown the Mock Service and Write pact files to disk
	pact.WritePact()
	pact.Teardown()

	err := publishPacts()

	if err != nil {
		log.Println("ERROR: ", err)
		os.Exit(1)
	}
	os.Exit(exitCode)
}

func publishPacts() error {

	publisher := dsl.Publisher{}

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../../../_pacts", dir)

	return publisher.Publish(types.PublishRequest{
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/order-service-customer-service.json", pactDir))},
		ConsumerVersion: "1.0.0",
		PactBroker:      "http://localhost:9292",
		BrokerUsername:  "admin",
		BrokerPassword:  "admin",
	})
}

var term = dsl.Term

type request = dsl.Request

func Test_ClientPact_Get_Basket_Items(t *testing.T) {
	t.Run("Get Basket Items...", func(t *testing.T) {
		sampleUUID := uuid.MustParse("e153ef59-5708-48a4-848b-a65bd2667ac4")
		UUIDRegex := "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"
		customerID := sampleUUID.String()

		basketItemsResponse := [1]dto.BasketItemDto{
			{
				ID:          sampleUUID,
				BasketID:    sampleUUID,
				ProductID:   sampleUUID,
				ProductName: "Sample Product",
				UnitPrice:   5,
				Quantity:    1,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			}}

		pact.
			AddInteraction().
			Given("There is an available basket for the customer").
			UponReceiving("A GET request to retrieve customer basket items.").
			WithRequest(request{
				Method: "GET",
				Path:   term("/customers/"+customerID+"/basketItems", "/customers/"+UUIDRegex+"/basketItems"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body:   dsl.Like(basketItemsResponse),
				Headers: dsl.MapMatcher{
					"Content-Type": term("application/json; charset=utf-8", `application\/json`),
				},
			})

		err := pact.Verify(func() error {
			_, err := client.GetBasketItems(customerID)
			return err
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}

var client *Client

func setupMockServer() {
	pact = createPact()
	// Start service to get access to the port
	pact.Setup(true)
	client = &Client{
		hostURL: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 "order-service",
		Provider:                 "customer-service",
		LogDir:                   "../../../_logs",
		PactDir:                  "../../../_pacts",
		LogLevel:                 "DEBUG",
		DisableToolValidityCheck: true,
	}
}
