const mongoose = require('mongoose');
const { Schema } = mongoose;

const userSchema = new Schema({
nip:  {
    type: Number,
    required : true,
}, 
nama:  {
    type: String,
    required : true
}, 
password: {
    type: String,
    required : true
}, 
email : {
    type: String,
    default: null
},
phone : {
    type: Number,
    default: null
},
jumlahSKS : {
    type: Number,
    default: 0,
    max: 24
},
matkul : [{
    type: mongoose.Schema.Types.ObjectId,
    ref: 'class'
}],
}, { timestamps: true });

module.exports = mongoose.model('user', userSchema)
