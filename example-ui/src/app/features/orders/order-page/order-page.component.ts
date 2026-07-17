import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AgGridAngular } from 'ag-grid-angular';
import { forkJoin } from 'rxjs';

import {
  AllCommunityModule,
  ColDef,
  ModuleRegistry
} from 'ag-grid-community';

import { WmsApiService } from '../../../core/api/wms-api.service';
import { Order } from '../../../core/models/order.model';
import { Router, RouterLink } from '@angular/router';

ModuleRegistry.registerModules([
  AllCommunityModule
]);

@Component({
  selector: 'app-order-page',
  standalone: true,
  imports: [
    CommonModule,
    AgGridAngular,
    RouterLink,    
  ],
  templateUrl: './order-page.component.html'
})
export class OrderPageComponent {

  private readonly api = inject(WmsApiService);
  selectedOrder?: Order;

  orders: Order[] = [];

  columnDefs: ColDef[] = [
    {
      field: 'id'
    },
    {
      field: 'status'
    },
    {
      field: 'comment'
    },
    {
      field: 'createdAt',
      headerName: 'Created',
      valueFormatter: params =>
        new Date(params.value).toLocaleString('pl-PL')
    },
  ];


  itemColumnDefs: ColDef[] = [
    {
      field: 'id'
    },
    {
      field: 'productName',
      headerName: 'Product'
    },
    {
      field: 'quantity'
    }
  ];


  defaultColDef: ColDef = {
    flex: 1,
    sortable: true,
    filter: true,
    resizable: true
  };

  ngOnInit(): void {
    this.loadOrders();
  }


  onRowClicked(event: any): void {
    this.api.getOrder(event.data.id)
      .subscribe({
        next: order => {

          const items = order.items ?? [];

          const requests = items.map(item =>
            this.api.getProduct(item.productId)
          );

          forkJoin(requests).subscribe({
            next: products => {

              items.forEach((item, index) => {
                item.productName = products[index].name;
              });

              this.selectedOrder = order;
            },
            error: console.error
          });
        },
        error: console.error
      });
  }

  loadOrders(): void {
    this.api.getOrders()
      .subscribe({
        next: orders => {
          this.orders = orders;
        },
        error: console.error
      });
  }
}