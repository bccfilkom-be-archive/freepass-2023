const mongoose = require("mongoose");

const userSchema = mongoose.Schema({
  username: {
    required: true,
    unique: true,
    type: String,
  },
  password: {
    required: true,
    type: String,
  },
  fullName: String,
  semester: Number,
  _class: [{ type: mongoose.Schema.Types.ObjectId, ref: "Class"}],
});

const User = mongoose.model("User", userSchema);

module.exports = User;
