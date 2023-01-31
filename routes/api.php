<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\{
    BabController,
    AuthController,
    MateriController,
    KursusController,
    JawabanController,
};

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});



// Auth

// Link untuk register user ke system
Route::post('/register', [AuthController::class, 'register']);

// Link untuk register admin ke system atau bisa juga dari AdminSeeder.php
Route::post('/adminrandombikinregisterpokoknya', [AuthController::class, 'registerAdmin']);

// Link untuk login
Route::post('/login',[AuthController::class, 'login']);

// Link untuk edit password
Route::post('/changePassword',[AuthController::class, 'editPassword']);

// Link untuk edit profile
Route::post('/editProfile',[AuthController::class, 'editProfile']);

// Link untuk membuat access pada class yang dituju oleh user
Route::post('/createAccess',[AuthController::class, 'createAccess']);

// Link untuk menghapus access pada class yang sudah didaftarkan
Route::post('/deleteAccess/{idUser}/{idKursus}',[AuthController::class, 'deleteAccess']);

// Link untuk melihat peserta yang terdaftar pada class
Route::get('/userkursus/{idKursus}',[AuthController::class, 'userKursus']);

// Link Untuk melihat class
Route::get('/kursus',[KursusController::class, 'showKursus']);

// Link untuk membuat class
Route::post('/saveKursus', [KursusController::class, 'saveKursus']);

// Link untuk edit class
Route::post('/updateKursus', [KursusController::class, 'updateKursus']);

// Link untuk menghapus class
Route::post('/deleteKursus/{id}',[KursusController::class, 'deleteKursus']);

// Link untuk melihat kursus
Route::get('/bab/{idKursus}',[BabController::class, 'showBab']);

// Link untuk edit kursus
Route::post('/updateBab', [BabController::class, 'updateBab']);

// Link untuk membuat kursus
Route::post('/saveBab', [BabController::class, 'saveBab']);

// Link untuk menghapus kursus

// Kodingan dibawah ini untuk membuat middleware, berhubung saya membuatnya api maka untuk saat ini blom di pakai
// Route::middleware(['auth:sanctum'])->group(function(){
//     Route::middleware(['role:admin'])->group(function(){
        
//     });
// });




// ---------------------------------------------------------------------------------------------------
Route::post('/deleteBab/{id}',[BabController::class, 'deleteBab']);
Route::post('/deleteUser/{id}',[AuthController::class, 'deleteUser']);
Route::get('/user',[AuthController::class, 'showUser']);
Route::post('/userAccess',[AuthController::class, 'userAccess']);


// Kursus





// Bab



// Materi
Route::get('/materi/{idKursus}/{idBab}',[MateriController::class, 'showMateri']);
Route::post('/updateMateri', [MateriController::class, 'updateMateri']);
Route::post('/saveMateri', [MateriController::class, 'saveMateri']);
Route::post('/deleteMateri/{id}',[MateriController::class, 'deleteMateri']);

// Jawaban
Route::get('/jawaban/{idKursus}/{idBab}/{idMateri}',[JawabanController::class, 'showJawaban']);
Route::get('/jawaban/{idMateri}/{email}',[JawabanController::class, 'khususJawaban']);
Route::post('/updateJawaban', [JawabanController::class, 'updateJawaban']);
Route::post('/saveJawaban', [JawabanController::class, 'saveJawaban']);
Route::post('/deleteJawaban/{id}',[JawabanController::class, 'deleteJawaban']);