<?php

namespace App\Http\Controllers;

use App\Models\Jawaban;
use Illuminate\Http\Request;
use Illuminate\Support\Str;
use Illuminate\Support\Facades\DB;

class JawabanController extends Controller
{
    public function khususJawaban($idMateri,$email)
    {
        $data = Jawaban::where('materi_id',$idMateri)->where('email',$email)->first();

        return response()->json([
            'status' => 'berhasil',
            'data' => $data,
        ]);
    }

    public function showJawaban($idKursus,$idBab,$idMateri)
    {
        $data = Jawaban::where('kursus_id',$idKursus)->where('bab_id',$idBab)->where('materi_id',$idMateri)->get();
        
        return response()->json([
            'status' => 'berhasil',
            'data' => $data,
        ]);
    }

    public function saveJawaban(Request $request)
    {
        $data = new Jawaban;

            $image  = $request->file('gambar');
            $result = CloudinaryStorage::upload($image->getRealPath(), $image->getClientOriginalName());

            $data->kursus_id = $request->idKursus;
            $data->bab_id = $request->idBab;
            $data->materi_id = $request->idMateri;
            $data->komen = $request->komen;
            $data->nilai = $request->nilai;
            $data->gambar = $result;
            $data->namauser = $request->namauser;
            $data->email = $request->email;
            $data->save();

            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
    }

    public function updateJawaban(Request $request)
    {
        $data = Jawaban::where('id',$request->id)->first();

        if ($request->file('gambar') === null){
            $data->kursus_id = $request->idKursus;
            $data->bab_id = $request->idBab;
            $data->materi_id = $request->idMateri;
            $data->komen = $request->komen;
            $data->nilai = $request->nilai;
            $data->gambar = $request->gambar;
            $data->namauser = $request->namauser;
            $data->email = $request->email;
            $data->save();
            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
         }else{
            $file  = $request->file('gambar');
            $image = $data->gambar;
            $result = CloudinaryStorage::replace($image, $file->getRealPath(), $file->getClientOriginalName());
        
            $data->kursus_id = $request->idKursus;
            $data->bab_id = $request->idBab;
            $data->materi_id = $request->idMateri;
            $data->komen = $request->komen;
            $data->nilai = $request->nilai;
            $data->gambar = $result;
            $data->namauser = $request->namauser;
            $data->email = $request->email;
            $data->save();

            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
         }
            
    }

    public function deleteJawaban($id)
    {
        $data = Jawaban::where('id',$id)->first();
    
        $data->delete();

        return response()->json([
            'status' => 'berhasil',
        ],200);
    }
}
