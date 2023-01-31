<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Kursus extends Model
{
    protected $table = "kursus";
    use HasFactory;

    public function bab()
    {
        return $this->hasMany(Bab::class);
    }

    public function materi()
    {
        return $this->hasMany(Materi::class);
    }

    public function kursusAcc()
    {
        return $this->hasMany(kursusAcc::class);
    }
}
