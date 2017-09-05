const http = require('http');

const requestHandler = (request, response) => {  
  const date = new Date();
  var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
  response.end(date.toLocaleDateString("en-US", options));
}

const server = http.createServer(requestHandler)

server.listen(3000, (err) => {  
  if (err) {
    response.write(500);
    response.end();
  }
})