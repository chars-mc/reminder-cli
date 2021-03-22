const express = require("express");
const { notify } = require("node-notifier");

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

// Sends an OS notification
app.post("/notify", (req, res) => {
  const notifyContent = {
    title: req.body.title || "Unknown title",
    message: req.body.message || "Unknown message",
    sound: true,
    wait: true,
    reply: true,
    closeLabel: "Completed?",
    timeout: 15,
  };

  notify(notifyContent, (err, response, reply) => {
    // sendding the notify reply to the response
    res.send(reply);
  });
});

module.exports = { app };
