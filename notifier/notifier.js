const express = require("express");

const app = express();
// server settings
app.set("PORT", process.env.PORT || 3000);

// MIDDLEWARES
app.use(express.json());

// ROUTES
// Sends http status 200 if server is running
app.get("/health", (req, res) => {
  res.status(200).send();
});

module.exports = { app };
