import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { LayoutComponent } from './pages/layout/layout.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';

export const routes: Routes = [
    {
        path: '', redirectTo: 'home', pathMatch: 'full' 
    },
    {
        path: 'home',
        component: DashboardComponent
    },
];