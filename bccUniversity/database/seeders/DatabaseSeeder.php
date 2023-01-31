<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Hash;


class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        DB::table('users')->insert([
            'name' => 'admin',
            'email' => 'admin@gmail.com',
            'password' => Hash::make('admin'),
            'role' => 'admin'
        ]);

        DB::table('users')->insert([
            'name' => 'reyhan',
            'email' => 'reyhan@gmail.com',
            'password' => Hash::make('password'),
            'role' => 'user'
        ]);

        DB::table('users')->insert([
            'name' => 'budi',
            'email' => 'budi@gmail.com',
            'password' => Hash::make('password'),
            'role' => 'user'
        ]);

        DB::table('users')->insert([
            'name' => 'jono',
            'email' => 'jono@gmail.com',
            'password' => Hash::make('password'),
            'role' => 'user'
        ]);

        DB::table('users')->insert([
            'name' => 'jaka',
            'email' => 'jaka@gmail.com',
            'password' => Hash::make('password'),
            'role' => 'user'
        ]);

        DB::table('users')->insert([
            'name' => 'putri',
            'email' => 'putri@gmail.com',
            'password' => Hash::make('password'),
            'role' => 'user'
        ]);

        DB::table('courses')->insert([
            'name' => 'Computer Science',
            'course_credit' => 20,
        ]);

        DB::table('courses')->insert([
            'name' => 'Information Technology',
            'course_credit' => 20,
        ]);

        DB::table('courses')->insert([
            'name' => 'History',
            'course_credit' => 4,
        ]);

        DB::table('course_classes')->insert([
            'name' => 'CS-A',
            'course_id' => 1,
        ]);

        DB::table('course_classes')->insert([
            'name' => 'CS-B',
            'course_id' => 1,
        ]);
        
        DB::table('course_classes')->insert([
            'name' => 'IT-A',
            'course_id' => 2,
        ]);

        DB::table('course_classes')->insert([
            'name' => 'IT-B',
            'course_id' => 2,
        ]);

        DB::table('course_classes')->insert([
            'name' => 'H-A',
            'course_id' => 3,
        ]);

        DB::table('course_classes')->insert([
            'name' => 'H-B',
            'course_id' => 3,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 2,
            'course_class_id' => 1,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 2,
            'course_class_id' => 3,
        ]);
        
        DB::table('user_classes')->insert([
            'user_id' => 3,
            'course_class_id' => 2,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 3,
            'course_class_id' => 3,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 4,
            'course_class_id' => 1,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 4,
            'course_class_id' => 3,
        ]);
        
        DB::table('user_classes')->insert([
            'user_id' => 5,
            'course_class_id' => 2,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 5,
            'course_class_id' => 3,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 6,
            'course_class_id' => 1,
        ]);

        DB::table('user_classes')->insert([
            'user_id' => 6,
            'course_class_id' => 3,
        ]);

        




    }
}
