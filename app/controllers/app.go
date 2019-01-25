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
		SetHeader("Client-Id", "947493fc-5fc8-4c9e-ab84-d7f03c95968c").
		SetHeader("Client-Secret", "a879248a-93ed-4a75-9f3d-35bc2eae2590").
		SetResult(&successfulAuthenticationDto).
		Get(fmt.Sprintf("https://eeze.io/api/v1/did-auth/challenge/%s/user", token))

	b, _ := json.Marshal(&successfulAuthenticationDto)

	log.Println(string(b))

	c.ViewArgs["firstName"] = successfulAuthenticationDto.FirstName
	c.ViewArgs["lastName"] = successfulAuthenticationDto.LastName
	return c.RenderTemplate("App/Success.html")
}
