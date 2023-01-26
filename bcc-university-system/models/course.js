const mongoose = require("mongoose");

const courseSchema = mongoose.Schema({
  name: {
    required: true,
    type: String,
  },
  sks: Number,
  _class: [
    {
      type: mongoose.Schema.Types.ObjectId,
      ref: "Class",
    },
  ],
});

const Course = mongoose.model("Course", courseSchema);

module.exports = Course;
