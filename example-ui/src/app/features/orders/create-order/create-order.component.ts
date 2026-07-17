import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  FormBuilder,
  FormsModule,
  ReactiveFormsModule,
  Validators
} from '@angular/forms';

import { WmsApiService } from '../../../core/api/wms-api.service';
import { Router } from '@angular/router';
import { Product } from '../../../core/models/product.model';

@Component({
  selector: 'app-create-order',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,

  ],
  templateUrl: './create-order.component.html'
})
export class CreateOrderComponent {

  private readonly fb = inject(FormBuilder);
  private readonly api = inject(WmsApiService);
  private readonly router = inject(Router);


  products: Product[] = [];

  items = [
    {
      productId: 1,
      quantity: 1
    }
  ];

  success = false;



  ngOnInit(): void {
    this.api.getProducts()
      .subscribe({
        next: products => {
          this.products = products;
        }
      });
  }


  addItem(): void {
    this.items.push({
      productId: this.products[0].id,
      quantity: 1
    });
  }


  removeItem(index: number): void {
    this.items.splice(index, 1);
  }



  form = this.fb.nonNullable.group({
    productId: [
      1,
      [
        Validators.required,
        Validators.min(1)
      ]
    ],
    quantity: [
      1,
      [
        Validators.required,
        Validators.min(1)
      ]
    ],
    comment: ['']
  });


  submit(): void {

    this.api.createOrder({
      comment: this.form.controls.comment.value,
      items: this.items
    })
      .subscribe({
        next: () => {
          this.router.navigate(['/orders']);
        },
        error: console.error
      });
  }

}