<?php

namespace App\Http\Controllers;

use App\Http\Resources\CourseClassResource;
use App\Models\Course;
use App\Models\CourseClass;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Gate;
use Illuminate\Support\Facades\Validator;

class CourseClassController extends Controller
{   
    public function index()
    {
        $data = CourseClassResource::collection(CourseClass::all());
        return response()->json([
            'message' => 'Successed',
            'data' => $data
        ], 200);
    }

    public function store(Request $request)
    {   
        if (!Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $validator = Validator::make($request->all(), [
            'name' => 'required|string',
            'course_id' => 'required|integer'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ], 400);
        }
        
        if (Course::find($request->course_id) == null) {
            return response()->json([
                'message' => 'Course not exist!'
            ], 400);
        }

        $courseClass = new CourseClass;
        $courseClass->name = $request->name;
        $courseClass->course_id = $request->course_id;
        $courseClass->save();

        $data = new CourseClassResource($courseClass);
        return response()->json([
            'message' => 'Create Successed',
            'data' => $data
        ], 201);
    }

    public function update(Request $request, $id)
    {   
        if (!Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $validator = Validator::make($request->all(), [
            'name' => 'required|string',
            'course_id' => 'required|integer'
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => $validator->errors()
            ]);
        }

        if (Course::find($request->course_id) == null) {
            return response()->json([
                'message' => 'Course not exist!'
            ], 400);
        }

        $courseClass = CourseClass::find($id);

        if ($courseClass == null) {
            return response()->json([
                'message' => 'Class not exist!'
            ], 400);
        }

        $courseClass->name = $request->name;
        $courseClass->course_id = $request->course_id;
        $courseClass->save();

        $data = new CourseClassResource($courseClass);
        return response()->json([
            'message' => 'Update Successed',
            'data' => $data
        ], 200);
    }

    public function destroy($id)
    {   
        if (!Gate::allows('isAdmin')) {
            return response()->json([
                'message' => 'Not Authorized'
            ], 403);
        }

        $courseClass = CourseClass::find($id);
        
        if ($courseClass == null) {
            return response()->json([
                'message' => 'Class not exist!'
            ], 400);
        }   

        $courseClass->delete();
        
        return response()->json([
            'message' => 'Delete Successed',
        ], 200);
    }
}

