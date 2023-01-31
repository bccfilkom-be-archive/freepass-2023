const express = require('express')
const router = express.Router()
const controller = require('../controller/users')

router.put('/profile/edit/:id_user', controller.actionEdit);
router.get('/profile/:id ', controller.index);
router.get('/class/', controller.kelas);
router.post('/add_class/:id_class/:id_user', controller.create);
router.delete('/delete_class/:id_class/:id_user', controller.deleteClass);
router.delete('/my_class/:id_class', controller.anggota_myClass);

module.exports = router
