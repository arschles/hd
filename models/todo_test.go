package models_test

import "testing"

func Test_Todo(t *testing.T) {
	t.Logf("If the todo has some special logic on it, then you can test it here")
}

func (m *ModelSuite) Test_Todo() {
	m.T().Logf("Or better yet, test it here!")
}
