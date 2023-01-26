require('dotenv').config()
const express = require("express");
const cors = require('cors')
const db = require("./util/database");
const routes = require('./routes/index')

const app = express();
const port = process.env.PORT

app.use(cors)
app.use(express.urlencoded({extended: true}))

app.use('/api', routes)

app.listen(port, () => {
    console.log(`Server is listening in port ${port}...`);
})
