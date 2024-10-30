import { Injectable } from '@angular/core';
import axios from 'axios';
import { Observable, from } from 'rxjs';
import { environment } from '../../../environments/environment';

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
  private apiUrl = environment.apiUrl;

  constructor() {}

  getMovies(): Observable<Movie[]> {
    return from(axios.get<Movie[]>(`${this.apiUrl}/movies`).then(response => response.data));
  }

  getShowingMovies(): Observable<Movie[]> {
    return from(axios.get<Movie[]>(`${this.apiUrl}/movie/nowshowing`).then(response => response.data));
  }

  getUpcomingMovies(): Observable<Movie[]> {
    return from(axios.get<Movie[]>(`${this.apiUrl}/movie/upcoming`).then(response => response.data));
  }

  getMovieById(id: string): Observable<Movie> {
    return from(axios.get<Movie>(`${this.apiUrl}/movie/${id}`).then(response => response.data));
  }
}