const express = require('express');
const app = express();
const PORT = 3000;

app.get('/', (req, res) => {
  res.json({
    status: 'success',
    message: 'Hello from Node.js Express running inside a Docker container!'
  });
});

// Bind to 0.0.0.0 so the container is accessible from outside
app.listen(PORT, '0.0.0.0', () => {
  console.log(`Server is running and listening on http://0.0.0.0:${PORT}`);
});
