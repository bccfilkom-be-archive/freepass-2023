require('dotenv').config()
const mongoose = require('mongoose')

mongoose.set("strictQuery", true);
exports.db = mongoose.connect(process.env.MONGO_URI, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
}).catch(err => console.error(err));