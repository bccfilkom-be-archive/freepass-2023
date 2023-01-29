const express = require("express");
const userController = require("../controllers/user.controllers");
const isLoggedIn = require("../middleware/auth.middleware");

const router = express.Router();

router.post("/register", userController.register);
router.post("/login", userController.login);

router.use(isLoggedIn);

router.get("/:id/account", userController.getUser);
router.patch("/:id/account/edit", userController.updateUser);
router
  .route("/:id/classes")
  .get(userController.getClassByUser)
  .post(userController.addClassByUser);
router.delete("/:userId/classes/:classId", userController.dropClass);
router.get(
  "/:userId/classes/:classId/students",
  userController.getStudentByClass
);

module.exports = router;
