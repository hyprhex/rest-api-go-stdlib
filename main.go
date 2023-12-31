package main

import (
	"net/http"
	"regexp"

	"github.com/hyprhex/rest-api-go-stdlib/pkg/recipes"
)

type homeHandler struct{}
type RecipesHandler struct{
	store recipeStore
}
type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List()(map[string]recipes.Recipe, error)
	Remove(name string) error
}

var (
    RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
    RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
    case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
        h.CreateRecipe(w, r)
        return
    case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
        h.ListRecipes(w, r)
        return
    case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
        h.GetRecipe(w, r)
        return
    case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
        h.UpdateRecipe(w, r)
        return
    case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
        h.DeleteRecipe(w, r)
        return
    default:
        return
    }
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

func main() {
	// Create a new request multiplexer
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &RecipesHandler{})
	mux.Handle("/recipes/", &RecipesHandler{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}






