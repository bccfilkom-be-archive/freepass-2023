<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;

class AuthController extends Controller
{
    // public function register(Request $request)
    // {
    //     $validator = Validator::make($request->all(), [
    //         'name' => 'required|string',
    //         'email' => 'required|email',
    //         'password' => 'required|string'
    //     ]);

    //     if ($validator->fails()) {
    //         return response()->json([
    //             'message' => $validator->errors()
    //         ], 401);
    //     }

    //     $user = new User;
    //     $user->name = $request->name;
    //     $user->email = $request->email;
    //     $user->password = Hash::make($request->password);
    //     $user->role = 'user';
    //     $user->save();

    //     return response()->json([
    //         'message' => 'Registration Succesed!',
    //         'name' => $request->name,
    //         'email' => $request->email
    //     ]);
    // }

    public function register(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'string|required',
            'email' => 'email|required',
            'password' => 'string|required'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ], 401);
        }

        $user = new User;
        $user->name = $request->name;
        $user->email = $request->email;
        $user->password = Hash::make($request->password);
        $user->role = 'user';
        $user->save();

        return response()->json([
            'message' => 'Registration Succesed!',
            'name' => $request->name,
            'email' => $request->email
        ], 201);
    }

    public function login(Request $request)
    {
        if (!Auth::attempt($request->only('email', 'password'))) {
            return response()->json([
                'message' => "Login failed, Try again ! Don't forget to Register"
            ], 401);
        }

        $user = User::where('email', $request->email)->firstOrFail();

        $token = $user->createToken('auth_token')->plainTextToken;

        return response()->json([
            'message' => 'Login Succesed!',
            'token' => $token
        ]);
    }

    public function logout()
    {
        $user = User::find(Auth::user()->id);
        $user->tokens()->delete();
        return response()->json([
            'message' => 'Logout Succesed!'
        ], 200);
    }
}
