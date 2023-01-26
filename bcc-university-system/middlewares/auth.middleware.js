require('dotenv').config()
const jwt = require('jsonwebtoken')

module.exports = (req, res, next) => {
    if (req.header("Authorization")) {
        let decoded
        try {
            decoded = jwt.verify(req.header("Authorization"), process.env.SECRET_TOKEN); 
            next()
        } catch (error) {
            res
              .status(400)
              .json({ error: true, message: "Wrong Token" });
        }
    } else {
        res.status(401).json({ error: true, message: "Please login to access this resource" });
    }
}