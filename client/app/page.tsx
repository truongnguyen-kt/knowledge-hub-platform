"use client";
import { useEffect, useState } from 'react';
import axios from 'axios';

type Lesson = {
  id: number;
  title: string;
};

export default function Home() {
  const [lessons, setLessons] = useState<Lesson[]>([]);

  useEffect(() => {
    axios.get('http://localhost:8080/api/v1/lessons')
      .then(res => setLessons(res.data))
      .catch(err => console.log("Lỗi gọi API:", err));
  }, []);

  return (
    <main className="p-10">
      <h1 className="text-3xl font-bold mb-5">Học tập lập trình</h1>
      <div className="bg-gray-100 p-5 rounded-lg">
        <h2 className="text-xl font-semibold mb-3">Danh sách bài học:</h2>
        <ul className="list-disc ml-5">
          {lessons.map(lesson => (
            <li key={lesson.id} className="text-lg">{lesson.title}</li>
          ))}
        </ul>
      </div>
    </main>
  );
}