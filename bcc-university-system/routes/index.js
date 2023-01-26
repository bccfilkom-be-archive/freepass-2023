const express = require("express");

const router = express.Router();

router.use("/users");
router.use("/admin");
router.use("/course");

module.exports = router;
