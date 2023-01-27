const express = require("express");

const router = express.Router();

router.use("/users", require('./user.routes'));
router.use("/admin", require('./admin.routes'));
router.use("/courses", require('./course.routes'));

module.exports = router;
