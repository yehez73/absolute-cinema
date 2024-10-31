import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MoviesService } from '../../services/movies/movies.service';
import { CommonModule } from '@angular/common';
import { ShowtimeService, GroupedShowtime } from '../../services/showtime/showtime.service';
import { catchError, of } from 'rxjs';

@Component({
  selector: 'app-movie-detail',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './movie-detail.component.html',
  styleUrls: ['./movie-detail.component.css'],
})

export class MovieDetailComponent implements OnInit {
  activeTab: string = 'schedule';
  selectedDate: string = '';
  isModalOpen: boolean = false;
  selectedSeats: number = 1;
  movie: any;
  showtimeData: GroupedShowtime | null = null;
  availableSeatsCount: number = 0;
  dateTabs: {
    date: string;
    day: string;
    month: string;
    year: string;
    isHeader: boolean;
  }[] = [];
  expandedTheaterId: string | null = null;
  selectedShowtime: any;

  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private moviesService: MoviesService,
    private showtimeService: ShowtimeService
  ) {}

  incrementSeats() {
    if (this.selectedSeats < 8) {
      this.selectedSeats++;
    }
  }

  decrementSeats() {
    if (this.selectedSeats > 1) {
      this.selectedSeats--;
    }
  }

  openModal(showtime: any) {
    this.selectedShowtime = showtime;
    this.isModalOpen = true;

    // Count available seats only for the selected showtime
    this.availableSeatsCount = showtime.available_seats
      ? showtime.available_seats.filter((seat: { is_available: any; }) => seat.is_available).length
      : 0;
  }

  closeModal() {
    this.isModalOpen = false;
    this.selectedSeats = 1;
  }

  ngOnInit(): void {
    const id = this.activatedRoute.snapshot.paramMap.get('id');
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
    const movieId = this.activatedRoute.snapshot.paramMap.get('id');
    if (movieId && this.selectedDate) {
      this.showtimeService.getGroupedShowtime(movieId, this.selectedDate)
        .pipe(
          catchError(error => {
            if (error.status === 204) {
              console.warn('No content available for the selected date.');
              this.showtimeData = null;
            } else {
              console.error('Error fetching showtime:', error);
              this.showtimeData = null;
            }
            return of(null);
          })
        )
        .subscribe(showtime => {
          if (showtime) {
            this.showtimeData = showtime;
          } else {
            this.showtimeData = null;
          }
        });
    }
  }

  toggleAccordion(theaterId: string) {
    this.expandedTheaterId = this.expandedTheaterId === theaterId ? null : theaterId;
  }

  onDateTabClick(date: string) {
    this.selectedDate = date;
    this.fetchShowtimeData();
  }

  // Method to check if a showtime is in the past based on local OS time
  isShowtimeInThePast(showtimeTime: string): boolean {
    const currentTime = new Date();
    const [hours, minutes] = showtimeTime.split(':').map(Number);
    const showtimeDate = new Date(currentTime);
    showtimeDate.setHours(hours, minutes, 0, 0);
    return showtimeDate < currentTime;
  }

  confirmSelection() {
    this.isModalOpen = false;
    const token = localStorage.getItem('token');
    if (!token) {
      this.router.navigate(['/login']);
    } else {
      this.router.navigate(['/seat-reservation'], {
        queryParams: {
          movieId: this.movie.id,
          showtimeId: this.selectedShowtime.id,
          selectedSeats: this.selectedSeats,
        },
      });
    }
  }
}