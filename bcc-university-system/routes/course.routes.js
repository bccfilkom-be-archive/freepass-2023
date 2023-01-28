const express = require("express");
const User = require("../models/user");
const Class = require("../models/class");
const Course = require("../models/course");

const router = express.Router();

router.post("/", async (req, res) => {
  const course = await Course.create({
    name: req.body.name,
    sks: req.body.sks,
  });

  try {
    await course.save();
    res.status(201).json(course);
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
});

router
  .route("/:courseId/classes/")
  .get(async (req, res) => {
    const classes = await Class.find({})
      .sort({ _course: 1 })
      .select("-_student")
      .populate("_course", "name sks -_id");

    try {
      res.status(200).json(classes);
    } catch (error) {
      res.status(400).json({ error: true, message: error });
    }
  })
  .post(async (req, res) => {
    let classFound;

    if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
      classFound = await Class.findOne({ name: req.body.name });
    } else {
      return res
        .status(400)
        .json({ error: true, message: "Class is already created" });
    }

    if (classFound) {
      return res
        .status(400)
        .json({ error: true, message: "Class is already created" });
    }

    const classVar = await Class.create({
      name: req.body.name,
      _course: req.params.courseId,
    });

    try {
      await classVar.save();
      await Course.findByIdAndUpdate(req.params.courseId, {
        $addToSet: { _class: classVar._id },
      });
      res.status(201).json(classVar);
    } catch (error) {
      res.status(400).json({ error: true, message: error });
    }
  });

router
  .route("/:courseId/classes/:classId")
  .patch(async (req, res) => {
    let classFound;

    if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
      classFound = await Class.findOne({ _id: req.params.classId });
    } else {
      return res
        .status(404)
        .json({ error: true, message: "Class is not found" });
    }

    if (!classFound) {
      return res
        .status(400)
        .json({ error: true, message: "Class is not found" });
    }

    if (req.body.name) {
      await Class.findByIdAndUpdate(req.params.classId, {
        name: req.body.name,
      });
      res.status(200).json({ message: "successfully edited" });
    } else {
      res.status(400).json({ error: true, message: "Nothing changed" });
    }
  })
  .delete(async (req, res) => {
    let classVar;

    if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
      classVar = await Class.findOne({ _id: req.params.classId });
    } else {
      return res
        .status(404)
        .json({ error: true, message: "Class is not found" });
    }

    if (!classVar) {
      return res.status(404).json({
        error: true,
        message: "Class is not found",
      });
    }

    try {
      await User.findOneAndUpdate(
        { _class: req.params.classId },
        { $pull: { _class: req.params.classId } }
      );
      await Course.findOneAndUpdate(
        { _class: req.params.classId },
        { $pull: { _class: req.params.classId } }
      );
      await Class.deleteOne({ _id: req.params.classId });
      res.status(200).json({ message: "successfully deleted" });
    } catch (error) {
      res.status(400).json({
        error: true,
        message: error,
      });
    }
  });

router.post("/:courseId/classes/:classId/students", async (req, res) => {
  let classFound;

  if (req.params.classId.match(/^[0-9a-fA-F]{24}$/)) {
    classFound = await Class.findById(req.params.classId);
  } else {
    return res.status(404).json({ error: true, message: "Class is not found" });
  }

  if (!classFound) {
    return res.status(404).json({ error: true, message: "Class is not found" });
  }

  try {
    await User.updateOne(
      { _id: req.body.userId },
      {
        $addToSet: { _class: req.params.classId },
      }
    );
    await Class.updateOne(
      { _id: req.params.classId },
      {
        $addToSet: { _student: req.body.userId },
      }
    );

    res.status(200).json({ message: "successfully added new user" });
  } catch (error) {
    res.status(400).json({
      error: true,
      message: error,
    });
  }
});

router.patch(
  "/:courseId/classes/:classId/students/:studentId",
  async (req, res) => {
    try {
      await User.findByIdAndUpdate(req.params.studentId, {
        $pull: { _class: req.body.classId },
      });
      await Class.findByIdAndUpdate(req.params.classId, {
        $pull: { _student: req.params.studentId },
      });

      res.status(200).json({ message: "successfully deleted an user" });
    } catch (error) {
      res.status(400).json({
        error: true,
        message: error,
      });
    }
  }
);

module.exports = router;
