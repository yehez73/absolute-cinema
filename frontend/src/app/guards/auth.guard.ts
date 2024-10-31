import { CanActivate, Router } from '@angular/router';

export class AuthGuard implements CanActivate {
  constructor() {}

  canActivate(): boolean {
    const token = localStorage.getItem('token');
    if (token) {
      return true;
    }
    return false;
  }
};