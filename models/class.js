const mongoose = require('mongoose');
const { Schema } = mongoose;

const classSchema = new Schema({
matkul:  {
    type: String
}, 
sks:  {
    type: Number
}, 
peserta: [{
    type: mongoose.Schema.Types.ObjectId,
    ref: 'user'
}]
}, { timestamps: true });

module.exports = mongoose.model('class', classSchema)
