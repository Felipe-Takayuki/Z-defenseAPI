# sistema-de-defesa-dsinCC

## **Sobre o projeto**

A minha solução para a etapa bônus(problemas 1 e 2) é uma API feita usando **`GoLang` + `Sqlite`** através do **[GORM](https://gorm.io/index.html)** 

---

## **Instalação**

### Requisitos:

- **`Golang` [](https://go.dev/doc/install)instalado: [https://go.dev/doc/install](https://go.dev/doc/install) ;**
- **`gcc`** instalado;

 para a instalação do `**gcc**` a forma mais simples foi através do choco usando o seguinte comando no terminal com admin:

`**choco install mingw -y`** 

**instalação do choco:** [https://chocolatey.org/install](https://chocolatey.org/install) 

---

## Inicializar API

comandos:

1. **`git clone  [https://github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC.git](https://github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC.git)`** 
2. **`git cd [sistema-de-defesa-dsinCC](https://github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC)`**
3. **`go mod download`** 
4. **`go run main.go`**

---

## **Rotas:**

endereço usado: [`**http://localhost:8080](http://localhost:8080)/api`**       

`/catalogar : POST`

- Cadastro dos potenciais hospedeiro ao banco de dados + analise dos mesmos. As informações são passadas via `json`

```json
{ 
  "nome"   : //<string>,
  "idade"  : //<int>,
  "sexo"   : //<string>,
	"peso"   : //<float>,
  "altura" : //<float>,
  "tipoSanguineo" : //<string>,
  "gostoMusical"  : //<string>,
  "esporte": //<string>,
  "jogoPreferido" : //<string>
}
```

`**/deletarHosp?id=<id> : DELETE`

- Deleta potencial hospedeiro e analise pelo `id`
    
    

**`/atualizarHosp?id=<id> : PUT`**

- Atualiza potencial hospedeiro pelo id via `json` passando os parâmetros que devem ser alterados.

**`/analiseHosp?id=<id> : GET`**

- Analisa o potencial hospedeiro buscando o pelo `id` e retorna o nível provável do hospedeiro como zumbi.

**`/analises: GET`**

- Mostra uma lista com todos hospedeiros analisados.

**`/hosps : GET`**    

- Mostra uma lista dos potenciais hospedeiros já cadastrados.

`**/filtrar?Classificacao=<numero> : GET`** 

- Lista as analises de acordo com a classificação sendo elas representadas por números de 1 a 4 para facilitar o uso da API

**1 - Zumbi extremamente perigoso** 

**2 - Zumbi perigoso** 

**3 - Zumbi normal** 

**4 - Zumbi fraco**
