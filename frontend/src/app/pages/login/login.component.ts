import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth/auth.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  email: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router) {}

  onLogin() {
    this.authService.login(this.email, this.password).subscribe({
      next: (response) => {
        // Store the token in local storage or a service
        localStorage.setItem('token', response.token);
        this.router.navigate(['/']); // Redirect to home or movie booking page
      },
      error: (error) => {
        console.error('Login failed', error);
        // Handle login error (e.g., show a message to the user)
      }
    });
  }
}