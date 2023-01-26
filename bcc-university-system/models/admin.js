const mongoose = require("mongoose");

const adminSchema = mongoose.Schema({
  username: {
    required: true,
    type: String,
  },
  password: {
    required: true,
    select: false,
    type: String,
  },
});

const Admin = mongoose.model("Admin", adminSchema);

module.exports = Admin;
