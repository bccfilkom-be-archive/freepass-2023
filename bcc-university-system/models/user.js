const mongoose = require("mongoose");

const userSchema = mongoose.Schema({
  username: {
    required: true,
    type: String,
  },
  password: {
    required: true,
    select: false,
    type: String,
  },
  fullName: String,
  NIM: {
    required: true,
    unique: true,
    type: String,
  },
  _class: [{ type: mongoose.Schema.Types.ObjectId, ref: "Class" }],
});

const User = mongoose.model("User", userSchema);

module.exports = User;
