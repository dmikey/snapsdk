// Example HTTP receiver for DogsAppSDK
// Run with: node receiver.js

const http = require('http');

// Mock dog database
const dogs = [
  { id: 1, name: "Buddy", breed: "Golden Retriever", age: 3 },
  { id: 2, name: "Max", breed: "German Shepherd", age: 5 },
  { id: 3, name: "Charlie", breed: "Beagle", age: 2 },
];

const server = http.createServer((req, res) => {
  if (req.method !== 'POST') {
    res.writeHead(405);
    res.end('Method not allowed');
    return;
  }

  let body = '';
  req.on('data', chunk => body += chunk);
  req.on('end', () => {
    try {
      const { method, arguments: args } = JSON.parse(body);
      console.log(`Received call: ${method}`, args);

      let result;
      switch (method) {
        case 'listDogs':
          result = dogs;
          break;
        case 'getDog':
          result = dogs.find(d => d.id === args.id) || { error: "Dog not found" };
          break;
        default:
          res.writeHead(400);
          res.end(JSON.stringify({ error: `Unknown method: ${method}` }));
          return;
      }

      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify(result, null, 2));
    } catch (err) {
      res.writeHead(400);
      res.end(JSON.stringify({ error: err.message }));
    }
  });
});

server.listen(8080, () => {
  console.log('DogsApp receiver running on http://localhost:8080/api');
});
