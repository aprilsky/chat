package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitIntercept)

}
func InitIntercept() {
	revel.InterceptMethod((*Application).checkUser, revel.BEFORE)

}
