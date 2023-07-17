
const serverless = require("serverless-http")
console.log("Hello world")
const app = express()
app.use(express.json())
app.get("/health", (_, res) => {
    res.send()
})
module.exports.handler = serverless(app)