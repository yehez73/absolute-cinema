import { Component, OnInit } from '@angular/core';
import { Movie, MoviesService } from '../../services/movies/movies.service';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  movies: Movie[] = [];
  nowshowingMovies: Movie[] = [];
  upcomingMovies: Movie[] = [];

  constructor(private moviesService: MoviesService) {}

  ngOnInit(): void {
    this.moviesService.getMovies().subscribe((data: Movie[]) => {
      this.movies = data;
    });
    this.moviesService.getShowingMovies().subscribe((data: Movie[]) => {
      this.nowshowingMovies = data;
    });
    this.moviesService.getUpcomingMovies().subscribe((data: Movie[]) => {
      this.upcomingMovies = data;
    });
  }

  getCountryCode(language: string): string {
    const languageToCountryMap: { [key: string]: string } = {
      'EN': 'gb',
      'ID': 'id',
      'FR': 'fr',
      'DE': 'de',
      'IT': 'it',
      'JP': 'jp',
      'KR': 'kr',
      'RU': 'ru',
    };
    return languageToCountryMap[language] || language.toLowerCase();
  }
}