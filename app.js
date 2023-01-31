const express = require('express')
const app = express()
const db = require('./db/index')
const bodyparser = require("body-parser")
const port = 3000

app.use(express.json())
app.use(bodyparser.urlencoded({ extended : true}))

const adminEndpoint = require('./routes/admin')
const userEndpoint = require('./routes/users')
const authEndpoint = require('./routes/auth')

app.use('/user', userEndpoint)
app.use('/admin', adminEndpoint)
app.use('/auth', authEndpoint)

app.listen(port,()=>{
    console.log(`port running in port ${port}`)
})

module.exports = app
