package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/url"
	"scooby-api/apps/gpt"
	"scooby-api/apps/ticker"
	"strings"
)

func main() {
	// ...
	e := echo.New()
	e.POST("/", func(c echo.Context) error {

		fmt.Println("Raw Request Body : %s", c.Request().Body)
		bodyData, err := io.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Raw Body: %s", bodyData)
		fmt.Println("String Body: %s", string(bodyData))

		queryData, err := url.ParseQuery(string(bodyData))
		if err != nil {
			fmt.Println(err)
		}

		if !queryData.Has("Body") {
			return c.String(http.StatusBadRequest, "Body not found")
		}
		msg := queryData.Get("Body")

		fmt.Println("Body message: ", msg)

		if strings.Contains(strings.ToLower(msg), "help") {
			return c.String(http.StatusOK, "Bem vindo ao Scooby chatbot.\n"+
				"Essas são as opções disponíveis:\n\n"+
				"*help*: Ajuda\n"+
				"*gpt*: Fazer uma pergunta para o chatGPT\n"+
				"*ticker*: Ver o valor de uma ação\n"+
				"*dolar*: Ver o valor do dólar\n")
		} else if strings.Contains(strings.ToLower(msg), "gpt") {
			resp, _ := gpt.GenerateGPTText(msg[4:])
			return c.String(http.StatusOK, resp)
		} else if strings.Contains(strings.ToLower(msg), "ticker") {
			tickerName := msg[7:]
			quoteSummary, _ := ticker.QuoteSummary(tickerName)
			return c.String(http.StatusOK, "Valor do Ticker é: "+quoteSummary.QuoteSummary.Result[0].Price.RegularMarketPrice.Fmt)
		} else if strings.Contains(strings.ToLower(msg), "dolar") {
			tickerName := "BRL=X"
			quoteSummary, _ := ticker.QuoteSummary(tickerName)
			return c.String(http.StatusOK, "Valor do Dolar é: "+quoteSummary.QuoteSummary.Result[0].Price.RegularMarketPrice.Fmt)
		}

		return c.String(http.StatusOK, "Opção inválida. Digite help para ver as opções disponíveis.")
	})
	e.Logger.Fatal(e.Start(":9000"))
}
