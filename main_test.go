package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/salgadoth/gin-rest-api/controller"
	"github.com/salgadoth/gin-rest-api/database"
	"github.com/salgadoth/gin-rest-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678900", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
}

func TestVarificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupRotas()
	r.GET("/:nome", controller.Saudacao)

	req, _ := http.NewRequest("GET", "/gui", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E ai gui tudo beleza?!"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockDaResposta, string(responseBody), "deveriam ser iguais")
}

func TestListarTodosAlunos(t *testing.T) {
	database.ConectaComDB()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotas()
	r.GET("/alunos", controller.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaComDB()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotas()
	r.GET("/alunos/cpf/:cpf", controller.ExibeAlunoPorCpf)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678900", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComDB()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotas()
	r.GET("/alunos/:id", controller.ExibeAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os dados devem ser iguais")
	assert.Equal(t, "12345678900", alunoMock.CPF, "Os dados devem ser iguais")
	assert.Equal(t, "123456789", alunoMock.RG, "Os dados devem ser iguais")
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComDB()

	CriaAlunoMock()

	r := SetupRotas()
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditaAluno(t *testing.T) {
	database.ConectaComDB()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "98745678900", RG: "987456789"}
	valorJson, _ := json.Marshal(aluno)

	r := SetupRotas()
	r.PATCH("/alunos/:id", controller.EditaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, aluno.Nome, alunoMock.Nome, "Os dados devem ser iguais")
	assert.Equal(t, aluno.CPF, alunoMock.CPF, "Os dados devem ser iguais")
	assert.Equal(t, aluno.RG, alunoMock.RG, "Os dados devem ser iguais")
}
