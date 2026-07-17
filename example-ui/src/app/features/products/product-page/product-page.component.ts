import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AgGridAngular } from 'ag-grid-angular';

import {
  AllCommunityModule,
  ColDef,
  ModuleRegistry,
} from 'ag-grid-community';

import { Product } from '../../../core/models/product.model';
import { WmsApiService } from '../../../core/api/wms-api.service';
import { FormsModule } from '@angular/forms';

ModuleRegistry.registerModules([
  AllCommunityModule
]);

@Component({
  selector: 'app-product-page',
  standalone: true,
  imports: [
    CommonModule,
    AgGridAngular,
    FormsModule,
  ],
  templateUrl: './product-page.component.html'
})
export class ProductPageComponent {

  private readonly api = inject(WmsApiService);


  newProductName = '';
  newProductQuantity = 0;


  products: Product[] = [];

  columnDefs: ColDef[] = [
    {
      field: 'id'
    },
    {
      field: 'name'
    },

    {
      field: 'quantity',
      headerName: 'Quantity'
    },

    {
      field: 'createdAt',
      headerName: 'Created',
      valueFormatter: params =>
        new Date(params.value).toLocaleString('pl-PL')
    }
  ];

  defaultColDef: ColDef = {
    flex: 1,
    sortable: true,
    filter: true,
    resizable: true
  };

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.api.getProducts()
      .subscribe({
        next: products => {
          this.products = products;
        },
        error: console.error
      });
  }

  createProduct(): void {

    const name = this.newProductName.trim();

    if (!name) {
      return;
    }

    this.api.createProduct({
      name,
      quantity: this.newProductQuantity
    })
      .subscribe({
        next: () => {
          this.newProductName = '';
          this.newProductQuantity = 0;

          this.loadProducts();
        },
        error: console.error
      });
  }
}