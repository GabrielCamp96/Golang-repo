package routes

import (
	"crud/db"
	"crud/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição! " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u := models.NewUsuario()
	if erro = json.Unmarshal(request, &u); erro != nil {
		w.Write([]byte("Falha ao parsear objeto usuario! " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values(?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement! " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(u.Nome, u.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement! " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// id, erro := insercao.LastInsertId()
	// if erro != nil {
	// 	w.Write([]byte("Erro ao obter o id inserido! " + erro.Error()))
	// 	return
	// }

	// w.Write([]byte(fmt.Sprintf("Usuario criado com sucesso! ID: %d", id)))
	w.WriteHeader(http.StatusNoContent)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuarios: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()

	users := models.NewSliceUsuarios()

	for result.Next() {
		user := models.NewUsuario()
		if erro := result.Scan(&user.ID, &user.Nome, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuários: " + erro.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, erro := db.Query("select * from usuarios u where u.id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuarios: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()

	user := models.NewUsuario()

	if result.Next() {
		if erro := result.Scan(&user.ID, &user.Nome, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuários: " + erro.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := models.NewUsuario()
	if erro := json.Unmarshal(request, &user); erro != nil {
		w.Write([]byte("Erro ao converter JSON para struct: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	if user.Nome != "" && user.Email != "" {
		statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
		if erro != nil {
			w.Write([]byte("Erro ao criar o statement: " + erro.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer statement.Close()

		if _, erro := statement.Exec(user.Nome, user.Email, ID); erro != nil {
			w.Write([]byte("Erro ao atualizar usuário! " + erro.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		if user.Nome != "" {
			statement, erro := db.Prepare("update usuarios set nome = ? where id = ?")
			if erro != nil {
				w.Write([]byte("Erro ao criar o statement: " + erro.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer statement.Close()

			if _, erro := statement.Exec(user.Nome, ID); erro != nil {
				w.Write([]byte("Erro ao atualizar usuário! " + erro.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		} else {
			statement, erro := db.Prepare("update usuarios set email = ? where id = ?")
			if erro != nil {
				w.Write([]byte("Erro ao criar o statement: " + erro.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer statement.Close()

			if _, erro := statement.Exec(user.Email, ID); erro != nil {
				w.Write([]byte("Erro ao atualizar usuário! " + erro.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios u where u.id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário: " + erro.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
