package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type AngularController struct {
	Application
}

func (c AngularController) IndexAngular() revel.Result {
	fmt.Println("request")
	return c.Render()
}
