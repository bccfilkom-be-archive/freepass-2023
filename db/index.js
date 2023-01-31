const mongoose = require('mongoose');
mongoose.set("strictQuery", false);
mongoose.connect('mongodb://127.0.0.1:27017/Bcc_University' , function(err,result){
    if(err){
        console.log(`pesan error : ${err}`)
    }else{
        console.log(`berhasil tersambung ke database`)
    }
});;
