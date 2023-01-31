<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class CourseClass extends Model
{
    use HasFactory;

    public function users() {
        return $this->belongsToMany(User::class, 'user_classes');
    }

    public function course() {
        return $this->belongsTo(Course::class);
    }

    protected $fillable = [
        'name',
        'course_id',
    ];

    public $timestamps = false;
}
