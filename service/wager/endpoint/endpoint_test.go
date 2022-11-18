package wagerep

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	"strings"
	"testing"
)

/* ----------------------------------------------------------------
						Place Wager Benchmark
---------------------------------------------------------------- */

// Input
var placeWagerInputs = map[int]string{
	0: `{"odds":90,"selling_percentage":100,"selling_price":100000,"total_wager_value":50}`,
	1: `{"odds":90,"selling_percentage":100,"selling_price":200,"total_wager_value":50}`,
	2: `{"odds":0,"selling_percentage":100,"selling_price":200,"total_wager_value":50}`,
	3: `{"odds":90,"selling_percentage":0,"selling_price":200,"total_wager_value":50}`,
	4: `{"odds":90,"selling_percentage":100,"selling_price":1,"total_wager_value":50}`,
	5: `{"odds":90,"selling_percentage":100,"selling_price":150,"total_wager_value":50}`,
	6: `{"odds":90,"selling_percentage":100,"selling_price":100,"total_wager_value":50}`,
	7: `{"odds":74.2,"selling_percentage":51.7,"selling_price":956.8422,"total_wager_value":50.33}`,
	8: `{"odds":74.2,"selling_percentage":51.7,"selling_price":956.99,"total_wager_value":50.33}`,
	9: `{"odds":74.2,"selling_percentage":51.7,"selling_price":956.8499,"total_wager_value":50.33}`,
}

// Output
var placeWagerOutputs = map[int]int{
	0: 201,
	1: 201,
	2: 422,
	3: 422,
	4: 400,
	5: 201,
	6: 201,
	7: 201,
	8: 201,
	9: 201,
}

// Body
var placeWagerBody = map[int]models.Wager{
	0: {
		Odds:              90,
		SellingPercentage: 100,
		SellingPrice:      100000,
		TotalWagerValue:   50,
	},
	1: {
		Odds:              90,
		SellingPercentage: 100,
		SellingPrice:      200,
		TotalWagerValue:   50,
	},
	5: {
		Odds:              90,
		SellingPercentage: 100,
		SellingPrice:      150,
		TotalWagerValue:   50,
	},
	6: {
		Odds:              90,
		SellingPercentage: 100,
		SellingPrice:      100,
		TotalWagerValue:   50,
	},
	7: {
		Odds:              74,
		SellingPercentage: 52,
		SellingPrice:      956.84,
		TotalWagerValue:   50,
	},
	8: {
		Odds:              74,
		SellingPercentage: 52,
		SellingPrice:      956.99,
		TotalWagerValue:   50,
	},
	9: {
		Odds:              74,
		SellingPercentage: 52,
		SellingPrice:      956.85,
		TotalWagerValue:   50,
	},
}

func makePlaceWagerInput(r int) FakeContext {
	input := FakeContext{
		Method: "POST",
		Body:   placeWagerInputs[r],
	}
	return input
}

func checkPlaceWager(b *testing.B) {
	r := rand.Intn(10)
	input := makePlaceWagerInput(r)
	ctx, rec := createEchoContext(input)
	_ = PlaceWager(ctx)

	//if rec.Code != placeWagerOutputs[r] {
	//	b.Errorf("Expected %d, got %d - %s - %d", placeWagerOutputs[r], rec.Code, rec.Body, r)
	//}

	assert.Equal(b, placeWagerOutputs[r], rec.Code)
	if rec.Code == 201 {
		data := models.Wager{}
		json.Unmarshal(rec.Body.Bytes(), &data)
		assert.Equal(b, placeWagerBody[r].Odds, data.Odds)
		assert.Equal(b, placeWagerBody[r].SellingPercentage, data.SellingPercentage)
		assert.Equal(b, placeWagerBody[r].SellingPrice, data.SellingPrice)
		assert.Equal(b, placeWagerBody[r].TotalWagerValue, data.TotalWagerValue)
	}
}

func BenchmarkPlaceWager(b *testing.B) {
	// Init Fake DB
	sqldb.InitFakeDB()

	for i := 0; i < b.N; i++ {
		checkPlaceWager(b)
	}
}

/* ----------------------------------------------------------------
						Buy Wager Benchmark
---------------------------------------------------------------- */

// Input
var buyWagerInputs = map[int]string{
	0: `{"buying_price":-1}`,
	1: `{"buying_price":999999}`,
	2: `{"buying_price":100}`,
	3: `{"buying_price":200.2222}`,
	4: `{"buying_price":200.9999}`,
	5: `{"buying_price":200.99}`,
}

// Output

type buyWagerOutput struct {
	Code        int
	BuyingPrice float64
}

var buyWagerOutputs = map[int]buyWagerOutput{
	0: {
		Code: 422,
	},
	1: {
		Code: 400,
	},
	2: {
		Code: 400,
	},
	3: {
		Code:        201,
		BuyingPrice: 200.22,
	},
	4: {
		Code:        201,
		BuyingPrice: 201,
	},
	5: {
		Code:        201,
		BuyingPrice: 200.99,
	},
}

func makeBuyWagerInput(r int) FakeContext {
	id := "1"
	if r == 2 {
		id = "999"
	} else if r == 1 {
		id = "2"
	}

	input := FakeContext{
		Method: "POST",
		Body:   buyWagerInputs[r],
		Param: map[string]string{
			"wager_id": id,
		},
	}
	return input
}

func checkBuyWager(b *testing.B) {
	r := rand.Intn(6)
	input := makeBuyWagerInput(r)
	ctx, rec := createEchoContext(input)
	_ = BuyWager(ctx)

	//if rec.Code != buyWagerOutputs[r].Code {
	//	b.Errorf("Expected %d, got %d - %s - %d", buyWagerOutputs[r].Code, rec.Code, rec.Body, r)
	//}

	out := buyWagerOutputs[r]
	assert.Equal(b, out.Code, rec.Code)
	if rec.Code == 201 {
		data := models.Transaction{}
		json.Unmarshal(rec.Body.Bytes(), &data)
		assert.Equal(b, out.BuyingPrice, data.BuyingPrice)
	}
}

func BenchmarkBuyWager(b *testing.B) {
	// Init Fake DB
	sqldb.InitFakeDB()
	sqldb.DB.Create(&models.Wager{Id: 1, SellingPrice: math.MaxFloat64, CurrentSellingPrice: math.MaxFloat64})
	sqldb.DB.Create(&models.Wager{Id: 2, SellingPrice: 1000, CurrentSellingPrice: 0})

	for i := 0; i < b.N; i++ {
		checkBuyWager(b)
	}
}

/* ----------------------------------------------------------------
						List Wager Benchmark
---------------------------------------------------------------- */

// Input

type listInput struct {
	Page  string
	Limit string
}

var listWagerInputs = map[int]listInput{
	0: {
		Page:  "a",
		Limit: "",
	},
	1: {
		Page:  "",
		Limit: "b",
	},
	2: {
		Page:  "",
		Limit: "",
	},
	3: {
		Page:  "100",
		Limit: "100",
	},
	4: {
		Page:  "1",
		Limit: "10",
	},
	5: {
		Page:  "4",
		Limit: "20",
	},
}

// Output

var listWagerOutputs = map[int]int{
	0: 400,
	1: 400,
	2: 200,
	3: 200,
	4: 200,
	5: 200,
}

func makeListInput(r int) FakeContext {
	input := FakeContext{
		Method: "GET",
		QueryParam: map[string]string{
			"page":  listWagerInputs[r].Page,
			"limit": listWagerInputs[r].Limit,
		},
	}
	return input
}

func checkListWager(b *testing.B) {
	r := rand.Intn(6)
	input := makeListInput(r)
	ctx, rec := createEchoContext(input)
	_ = ListWager(ctx)

	if rec.Code != listWagerOutputs[r] {
		b.Errorf("Expected %d, got %d - %s - %d", listWagerOutputs[r], rec.Code, rec.Body, r)
	}

	assert.Equal(b, listWagerOutputs[r], rec.Code)
}

func BenchmarkListWager(b *testing.B) {
	// Init Fake DB
	sqldb.InitFakeDB()
	for i := 0; i < 1000; i++ {
		sqldb.DB.Create(&models.Wager{})
	}

	for i := 0; i < b.N; i++ {
		checkListWager(b)
	}
}

/* ----------------------------------------------------------------
							Test Setups
---------------------------------------------------------------- */

type FakeContext struct {
	Method     string
	Body       string
	Param      map[string]string
	QueryParam map[string]string
}

// Test Setups
func createEchoContext(f FakeContext) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := makeUrl(f)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	initParam2Ctx(f, c)
	return c, rec
}

func makeUrl(f FakeContext) *http.Request {
	q := make(url.Values)
	for k, v := range f.QueryParam {
		q.Set(k, v)
	}
	req := httptest.NewRequest(f.Method, "/?"+q.Encode(), strings.NewReader(f.Body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func initParam2Ctx(f FakeContext, ctx echo.Context) {
	for k, v := range f.Param {
		ctx.SetParamNames(k)
		ctx.SetParamValues(v)
	}
}
