const express = require("express");
const bcrypt = require("bcrypt");
const Admin = require("../models/admin");

const router = express.Router();

router.post("/register", async (req, res) => {
  const registeredAdmin = Admin.findOne({ username: req.body.username });

  if (registeredAdmin) {
    return res
      .status(400)
      .json({ error: true, message: "Username is already registered" });
  }

  const salt = await bcyrpt.genSalt(10);
  const hashedPassword = await bcyrpt.hash(req.body.password, salt);

  const admin = await Admin.create({
    username: req.body.username,
    password: hashedPassword,
  });

  try {
    await admin.save();
    res.status(200).json({ username: admin.username });
  } catch (error) {
    res.status(400).json(error);
  }
});

router.post("/login", async (req, res) => {
  const admin = await Admin.findOne({ username: req.body.username });

  if (!admin) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const validPassword = await bcyrpt.compare(req.body.password, admin.password);

  if (!validPassword) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const token = jwt.sign({ _id: admin._id }, process.env.ADMIN_SECRET_TOKEN);
  res.header("Authorization", token);
  res.status(200).json({ token: token });
});

module.exports = router;
