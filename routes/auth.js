const express = require('express')
const router = express.Router()
const controller = require('../controller/auth')

router.post('/signup', controller.Signup);

module.exports = router
