import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MoviesService } from '../../services/movies/movies.service';
import { CommonModule } from '@angular/common';
import { ShowtimeService } from '../../services/showtime/showtime.service';
import { catchError, of } from 'rxjs';

@Component({
  selector: 'app-movie-detail',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './movie-detail.component.html',
  styleUrls: ['./movie-detail.component.css'],
})
export class MovieDetailComponent implements OnInit {
  activeTab: string = 'details';
  selectedDate: string = ''; // Initialized to today's date
  movie: any;
  showtimeData: any; // Holds showtime data for the selected date
  availableSeatsCount: number = 0; // To hold available seats count
  dateTabs: {
    date: string;
    day: string;
    month: string;
    year: string;
    isHeader: boolean;
  }[] = [];

  constructor(
    private route: ActivatedRoute,
    private moviesService: MoviesService,
    private showtimeService: ShowtimeService
  ) {}

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.moviesService.getMovieById(id).subscribe((data) => {
        this.movie = data;
      });
    }

    // Set selected date to today
    this.selectedDate = new Date().toISOString().split('T')[0];
    this.generateDateTabs();
    this.fetchShowtimeData(); // Fetch showtime data for today initially
  }

  generateDateTabs() {
    const today = new Date();
    let lastMonth = today.getMonth();

    for (let i = 0; i < 7; i++) {
      const date = new Date(today);
      date.setDate(today.getDate() + i);
      const day = date.getDate();
      const month = date.toLocaleString('default', { month: 'short' });
      const year = date.getFullYear();
      const currentMonth = date.getMonth();

      // Add a header for a new month
      if (currentMonth !== lastMonth) {
        this.dateTabs.push({
          date: '',
          day: '',
          month: month,
          year: year.toString(),
          isHeader: true,
        });
        lastMonth = currentMonth;
      }

      // Add a date tab
      this.dateTabs.push({
        date: date.toISOString().split('T')[0],
        day: i === 0 ? 'Today' : date.toLocaleString('default', { weekday: 'short' }),
        month: '',
        year: '',
        isHeader: false,
      });
    }
  }

  fetchShowtimeData() {
    const movieId = this.route.snapshot.paramMap.get('id');
    if (movieId && this.selectedDate) {
      this.showtimeService.getShowtimeByMovieIdAndDate(movieId, this.selectedDate)
        .pipe(
          catchError(error => {
            if (error.status === 204) {
              console.warn('No content available for the selected date.');
              this.showtimeData = null;
              this.availableSeatsCount = 0;
            } else {
              console.error('Error fetching showtime:', error);
            }
            return of(null);
          })
        )
        .subscribe(showtime => {
          if (showtime) {
            // If there's showtime data, process it
            this.showtimeData = showtime;
            this.availableSeatsCount = showtime.available_seats.filter(seat => seat.is_available).length;
          } else {
            this.showtimeData = null;
            this.availableSeatsCount = 0;
          }
        });
    }
  }  

  onDateTabClick(date: string) {
    this.selectedDate = date;
    this.fetchShowtimeData();
  }
}
