import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, HttpClientModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  loginData: Login;

  constructor(private http: HttpClient) {
    this.loginData = new Login();
  }

  onLogin() {
    this.http.post('http://localhost:8080/login', this.loginData).subscribe((res: any) => {
      if (res.result) {
        alert('Login Success');
      } else {
        alert('Login Failed');
      }
    }
  )};
}

export class Login {
  email: string;
  password: string;
  constructor() {
    this.email = '';
    this.password = '';
  }
}