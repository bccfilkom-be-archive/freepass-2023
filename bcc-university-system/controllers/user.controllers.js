const bcrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const User = require("../models/user");
const Class = require("../models/class");
const Course = require("../models/course");

exports.register = async (req, res) => {
  const emailExists = await User.findOne({ username: req.body.username });

  if (emailExists) {
    return res
      .status(400)
      .json({ error: true, message: "Username is already registered" });
  }

  const salt = await bcrypt.genSalt(10);
  const hashedPassword = await bcrypt.hash(req.body.password, salt);

  const user = await User.create({
    username: req.body.username,
    password: hashedPassword,
    fullName: req.body.fullName,
    semester: req.body.semester,
  });

  try {
    await user.save();
    res.status(201).json({ username: user.username });
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
};

exports.login = async (req, res) => {
  const user = await User.findOne({ username: req.body.username });

  if (!user) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const validPassword = await bcrypt.compare(req.body.password, user.password);

  if (!validPassword) {
    return res
      .status(400)
      .json({ error: true, message: "Wrong username or password" });
  }

  const token = jwt.sign({ _id: user._id }, process.env.SECRET_TOKEN);
  res.header("Authorization", token);
  res.status(200).json({ token: token });
};

exports.getUser = async (req, res) => {
  let user;

  if (req.params.id.match(/^[0-9a-fA-F]{24}$/)) {
    user = await User.findById(req.params.id);
  } else {
    return res
      .status(404)
      .json({ error: true, message: "User does not exist" });
  }

  if (!user) {
    return res
      .status(404)
      .json({ error: true, message: "User does not exist" });
  }

  res.status(200).json(user);
};

exports.updateUser = async (req, res) => {
  if (req.body.username) {
    await User.findByIdAndUpdate(req.params.id, {
      username: req.body.username,
    }).catch((err) => console.error(err));
  } else if (req.body.password) {
    const salt = await bcrypt.genSalt(10);
    const hashedPassword = await bcrypt.hash(req.body.password, salt);
    await User.findByIdAndUpdate(req.params.id, {
      password: hashedPassword,
    }).catch((err) => console.error(err));
  } else if (req.body.fullName) {
    await User.findByIdAndUpdate(req.params.id, {
      fullName: req.body.fullName,
    }).catch((err) => console.error(err));
  } else if (req.body.semester) {
    await User.findByIdAndUpdate(req.params.id, {
      semester: Number(req.body.semester),
    }).catch((err) => console.error(err));
  } else {
    return res.status(400).json({ error: true, message: "Field not found" });
  }

  res.status(200).json({ message: "successfully edited" });
};

exports.getClassByUser = async (req, res) => {
  let user;

  if (req.params.id.match(/^[0-9a-fA-F]{24}$/)) {
    user = await User.findById(req.params.id).populate({
      path: "_class",
      select: "-_student",
      populate: { path: "_course", select: "name sks -_id" },
    });
  } else {
    return res
      .status(404)
      .json({ error: true, message: "User does not exist" });
  }

  if (!user) {
    res.status(404).json({ error: true, message: "User does not exist" });
  }

  res.status(200).json(user._class);
};

exports.addClassByUser = async (req, res) => {
  let classVar;

  if (req.body.classId.match(/^[0-9a-fA-F]{24}$/)) {
    classVar = await Class.findById(req.body.classId);
  } else {
    return res
      .status(404)
      .json({ error: true, message: "Class does not exist" });
  }

  if (!classVar) {
    return res
      .status(404)
      .json({ error: true, message: "Class does not exist" });
  }

  // const user = await User.findById(req.params.id);
  // const checkCourse = await Course.find({
  //   _class: { $in: user._class },
  // });

  // console.log(checkCourse);

  try {
    await User.updateOne(
      { _id: req.params.id },
      {
        $addToSet: { _class: req.body.classId },
      }
    );
    await Class.updateOne(
      { _id: req.body.classId },
      {
        $addToSet: { _student: req.params.id },
      }
    );

    res.status(200).json({ message: "successfully added new class" });
  } catch (error) {
    res.status(400).json({
      error: true,
      message: error,
    });
  }
};

exports.dropClass = async (req, res) => {
  let classVar;

  if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
    classVar = await Class.findOne({ _id: req.params.classId });
  } else {
    return res
      .status(404)
      .json({ error: true, message: "Class does not exist" });
  }

  if (!classVar) {
    return res
      .status(404)
      .json({ error: true, message: "Class does not exist" });
  }

  try {
    await User.findByIdAndUpdate(req.params.userId, {
      $pull: { _class: req.body.classId },
    });
    await Class.findByIdAndUpdate(req.body.classId, {
      $pull: { _student: req.params.userId },
    });

    res.status(200).json({ message: "successfully dropped a class" });
  } catch (error) {
    res.status(400).json({
      error: true,
      message: error,
    });
  }
};

exports.getStudentByClass = async (req, res) => {
  let classVar;

  if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
    classVar = await Class.findById(req.params.classId)
      .select("-_course")
      .populate("_student", "fullName");
  } else {
    return res
      .status(404)
      .json({ error: true, message: "Class does not exist" });
  }

  if (!classVar) {
    res.status(404).json({ error: true, message: "Class does not exist" });
  }

  res.status(200).json(classVar);
};
