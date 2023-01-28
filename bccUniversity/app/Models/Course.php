<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Course extends Model
{
    use HasFactory;

    public function courseClasses() {
        return $this->hasMany(Course::class);
    }

    public function studyProgram() {
        return $this->belongsTo(StudyProgram::class);
    }
}
