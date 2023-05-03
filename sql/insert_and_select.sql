INSERT INTO public.items(
	 customer_name, order_date, product, quantity, price)
	VALUES ( 'Florinda', '2023-04-02', 'Florero', 5, 5000),
			( 'Juan', '2023-04-20', 'Cenicero', 3, 2000),
			( 'Soledad', '2023-04-27', 'Taza', 6, 1000),
			( 'Jacinto', '2023-04-30', 'Cuchillos', 2, 1500),
			( 'Giselle', '2023-05-03', 'Velador', 1, 2100);
			
			
SELECT id, customer_name, order_date, product, quantity, price
	FROM public.items
	WHERE quantity > 2 AND price > 50;