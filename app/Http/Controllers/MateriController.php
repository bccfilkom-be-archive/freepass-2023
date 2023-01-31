<?php

namespace App\Http\Controllers;

use App\Models\Materi;
use Illuminate\Http\Request;
use Illuminate\Support\Str;
use Illuminate\Support\Facades\DB;

class MateriController extends Controller
{
    public function showMateri($idKursus,$idBab)
    {
        $data = Materi::where('kursus_id',$idKursus)->where('bab_id',$idBab)->orderBy('created_at', 'asc')->get();
        
        return response()->json([
            'status' => 'berhasil',
            'data' => $data,
        ]);
    }

    public function saveMateri(Request $request)
    {
        $data = new Materi;

            $data->judul = $request->judul;
            $data->kursus_id = $request->idKursus;
            $data->bab_id = $request->idBab;
            $data->judul = $request->judul;
            $data->tipe = $request->tipe;
            $data->isi = $request->isi;
            $data->save();

            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
    }

    public function updateMateri(Request $request)
    {
        $data = Materi::where('id',$request->id)->first();
        
            $data->judul = $request->judul;
            $data->kursus_id = $request->idKursus;
            $data->bab_id = $request->idBab;
            $data->judul = $request->judul;
            $data->tipe = $request->tipe;
            $data->isi = $request->isi;
            $data->save();

            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
    }

    public function deleteMateri($id)
    {
        $data = Materi::where('id',$id)->first();
        DB::table('jawaban')->where('kursus_id',$data->kursus_id)->where('bab_id',$data->bab_id)->where('materi_id',$id)->delete();
        $data->delete();

        return response()->json([
            'status' => 'berhasil',
        ],200);
    }
}
