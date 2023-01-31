const dataKelas = require('../models/class')
const user = require('../models/users')

module.exports = {
    actionEdit: async (req, res) => {
        try {
        const { id_user } = req.params
        const { nama, nip, email, phone } = req.body

        typeof (id_user)
    
        const users = await user.findOneAndUpdate({ 
          _id: id_user 
        }, { nama, nip, email, phone })
    
        res.status(200).json({
            data : users,
            metadata : "successfully edited"
        })
      
      } catch (err) {
        console.log(err)
      }
    },
    index: async (req, res) => {
        try {
            const { id } = req.params
            console.log(id)
            const users = await user.findOne({_id: id}).populate('matkul')
    
            res.status(200).json({
                data : users
            })
    
          } catch (err) {
            console.log(err)
          }
    },
    kelas: async (req, res) => {
        try {
            const kelas = await dataKelas.find().populate('peserta')
    
            res.status(200).json({
                data : kelas,
                "metadata" : "response success"
            })
          } catch (err) {
            console.log(err)
          }
    },
    create: async (req, res) => {
        try {
        const { id_class, id_user } = req.params
        const kelas = await dataKelas.findOne({_id: id_class}).populate('peserta')
        const users = await user.findOne({_id: id_user}).populate('matkul')
    
        var arr = []
        var i = 0
        users.matkul.forEach( async matkul=>{
          if(matkul._id.toString() !== kelas._id.toString()){
            arr[i] = "t"
            i++
          }else{
            arr[i] = "f"
            i++
          }
        })
    
        arr.sort()
        var jumlahsks = users.jumlahSKS + kelas.sks
        if(arr[0] !== "f" || users.matkul.length === 0){
          if(jumlahsks <= 24){
            await user.findOneAndUpdate({ _id: id_user}, { jumlahSKS: jumlahsks })
            await user.findOneAndUpdate({ 
              _id: id_user
            }, { 
              "$push": {
              "matkul": id_class
            }})
            await dataKelas.findOneAndUpdate({ 
              _id: id_class
            }, { 
              "$push": {
              "peserta": id_user
            }})
            res.status(200).json({
                metadata : "successfully added new class"
            })
    
          }else{
            res.status(200).json({
                metadata : "Melebihi jumlah sks"
            })
          }
        }else{
            res.status(200).json({
                metadata : "user sudah berada di kelas"
            })
        }
        
      } catch (err) {
        console.log(err)
      }
    },
    deleteClass: async (req, res) => {
        try {
        const { id_class, id_user } = req.params
            
        const users = await user.findOne({_id: id_user}).populate('matkul')
        const kelas = await dataKelas.findOne({_id: id_class}).populate('peserta')
    
        await user.findOneAndUpdate({ 
          _id: id_user
        }, { 
          "$pull": {
          "matkul": id_class
        }})
    
        await dataKelas.findOneAndUpdate({ 
          _id: id_class
        }, { 
          "$pull": {
          "peserta": id_user
        }})
    
        res.status(200).json({
            data : kelas,
            metadata : "successfully dropped a class"
        })
      
      } catch (err) {
        console.log(err)
      }
    },
    anggota_myClass: async (req, res) => {
        try {
            const { id_class } = req.params
            const kelas = await dataKelas.findOne({_id: id_class}).populate('peserta')
    
            res.status(200).json({
                data : kelas
            })
          } catch (err) {
            console.log(err)
          }
    }

}
