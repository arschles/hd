package actions

import (
	"github.com/arschles/hd/models"
)

func (as *ActionSuite) Test_TodosResource_List() {
	r := as.Require()
	db := as.DB
	todo := &models.Todo{Name: "Test1"}
	r.NoError(db.Create(todo))
	res := as.HTML("/todos").Get()
	r.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TodosResource_Show() {
	r := as.Require()
	db := as.DB
	todo := &models.Todo{Name: "Test1"}
	r.NoError(db.Create(todo))
	res := as.HTML("/todos/%s", todo.ID.String()).Get()
	r.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TodosResource_New() {
	// just make sure that the "new todo" form shows up
	r := as.Require()
	res := as.HTML("/todos/new").Get()
	r.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TodosResource_Create() {
	r := as.Require()
	db := as.DB
	todo := &models.Todo{Name: "Hello World!"}
	/*res := */ as.HTML("/todos").Post(todo)
	// r.Equal(201, res.Code)
	todos := &models.Todos{}
	r.NoError(db.All(todos))
	r.Equal(1, len(*todos))
}

func (as *ActionSuite) Test_TodosResource_Edit() {
	as.T().Log("Not Implemented. Left as an exercise for the reader!")
}

func (as *ActionSuite) Test_TodosResource_Update() {
	as.T().Log("Not Implemented. Left as an exercise for the reader!")
}

func (as *ActionSuite) Test_TodosResource_Destroy() {
	as.T().Log("Not Implemented. Left as an exercise for the reader!")
}
