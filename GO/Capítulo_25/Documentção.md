# Documentação em GO
---

## Introdução
A documentação na linguagem GO é muito fácil de ser feita e atualizada, pois os desenvolvedores da linguagem para facilitar a organização da documentação, automatizaram a criação da mesma.

---
## Onde encontrar
A documetação da linguagem GO pode ser encontrada em vários lugares. Os 3 principais são:  
- Site [go.dev](https://pkg.go.dev)
- Terminal (go doc)
    ```terminal
    go doc
    ```
- Terminal/HTTP (godoc)
    ```terminal
    godoc -http=:[porta_desejada]
    ```

---
## Escrevendo a documentação
Para documentar pacotes, tipos, variáveis ou constantes, escreva um comentário diretamente antes dos mesmos, **sem linhas em branco**.  
A primeira linha escrita no arquivo principal do package aparece no package list.  

Em casos de uma documentação extensa, que atrapalharia a leitura do código, pode ser criado um arquivo chamado **doc.go**.  

Nele será escrito, toda a documentação do código, e na última linha será colocado o package a qual ele pertence fora do comentário, sem linha em branco após o comentário da documentação. Exemplo: [doc.go do package fmt](https://cs.opensource.google/go/go/+/refs/tags/go1.26.3:src/fmt/doc.go).

---
## Upload de packages
Para criar packages e fazer sua documentação aparecer no site go.dev, é só criar o package com a documentação e dar upload no GitHub, automaticamente com o link do repositório será mostrado na aba de packages do site.