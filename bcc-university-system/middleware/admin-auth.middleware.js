require("dotenv").config();
const jwt = require("jsonwebtoken");

module.exports = (req, res, next) => {
  if (req.header("Authorization")) {
    let decoded;
    try {
      decoded = jwt.verify(
        req.header("Authorization"),
        process.env.ADMIN_SECRET_TOKEN
      );
      next();
    } catch (error) {
      res
        .status(400)
        .json({ error: true, message: "Only admin can access this resource" });
    }
  } else {
    res
      .status(403)
      .json({ error: true, message: "Only admin can access this resource" });
  }
};
