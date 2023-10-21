# Z-defenseAPI

A minha solução para a etapa bônus do Coder Challenge é uma API que basicamente analisa potenciais hospedeiros para zumbis e ela foi feita em GoLang. 

---

## Inicialização

para inicializar o projeto primeiro entre no diretório da API e rode os seguintes comandos tendo o Docker instalado na maquina caso não tenha faça  o download em [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/)

**`git clone https://github.com/Felipe-Takayuki/Z-defenseAPI.git`**

`cd Z-defenseAPI`

`docker build --tag z-defense .` 

`docker run -p 8080:8080 z-defense`  

---

## Como testar a API ?

Recomendo que utilize o aplicativo **postman** e importe este arquivo:
"z-defenderAPI.postman_collection.json" que esta dentro desse repositório
para dentro do app caso não tenha baixe em: https://www.postman.com/downloads/

---

### video mostrando esses passos e a API funcionando: 
 **https://youtu.be/iLH7_SQOi2Y?si=gtpXXxU2QlSH-mHF** [ASSISTA]


--- 
## **Rotas**

endereço usado: **`http://localhost:8080/api`**

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
  "pratica-esporte" : //<esporte string>,
  "jogo-preferido" : //<string>
}
```

`/deletarHosp?id=<id> : DELETE`

- Deleta potencial hospedeiro e analise pelo `id`
    
    

`/atualizarHosp?id=<id> : PUT`

- Atualiza potencial hospedeiro pelo `id` via `json` passando o ou os parâmetros que devem ser alterados.

`/analiseHosp?id=<id> : GET`

- Analisa o potencial hospedeiro e analise buscando o pelo `id` e retorna o nível do hospedeiro como zumbi.

**`/analises: GET`**

- Mostra uma lista com todos potenciais hospedeiros analisados.

`/hosps : GET`    

- Mostra uma lista dos potenciais hospedeiros já cadastrados.

`/filtrar?classificacao=<numero> : GET`

- Lista as analises de acordo com a classificação sendo elas representadas por números de 1 a 4 para facilitar o uso da API

**1 - Zumbi extremamente perigoso** 

**2 - Zumbi perigoso** 

**3 - Zumbi normal** 

**4 - Zumbi fraco**
