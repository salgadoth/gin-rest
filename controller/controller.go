package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salgadoth/gin-rest-api/database"
	"github.com/salgadoth/gin-rest-api/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(200, alunos)
}

func ExibeAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Aluno não encontrado",
		})

		return
	}

	c.JSON(200, aluno)
}

func ExibeAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno

	cpf := c.Params.ByName("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Aluno não encontrado",
		})

		return
	}

	c.JSON(200, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + " tudo beleza?!",
	})
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"status": "Aluno removido com sucesso",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}
