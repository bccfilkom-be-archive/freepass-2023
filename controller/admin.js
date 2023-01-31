const dataKelas = require('../models/class')
const user = require('../models/users')

module.exports = {
    deleteUser: async (req, res) => {
	    try {
        const { id_class, id_user } = req.params

        await dataKelas.findOneAndUpdate({ 
          _id: id_class
        }, { 
          "$pull": {
          "peserta": id_user
        }})

        await user.findOneAndUpdate({ 
          _id: id_user
        }, { 
          "$pull": {
          "matkul": id_class
        }})

        res.status(200).json({
            metadata : "successfully deleted an user"
        })
      
      } catch (err) {
        console.log(err)
      }
    },
    createUser: async (req, res) => {
	    try {
            const { id_class, id_user } = req.params

            await dataKelas.findOneAndUpdate({ 
              _id: id_class 
            }, { 
              "$push": {
              "peserta": id_user
            }})

            await user.updateMany({ 
              _id: id_user 
            }, { 
              "$push": {
              "matkul": id_class
            }})

            res.status(200).json({
                metadata : "successfully added new user"
            })
      
          } catch (err) {
            console.log(err)
          }
    },
    deleteClass: async (req, res) => {
	    try {
        const { id_class } = req.params

        await dataKelas.findOneAndRemove({
          _id: id_class
        });

        await user.updateMany({ 
          matkul : id_class
        }, { 
          "$pull": {
          "matkul": id_class
        }})

        res.status(200).json({
            metadata : "successfully deleted"
        })
      
      } catch (err) {
        console.log(err)
      }
    },
    CreateClass: async (req, res) => {
	    try {
            const { matkul, sks } = req.body
      
            const kelas = await dataKelas.insertMany({
              matkul,sks
            })
            
            res.status(200).json({
                data : kelas[0]._id
            })
      
          } catch (err) {
            console.log(err)
          }
    },
    editClass: async (req, res) => {
        try {
        const { id_class } = req.params
        const { matkul, sks } = req.body

        const kelas = await dataKelas.findOneAndUpdate({ 
          _id: id_class 
        }, { matkul, sks})

        res.status(200).json({
            data : kelas,
            metadata : "successfully edited"
        })
      
      } catch (err) {
        console.log(err)
      }
    },
    index: async (req, res) => {
	    try {
            const kelas = await dataKelas.find().populate('peserta')
      
            res.status(200).json({
                data : kelas,
                "metadata" : "response success"
            })
          } catch (err) {
            console.log(err)
          }
    }

}
