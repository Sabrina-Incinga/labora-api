export interface item {
	id: number,   
	name: string,    
	order_date: Date,
	product: string,    
	quantity: number,    
	price: number,
	details: string,
	total_price: number
}

export interface itemDTO{
	name: string,  
	order_date: string, 
	product: string,  
	quantity: number,     
	price: number, 
	details: string, 
}