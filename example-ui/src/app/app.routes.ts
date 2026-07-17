import { Routes } from '@angular/router';


export const routes: Routes = [
  {
    path: '',
    redirectTo: 'orders',
    pathMatch: 'full'
  },
  {
    path: 'products',
    loadComponent: () =>
      import('./features/products/product-page/product-page.component')
        .then(m => m.ProductPageComponent)
  },
  {
    path: 'orders',
    loadComponent: () =>
      import('./features/orders/order-page/order-page.component')
        .then(m => m.OrderPageComponent)
  },
  {
    path: 'orders/new',
    loadComponent: () =>
      import('./features/orders/create-order/create-order.component')
        .then(m => m.CreateOrderComponent)

  }
];