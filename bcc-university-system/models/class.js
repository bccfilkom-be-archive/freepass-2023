const mongoose = require("mongoose");

const classSchema = mongoose.Schema({
  name: {
    required: true,
    type: String,
  },
  _course: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "Course",
  },
  _student: [
    {
      type: mongoose.Schema.Types.ObjectId,
      ref: "User",
    },
  ],
});

const Class = mongoose.model("Class", classSchema);

module.exports = Class;
