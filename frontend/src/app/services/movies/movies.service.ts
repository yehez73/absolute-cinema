import { Injectable } from '@angular/core';
import axios from 'axios';  // Impor Axios
import { Observable, from } from 'rxjs';

export interface Movie {
  id: number;
  title: string;
  description: string;
  image: string;
  genre: string;
  language: string;
  duration: string;
  release_date: string;
  rating: string;
  created_at: string;
  updated_at: string;
}

@Injectable({
  providedIn: 'root'
})
export class MoviesService {
  private apiUrl = 'http://localhost:8080/movies';

  constructor() {}

  getMovies(): Observable<Movie[]> {
    return from(axios.get<Movie[]>(this.apiUrl).then(response => response.data));
  }
}