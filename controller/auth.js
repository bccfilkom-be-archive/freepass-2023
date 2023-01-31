const data = require('../models/users')
const bcrypt = require('bcrypt')

module.exports = {
    Signup: async (req, res) => {
        const {nip, nama, password} = req.body
        const encryptedPassword = await bcrypt.hash(password, 10)
    
        const user = await data.insertMany({
            nip,nama,password : encryptedPassword
        })

        res.status(200).json({
            data : user[0].nama,
            "metadata" : "berhasil register"
        })
    }
  }