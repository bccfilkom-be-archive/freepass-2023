<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class KursusAcc extends Model
{
    use HasFactory;

    protected $table = "kursusAcc";
    protected $fillable = [
        'kursus_id',
        'user_id',
    ];
    use HasFactory;

    public function user()
    {
        return $this->belongsTo(User::class,'user_id');
    }

    public function kursus()
    {
        return $this->belongsTo(Kursus::class,'kursus_id');
    }
}
