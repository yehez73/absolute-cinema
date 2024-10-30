import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Movie, MoviesService } from '../../services/movies/movies.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-movie',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './movie.component.html',
  styleUrl: './movie.component.css'
})

export class MovieComponent implements OnInit {
  nowshowingMovies: Movie[] = [];
  upcomingMovies: Movie[] = [];
  activeTab: 'nowshowing' | 'coming-soon' = 'nowshowing';

  constructor(private movieService: MoviesService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.activeTab = params['tab'] === 'coming-soon' ? 'coming-soon' : 'nowshowing';
    });

    this.movieService.getShowingMovies().subscribe((data: Movie[]) => {
      this.nowshowingMovies = data;
    });

    this.movieService.getUpcomingMovies().subscribe((data: Movie[]) => {
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

  goToMovieDetail(id: string): void {
    this.router.navigate(['movie', id]);
  }
}
