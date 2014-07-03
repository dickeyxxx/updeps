package api

import (
	"github.com/dickeyxxx/updeps/category"

	"github.com/martini-contrib/render"
)

func categoryCreate(category category.Category, categoryClient *category.Client) int {
	if _, err := categoryClient.UpsertCategory(&category); err != nil {
		panic(err)
	}
	return 201
}

func categoryList(r render.Render, category *category.Client) {
	categories, err := category.All()
	if err != nil {
		panic(err)
	}
	response := make([]map[string]interface{}, len(categories))
	for i := 0; i < len(categories); i++ {
		response[i] = map[string]interface{}{
			"name": categories[i].Name,
			//"top":  category.TopPackages(categories[i]),
		}
	}
	r.JSON(200, response)
}
