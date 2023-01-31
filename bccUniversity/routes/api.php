<?php

use App\Http\Controllers\Api\AuthController;
use App\Http\Controllers\CourseClassController;
use App\Http\Controllers\UserClassController;
use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

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


Route::post('/register', [AuthController::class, 'register']);
Route::post('/login', [AuthController::class, 'login']);
// Route::middleware('auth:sanctum')->get('logout', [AuthController::class, 'logout']);

// Route::middleware('auth:sanctum')->get('/account', [UserController::class, 'index']);
// Route::middleware('auth:sanctum')->post('/account', [UserController::class, 'store']);

// Route::middleware('auth:sanctum')->get('/course-class', [CourseClassController::class, 'index']);
// Route::middleware('auth:sanctum')->post('/course-class', [CourseClassController::class, 'store']);
// Route::middleware('auth:sanctum')->post('/course-class/{id}', [CourseClassController::class, 'update']);
// Route::middleware('auth:sanctum')->delete('/course-class/{id}', [CourseClassController::class, 'destroy']);

// Route::middleware('auth:sanctum')->get('/user-class', [UserClassController::class, 'index']);
// Route::middleware('auth:sanctum')->post('/user-class', [UserClassController::class, 'store']);
// Route::middleware('auth:sanctum')->delete('/user-class/{id}', [UserClassController::class, 'destroy']);

Route::middleware('auth:sanctum')->group(function () {
    Route::get('logout', [AuthController::class, 'logout']);

    Route::get('/account', [UserController::class, 'index']);
    Route::post('/account', [UserController::class, 'store']);

    Route::get('/course-classes', [CourseClassController::class, 'index']);
    Route::post('/course-classes', [CourseClassController::class, 'store']);
    Route::post('/course-classes/{id}', [CourseClassController::class, 'update']);
    Route::delete('/course-classes/{id}', [CourseClassController::class, 'destroy']);

    Route::get('/user-classes', [UserClassController::class, 'index']);
    Route::post('/user-classes', [UserClassController::class, 'store']);
    Route::delete('/user-classes/{id}', [UserClassController::class, 'destroy']);
});
