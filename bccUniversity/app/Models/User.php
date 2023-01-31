<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Laravel\Sanctum\HasApiTokens;

class User extends Authenticatable
{
    use HasApiTokens, HasFactory, Notifiable;

    public function courseClasses()
    {
        return $this->belongsToMany(CourseClass::class, 'user_classes');
    }

    protected $fillable = [
        'name',
        'email',
        'password',
        'role'
    ];

    /**
     * The model's default values for attributes.
     *
     * @var array
     */
    protected $attribute = [
        'role' => 'user'
    ];

    /**
     * The attributes that should be hidden for serialization.
     *
     * @var array<int, string>
     */
    protected $hidden = [
        'password',

    ];

    public $timestamps = false;
}
