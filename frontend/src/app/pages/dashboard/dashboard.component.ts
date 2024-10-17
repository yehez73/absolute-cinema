import { Component, OnInit } from '@angular/core';
import { Movie, MoviesService } from '../../services/movies/movies.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  movies: Movie[] = [];

  constructor(private moviesService: MoviesService) {}

  ngOnInit(): void {
    this.moviesService.getMovies().subscribe((data: Movie[]) => {
      this.movies = data;
    });
  }

  getCountryCode(language: string): string {
    const languageToCountryMap: { [key: string]: string } = {
      'EN': 'gb', // English -> Great Britain flag
      'ID': 'id', // Indonesia -> Indonesia flag
      // Tambahkan lebih banyak konversi jika diperlukan
    };
    return languageToCountryMap[language] || language.toLowerCase();
  }
}