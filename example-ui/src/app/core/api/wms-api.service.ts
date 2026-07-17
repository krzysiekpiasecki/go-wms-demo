import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable } from 'rxjs';

import { Product } from '../models/product.model';
import { Order } from '../models/order.model';

@Injectable({
  providedIn: 'root'
})
export class WmsApiService {
  private readonly http = inject(HttpClient);

  private readonly baseUrl = 'http://localhost:8080';

  getProduct(id: number): Observable<Product> {
    return this.http.get<Product>(
      `${this.baseUrl}/products/${id}`
    );
  }

  getOrder(id: number): Observable<Order> {
    return this.http.get<Order>(
      `${this.baseUrl}/orders/${id}`
    );
  }

createOrder(request: {
  comment: string;
  items: {
    productId: number;
    quantity: number;
  }[];
}) {
  return this.http.post(
    `${this.baseUrl}/orders`,
    request
  );
}

  updateOrderStatus(
    id: number,
    status: string
  ): Observable<void> {
    return this.http.patch<void>(
      `${this.baseUrl}/orders/${id}/status`,
      { status }
    );
  }

  getOrders() {
    return this.http.get<Order[]>(
      `${this.baseUrl}/orders`
    );
  }

  getProducts() {
    return this.http.get<Product[]>(
      `${this.baseUrl}/products`
    );
  }

  createProduct(request: {
  name: string;
  quantity: number;
}) {
  return this.http.post(
    `${this.baseUrl}/products`,
    request
  );
}
}