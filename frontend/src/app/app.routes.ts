import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { LayoutComponent } from './pages/layout/layout.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { MovieComponent } from './pages/movie/movie.component';
import { MovieDetailComponent } from './pages/movie-detail/movie-detail.component';
import { SeatReservationComponent } from './pages/seat-reservation/seat-reservation.component';

export const routes: Routes = [
    {
        path: '', redirectTo: 'home', pathMatch: 'full'
    },
    {
        path: 'login',
        component: LoginComponent
    },
    {
        path: 'home',
        component: DashboardComponent
    },
    {
        path: 'movies',
        component: MovieComponent
    },
    {
        path: 'movie/:id',
        component: MovieDetailComponent
    },
    {
        path: 'seat-reservation',
        component: SeatReservationComponent
    },
];