import { OrderItem } from './order-item.model';

export interface Order {
  id: number;
  status: string;
  comment: string | null;
  createdAt: string;
  items: OrderItem[] | null;
}
