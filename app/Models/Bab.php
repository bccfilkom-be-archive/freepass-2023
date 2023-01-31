<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Bab extends Model
{
    protected $table = "bab";
    use HasFactory;

    protected $fillable = [
        'kursus_id',
        'judul',
    ];

    public function materi()
    {
        return $this->hasMany(Materi::class)->orderBy('created_at','asc');
    }

    public function kursus()
    {
        return $this->belongsTo(Kursus::class,'kursus_id');
    }
}
