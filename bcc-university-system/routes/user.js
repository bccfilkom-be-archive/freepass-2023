const express = require("express");

const router = express.Router();

router.post("/register");
router.post("/login");
// need authorization middleware
router.get("/:id/account");
router.post("/:id/account/edit");
router.route("/:id/classes").get().post();
router.delete("/:userID/classes/:classID");
router.get("/:userID/classes/:classID/students");
