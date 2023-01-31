<?php

namespace App\Http\Controllers;

use App\Http\Resources\UserClassResource;
use App\Models\CourseClass;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Gate;
use Illuminate\Support\Facades\Validator;

class UserClassController extends Controller
{
    public function index()
    {
        if (Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $user = User::find(Auth::user()->id);

        $id = $user->courseClasses()->pluck('course_class_id');

        $data = UserClassResource::collection(CourseClass::find($id));
        return response()->json([
            'message' => 'Successed',
            'data' => $data
        ], 200);
    }

    public function store(Request $request)
    {
        if (Auth::user()->role == 'user') {
            if ($request->user_id != null) {
                return response()->json([
                    'message' => "Not Authorized"
                ], 403);
            }

            $validator = Validator::make($request->all(), [
                'course_class_id' =>  'integer|required'
            ]);

            if ($validator->fails()) {
                return response()->json([
                    'message' => $validator->errors()
                ], 400);
            }

            $user = User::find(Auth::user()->id);

            if (CourseClass::find($request->course_class_id) == null) {
                return response()->json([
                    'message' => "class not exist"
                ], 400);
            }

            $user_sks = 0;
            foreach ($user->courseClasses as $cc) {
                if ($cc->pivot->course_class_id == $request->course_class_id) {
                    return response()->json([
                        'message' => "You already join this class"
                    ], 400);
                }

                $user_sks += $cc->course()->first()->course_credit;
            }

            $temp = $user_sks + (CourseClass::find($request->course_class_id)->course->course_credit);

            if ($temp > 24) {
                return response()->json([
                    'message' => "Maximum course credit is 24"
                ], 400);
            }

            $user->courseClasses()->attach($request->course_class_id);
            return response()->json([
                'message' => 'Successed',
            ], 201);
        }

        $validator = Validator::make($request->all(), [
            'course_class_id' => 'integer|required',
            'user_id' => 'integer|required'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ], 400);
        }

        $user = User::find($request->user_id);

        if ($user == null) {
            return response()->json([
                'message' => "User not exist"
            ], 400);
        }

        if (CourseClass::find($request->course_class_id) == null) {
            return response()->json([
                'message' => "class not exist"
            ], 400);
        }

        $user_sks = 0;
        foreach ($user->courseClasses as $cc) {
            if ($cc->pivot->course_class_id == $request->course_class_id) {
                return response()->json([
                    'message' => "User already join this class"
                ], 400);
            }

            $user_sks += $cc->course()->first()->course_credit;
        }

        $temp = $user_sks + (CourseClass::find($request->course_class_id)->course->course_credit);

        if ($temp > 24) {
            return response()->json([
                'message' => "Maximum course credit is 24"
            ], 400);
        }

        $user->courseClasses()->attach($request->course_class_id);
        return response()->json([
            'message' => 'Successed',
        ], 201);
    }


    public function destroy(Request $request, $id)
    {
        if (Auth::user()->role == 'user') {
            if ($request->all() != null) {
                return response()->json([
                    'message' => "Not Authorized"
                ], 403);
            }

            $user = User::find(Auth::user()->id);

            if (CourseClass::find($id) == null) {
                return response()->json([
                    'message' => 'Class not exist!'
                ], 400);
            }

            $condition = true;

            foreach ($user->courseClasses as $cc) {
                if ($cc->pivot->course_class_id == $id) {
                    $condition = false;
                }
            }

            if ($condition) {
                return response()->json([
                    'message' => 'You are not participating in this class!'
                ], 400);
            }

            $user->courseClasses()->detach($id);

            return response()->json([
                'message' => 'Successed',
            ], 200);
        }

        $validator = Validator::make($request->all(), [
            'user_id' => 'integer|required'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ], 400);
        }

        $user = User::find($request->user_id);

        if ($user == null) {
            return response()->json([
                'message' => "User not exist"
            ], 400);
        }

        if (CourseClass::find($id) == null) {
            return response()->json([
                'message' => 'Class not exist!'
            ], 400);
        }

        $condition = true;

        foreach ($user->courseClasses as $cc) {
            if ($cc->pivot->course_class_id == $id) {
                $condition = false;
            }
        }

        if ($condition) {
            return response()->json(
                [
                    'message' => 'User are not participating in this class!'
                ],
                400
            );
        }

        $user->courseClasses()->detach($id);

        return response()->json([
            'message' => 'Successed',
        ], 200);
    }
}
