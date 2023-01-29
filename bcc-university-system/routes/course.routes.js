const express = require("express");
const courseController = require('../controllers/course.controllers')
const isAdmin = require("../middleware/admin-auth.middleware");

const router = express.Router();

router.use(isAdmin);

router
  .route("/")
  .get(courseController.getCourse)
  .post(courseController.addCourse);

router
  .route("/:courseId")
  .patch(courseController.updateCourse)
  .delete(courseController.deleteCourse);

module.exports = router;
