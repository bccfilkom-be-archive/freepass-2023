<?php

namespace App\Http\Controllers;

use App\Http\Resources\UserResource;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Gate;
use Illuminate\Support\Facades\Validator;

class UserController extends Controller
{
    public function index()
    {
        if (Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $data = new UserResource(User::findOrFail(Auth::user()->id));
        return response()->json([
            'message' => 'Successed',
            'data' => $data
        ], 200);
    }

    public function store(Request $request)
    {
        if (Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $validator = Validator::make($request->all(), [
            'name' => 'required|string'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ], 400);
        }

        $user = User::find(Auth::user()->id);
        $user->name = $request->name;
        $user->save();

        $data = new UserResource($user);
        return response()->json([
            'message' => 'Update Successed',
            'data' => $data
        ], 200);
    }
}
