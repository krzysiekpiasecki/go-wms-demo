export interface OrderItem {
  id: number;
  orderId: number;
  productId: number;
  quantity: number;
  productName?: string;
}


export interface CreateOrderItem {
  productId: number;
  quantity: number;
}
