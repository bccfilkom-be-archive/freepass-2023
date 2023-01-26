const express = require("express");

const router = express.Router();

router.use("/users", require('./user.routes'));
// router.use("/admin");
// router.use("/course");

module.exports = router;
