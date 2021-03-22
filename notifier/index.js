const { app } = require("./notifier");

(function main() {
  // starting the server
  app.listen(app.get("PORT"), () => {
    console.log(`server on port ${app.get("PORT")}`);
  });
})();
