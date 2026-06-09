var carros = '[' +
'{"id": 1, "marca": "Honda", "modelo": "HRV"},' +
'{"id": 2, "marca": "Volkswagen", "modelo": "Golf"},' +
'{"id": 3, "marca": "Fiat", "modelo": "Toro"},' +
'{"id": 4, "marca": "Ford", "modelo": "Ka"}' +
']'

var http = require('http')
var server = http.createServer(function(request, response) {
    response.setHeader('Access-Control-Allow-Origin', '*')
    response.writeHead(200, {"Content-Type": "text/html"})
    response.write(carros)
    response.end()
})

server.listen(3000,  () => {
    console.log('Servidor rodando em http://localhost:3000');
})