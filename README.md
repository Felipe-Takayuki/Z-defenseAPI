# Z-defenseAPI

A minha solução é um API que foi feita em GoLang utilizando algumas bibliotecas como Gin e Gorm. 

---

## Inicialização

para inicializar o projeto primeiro entre no diretório da API e rode os seguintes comandos tendo o Docker instalado na maquina caso não tenha faça  o download em [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/)

`**git clone https://github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC**`

`docker build --tag z` 

`docker run -p 8080:8080 zombie defense`  

## Como testar a API ?

Recomendo que utilize o aplicativo **postman** e importe este arquivo:

para dentro do mesmo

---

## **Rotas**

endereço usado: [`**http://localhost:8080](http://localhost:8080)/api`**       

`/catalogar : POST`

- Cadastro dos potenciais hospedeiro+ analise dos mesmos. As informações são passadas dvia `json`

```json
{ 
  "nome"   : //<string>,
  "idade"  : //<int>,
  "sexo"   : //<string>,
	"peso"   : //<float>,
  "altura" : //<float>,
  "tipo-sanguineo" : //<string>,
  "gosto-musical"  : //<string>,
  "qual-esporte" : //<string>,
  "jogo-preferido" : //<string>
}
```

`**/deletarHosp?id=<id> : DELETE**`

- Deleta potencial hospedeiro e analise pelo `id`
    
    

`**/atualizarHosp?id=<id> : PUT**`

- Atualiza potencial hospedeiro pelo `id` via `json` passando o ou os parâmetros que devem ser alterados.

`**/analiseHosp?id=<id> : GET**`

- Analisa o potencial hospedeiro buscando o pelo `id` e retorna o nível do hospedeiro como zumbi.

**`/analises: GET`**

- Mostra uma lista com todos potenciais hospedeiros analisados.

`**/hosps : GET`**    

- Mostra uma lista dos potenciais hospedeiros já cadastrados.

`**/filtrar?classificacao=<numero> : GET`** 

- Lista as analises de acordo com a classificação sendo elas representadas por números de 1 a 4 para facilitar o uso da API

**1 - Zumbi extremamente perigoso** 

**2 - Zumbi perigoso** 

**3 - Zumbi normal** 

**4 - Zumbi fraco**
