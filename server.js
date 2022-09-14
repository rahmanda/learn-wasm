const fs = require('fs')
const http = require('http')

http.createServer((req, res) => {
  fs.readFile(__dirname + req.url, (err, data) => {
    if (err) {
      res.writeHead(404)
      res.end(JSON.stringify(err))
      return
    }
    if (req.url.includes('/main.wasm')) {
      res.setHeader('Content-Type', 'application/wasm')
    }
    res.writeHead(200)

    res.end(data)
  })
}).listen(8000)
