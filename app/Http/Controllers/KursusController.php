<?php

namespace App\Http\Controllers;

use App\Models\Kursus;
use App\Models\Jawaban;
use App\Http\Controllers\CloudinaryStorage;
use Illuminate\Http\Request;
use Illuminate\Support\Str;
use Illuminate\Support\Facades\DB;

class KursusController extends Controller
{
    public function showKursus()
    {
        $data = Kursus::get();
        
        if(count($data)===0){
            return response()->json([
               'data' => [],
            ]);
        }else{
            foreach ($data as $item){
                $res[] = [
                    'materi' => $item->materi,
                ];           
            }
            
            return response()->json([
                'status' => 'berhasil',
                'data'=>$data,
            ]);
        }
    }

    public function saveKursus(Request $request)
    {
        $data = new Kursus;
        
            $image  = $request->file('gambar');
            $result = CloudinaryStorage::upload($image->getRealPath(), $image->getClientOriginalName());

            $data->judul = $request->judul;
            $data->deskripsi = $request->deskripsi;
            $data->gambar = $result;
            $berhasil = $data->save();
            if($berhasil){
                return response()->json([
                    'status' => 'berhasil',
                    'data' => $data,
                ],200);
            }else{
                return response()->json([
                    'status' => 'gagal',
                ]);
            }
           
    }

    public function updateKursus(Request $request)
    {
        $data = Kursus::where('id',$request->id)->first();
         if ($request->file('gambar') === null){
            $data->judul = $request->judul;
            $data->deskripsi = $request->deskripsi;
            $data->save();
            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
         }else{
            $file  = $request->file('gambar');
            $image = $data->gambar;
            $result = CloudinaryStorage::replace($image, $file->getRealPath(), $file->getClientOriginalName());

            $data->judul = $request->judul;
            $data->gambar = $result;
            $data->deskripsi = $request->deskripsi;
            $data->save();
            return response()->json([
                'status' => 'berhasil',
                'data' => $data,
            ],200);
         }
            
    }

    public function deleteKursus($id)
    {
        $data = Kursus::where('id',$id)->first();
        DB::table('bab')->where('kursus_id',$data->id)->delete();
        DB::table('materi')->where('kursus_id',$data->id)->delete();
        DB::table('jawaban')->where('kursus_id',$data->id)->delete();
        $data->delete();
        return response()->json([
            'status' => 'berhasil',
        ],200);
    }
}
