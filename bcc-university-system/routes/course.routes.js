const express = require("express");
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

router.post("/:courseId/classes/", async (req, res) => {
  const classVar = await Class.create({
    name: req.body.name,
    _course: req.params.courseId
  });

  try {
    await classVar.save();
    await Course.findByIdAndUpdate(req.params.courseId, {$addToSet: {_class: classVar._id}})
    res.status(201).json(classVar);
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
});

module.exports = router;
