SELECT  customer_name, count(customer_name) AS Cantidad_repeticiones
	FROM public.items
	GROUP BY customer_name
	HAVING count(customer_name)>1;
	
CREATE INDEX idx_items_customer_name ON items (customer_name);
	
EXPLAIN ANALYZE SELECT id, customer_name, order_date, product, quantity, price
	FROM public.items
	WHERE customer_name = 'Florinda';
	
EXPLAIN ANALYZE SELECT  id, customer_name, order_date, product, quantity, price
	FROM public.items
	WHERE product = 'Florero';
	
CREATE INDEX idx_items_product ON items (product);
