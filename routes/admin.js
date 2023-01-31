const express = require('express')
const router = express.Router()
const controller = require('../controller/admin')

router.delete('/delete-user/:id_class/:id_user', controller.deleteUser);
router.post('/add_user/:id_class/:id_user/ ', controller.createUser);
router.delete('/delete_class/:id_class/', controller.deleteClass);
router.post('/create_class', controller.CreateClass);
router.put('/edit_class/:id_class', controller.editClass);
router.get('/class', controller.index);

module.exports = router
