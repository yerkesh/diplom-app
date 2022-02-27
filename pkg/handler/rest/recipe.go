package rest

import "net/http"

type RecipeRest struct {

}

func NewRecipeRest() *RecipeRest {
	return &RecipeRest{}
}

func (r *RecipeRest) Hello(res http.ResponseWriter, req *http.Request) {

}