package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/salgadoth/gin-rest-api/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/alunos/:id", controller.ExibeAluno)
	r.GET("/alunos/cpf/:cpf", controller.ExibeAlunoPorCpf)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos", controller.CriaAluno)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.Run()
}
