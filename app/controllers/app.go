package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vivvo/go-sdk/utils"
	"github.com/revel/revel"
	"log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type SuccessfulAuthenticationDto struct {
	IdentityId string `json:"identityId,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
}

func (c App) Login(token string) revel.Result {
	successfulAuthenticationDto := SuccessfulAuthenticationDto{}

	utils.Resty(context.Background()).R().
		SetHeader("Accept", "application/json").
		SetHeader("Client-Id", "bf40e6ee-fcb3-480d-ac73-31b66d30d871").
		SetHeader("Client-Secret", "2524f14f-f136-40dc-aba2-14dc08e1f441").
		SetResult(&successfulAuthenticationDto).
		Get(fmt.Sprintf("https://eeze.io/api/v1/did-auth/challenge/%s/user", token))

	b, _ := json.Marshal(&successfulAuthenticationDto)

	log.Println(string(b))

	c.ViewArgs["firstName"] = successfulAuthenticationDto.FirstName
	c.ViewArgs["lastName"] = successfulAuthenticationDto.LastName
	return c.RenderTemplate("App/Success.html")
}
