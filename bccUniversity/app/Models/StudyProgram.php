<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class StudyProgram extends Model
{
    use HasFactory;

    public function users() {
        return $this->hasMany(User::class);
    }

    public function courses() {
        return $this->hasMany(Course::class);
    }

    public function department() {
        return $this->belongsTo(Department::class);
    }
}
