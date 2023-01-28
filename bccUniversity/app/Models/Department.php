<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Department extends Model
{
    use HasFactory;

    public function studyPrograms() {
        return $this->hasMany(StudyProgram::class);
    }

    public function faculty() {
        return $this->belongsTo(Faculty::class);
    }
}
