const express = require("express");
const bcyrpt = require("bcrypt");
const jwt = require("jsonwebtoken");
const User = require("../models/user");
const isLogin = require('../middlewares/auth.middleware')

const router = express.Router();

router.get("/", (req, res) => {
  res.send("test");
});

router.post("/register", async (req, res) => {
  const emailExists = await User.findOne({ username: req.body.username });

  if (emailExists) {
    return res
      .status(400)
      .json({ error: true, message: "Username is already registered" });
  }

  const salt = await bcyrpt.genSalt(10);
  const hashedPassword = await bcyrpt.hash(req.body.password, salt);

  const user = await User.create({
    username: req.body.username,
    password: hashedPassword,
    fullName: req.body.fullName,
  });

  try {
    await user.save();
    res.status(201).json({ username: user.username });
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
});

router.post("/login", async (req, res) => {
  const user = await User.findOne({ username: req.body.username });

  if (!user) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const validPassword = await bcyrpt.compare(req.body.password, user.password);

  if (!validPassword) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const token = jwt.sign({ _id: user._id }, process.env.SECRET_TOKEN);
  res.header("Authorization", token);
  res.status(200).json({ token: token });
});
// need authorization middleware
router.get("/:id/account", isLogin, async (req, res) => {
  const user = await User.findById(req.params.id);

  if (!user) {
    return res
      .status(404)
      .json({ error: true, message: "User does not exist" });
  }

  res.status(200).json(user)
});
// router.post("/:id/account/edit");
// router.route("/:id/classes").get().post();
// router.delete("/:userID/classes/:classID");
// router.get("/:userID/classes/:classID/students");

module.exports = router;
