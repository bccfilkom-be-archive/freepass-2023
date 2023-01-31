const Class = require("../models/class");
const Course = require("../models/course");

exports.getCourse = async (req, res) => {
  const courses = await Course.find({}).select("-_class").sort({ name: 1 });

  try {
    res.status(200).json(courses);
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
};

exports.addCourse = async (req, res) => {
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
};

exports.updateCourse = async (req, res) => {
  let courseFound;

  if (req.params.courseId.match(/^[0-9a-fA-F]{24}$/)) {
    courseFound = await Course.findOne({ _id: req.params.courseId });
  } else {
    return res
      .status(404)
      .json({ error: true, message: "Course is not found" });
  }

  if (!courseFound) {
    return res
      .status(404)
      .json({ error: true, message: "Course is not found" });
  }

  if (req.body.name) {
    await Course.findByIdAndUpdate(req.params.courseId, {
      name: req.body.name,
    });
  } else if (req.body.sks) {
    await Course.findByIdAndUpdate(req.params.courseId, {
      sks: req.body.sks,
    });
  } else {
    res.status(400).json({ error: true, message: "Nothing changed" });
  }

  res.status(200).json({ message: "successfully edited" });
};

exports.deleteCourse = async (req, res) => {
  let course;

  if (req.params.courseId.match(/^[0-9a-fA-F]{24}$/)) {
    course = await Course.findOne({ _id: req.params.courseId });
  } else {
    return res
      .status(404)
      .json({ error: true, message: "Course is not found" });
  }

  if (!course) {
    return res
      .status(404)
      .json({ error: true, message: "Course is not found" });
  }

  try {
    await Class.deleteMany({ _course: req.params.courseId });
    await Course.deleteOne({ _id: req.params.courseId });
    res.status(200).json({ message: "successfully deleted" });
  } catch (error) {
    res.status(400).json({ error: true, message: error });
  }
};
