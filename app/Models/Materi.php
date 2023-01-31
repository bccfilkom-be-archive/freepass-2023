<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Materi extends Model
{
    protected $table = "materi";
    use HasFactory;

    protected $fillable = [
        'kursus_id',
        'bab_id',
        'judul',
        'tipe',
        'isi',
    ];

    public function bab()
    {
        return $this->belongsTo(Bab::class, 'bab_id');
    }

    public function kursus()
    {
        return $this->belongsTo(Kursus::class, 'kursus_id');
    }

}
