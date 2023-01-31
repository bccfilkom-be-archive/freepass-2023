<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Jawaban extends Model
{
    protected $table = "jawaban";

    protected $fillable = [
        'kursus_id',
        'bab_id',
        'materi_id',
        'komen',
        'nilai',
        'gambar',
        'email',
        'namauser',
    ];
    use HasFactory;
}
